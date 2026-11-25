/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// kustomize-build applies a kustomization to a generated manifests file and
// replaces the file with the result. It uses the kustomize library vendored
// for manifests-gen, so it needs no kustomize binary. It is meant to run on
// freshly generated output, as the ocp-manifests target in openshift/Makefile
// does.
//
// Local file access is limited to the regular files directly in the
// kustomization directory and to the manifests file named on the command line:
// only those are copied into an in-memory filesystem, at their on-disk paths,
// before the build. Subdirectories are not loaded. Remote (http or git)
// references are not blocked by kustomize and must not be used.
//
// The kustomization may only patch the resources of the manifests file. The
// build fails if it adds, removes, renames or reorders resources. The manifests
// file is left unchanged if the build or these checks fail.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/filesys"
	"sigs.k8s.io/kustomize/kyaml/kio"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <kustomization-dir> <manifests-file>\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	if err := run(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(kustomizationDir, manifestsFile string) (err error) {
	kustomizationDir, err = filepath.Abs(kustomizationDir)
	if err != nil {
		return fmt.Errorf("error resolving kustomization directory: %w", err)
	}

	manifestsFile, err = filepath.Abs(manifestsFile)
	if err != nil {
		return fmt.Errorf("error resolving manifests file: %w", err)
	}

	fSys := filesys.MakeFsInMemory()

	if err := copyDirFiles(fSys, kustomizationDir); err != nil {
		return err
	}

	// The manifests file is read and replaced through a root scoped to its
	// directory.
	manifestsRoot, err := os.OpenRoot(filepath.Dir(manifestsFile))
	if err != nil {
		return fmt.Errorf("error opening manifests directory: %w", err)
	}
	defer func() {
		err = errors.Join(err, manifestsRoot.Close())
	}()

	manifestsName := filepath.Base(manifestsFile)

	input, err := copyFile(fSys, manifestsRoot, manifestsName)
	if err != nil {
		return err
	}

	opts := krusty.MakeDefaultOptions()
	// Keep the input resource order so that only patched fields change.
	opts.Reorder = krusty.ReorderOptionNone
	// The kustomization refers to the manifests file outside its own
	// directory. Local file access is still limited to the files in fSys.
	opts.LoadRestrictions = types.LoadRestrictionsNone

	resources, err := krusty.MakeKustomizer(opts).Run(fSys, kustomizationDir)
	if err != nil {
		return fmt.Errorf("error building kustomization %q (see the comments in its kustomization.yaml): %w", kustomizationDir, err)
	}

	output, err := resources.AsYaml()
	if err != nil {
		return fmt.Errorf("error rendering kustomization %q: %w", kustomizationDir, err)
	}

	if err := checkSameResources(input, output); err != nil {
		return fmt.Errorf("kustomization %q must only patch the resources of %q: %w", kustomizationDir, manifestsFile, err)
	}

	// The manifests file is only replaced after all checks have passed.
	if err := manifestsRoot.WriteFile(manifestsName, output, 0o600); err != nil {
		return fmt.Errorf("error writing manifests file %q: %w", manifestsFile, err)
	}

	return nil
}

// checkSameResources returns an error unless output contains the resources of
// a non-empty input, identified by apiVersion, kind, namespace and name, in the
// same order.
func checkSameResources(input, output []byte) error {
	inputIDs, err := resourceIDs(input)
	if err != nil {
		return fmt.Errorf("error parsing input: %w", err)
	}

	if len(inputIDs) == 0 {
		return errors.New("input has no resources")
	}

	outputIDs, err := resourceIDs(output)
	if err != nil {
		return fmt.Errorf("error parsing output: %w", err)
	}

	if len(outputIDs) != len(inputIDs) {
		return fmt.Errorf("output has %d resources, expected %d", len(outputIDs), len(inputIDs))
	}

	for i := range inputIDs {
		if outputIDs[i] != inputIDs[i] {
			return fmt.Errorf("resource at index %d is %q, expected %q", i, outputIDs[i], inputIDs[i])
		}
	}

	return nil
}

// resourceIDs returns the apiVersion/kind/namespace/name of each resource in
// data.
func resourceIDs(data []byte) ([]string, error) {
	nodes, err := kio.FromBytes(data)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(nodes))
	for _, node := range nodes {
		ids = append(ids, fmt.Sprintf("%s/%s/%s/%s", node.GetApiVersion(), node.GetKind(), node.GetNamespace(), node.GetName()))
	}

	return ids, nil
}

// copyDirFiles copies the regular files directly in dir into fSys.
func copyDirFiles(fSys filesys.FileSystem, dir string) (err error) {
	root, err := os.OpenRoot(dir)
	if err != nil {
		return fmt.Errorf("error opening kustomization directory: %w", err)
	}
	defer func() {
		err = errors.Join(err, root.Close())
	}()

	entries, err := fs.ReadDir(root.FS(), ".")
	if err != nil {
		return fmt.Errorf("error reading kustomization directory %q: %w", dir, err)
	}

	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}

		if _, err := copyFile(fSys, root, entry.Name()); err != nil {
			return err
		}
	}

	return nil
}

// copyFile copies the file name in root into fSys, at its on-disk path, and
// returns its content.
func copyFile(fSys filesys.FileSystem, root *os.Root, name string) ([]byte, error) {
	path := filepath.Join(root.Name(), name)

	data, err := root.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("error reading %q: %w", path, err)
	}

	if err := fSys.WriteFile(path, data); err != nil {
		return nil, fmt.Errorf("error copying %q into the in-memory filesystem: %w", path, err)
	}

	return data, nil
}
