/*
Copyright 2019 The Kubernetes Authors.

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

package machine

import (
	"encoding/json"

	machinev1 "github.com/openshift/api/machine/v1beta1"
)

// updateMachineAnnotationJSON updates the `annotation` on `machine` with
// `content`. `content` in this case should be a `map[string]interface{}`
// suitable for turning into JSON. This `content` map will be marshalled into a
// JSON string before being set as the given `annotation`.
func (a *Actuator) updateMachineAnnotationJSON(machine *machinev1.Machine, annotation string, content map[string]interface{}) error {
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}

	a.updateMachineAnnotation(machine, annotation, string(b))
	return nil
}

// updateMachineAnnotation updates the `annotation` on the given `machine` with
// `content`.
func (a *Actuator) updateMachineAnnotation(machine *machinev1.Machine, annotation string, content string) {
	// Get the annotations
	annotations := machine.Annotations
	if annotations == nil {
		machine.Annotations = map[string]string{}
	}
	// Set our annotation to the given content.
	machine.Annotations[annotation] = content
}

// Returns a map[string]interface from a JSON annotation.
// This method gets the given `annotation` from the `machine` and unmarshalls it
// from a JSON string into a `map[string]interface{}`.
func (a *Actuator) machineAnnotationJSON(machine *machinev1.Machine, annotation string) (map[string]interface{}, error) {
	out := map[string]interface{}{}

	jsonAnnotation := a.machineAnnotation(machine, annotation)
	if len(jsonAnnotation) == 0 {
		return out, nil
	}

	err := json.Unmarshal([]byte(jsonAnnotation), &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// Fetches the specific machine annotation.
func (a *Actuator) machineAnnotation(machine *machinev1.Machine, annotation string) string {
	if machine.Annotations == nil {
		return ""
	}
	return machine.Annotations[annotation]
}
