/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"flag"
	"time"

	clusterapis "github.com/openshift/cluster-api/pkg/apis"
	"github.com/openshift/cluster-api/pkg/controller/machine"
	"k8s.io/klog"
	machineactuator "sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators/machine"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

func main() {
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)

	flag.Parse()
	flag.VisitAll(func(f1 *flag.Flag) {
		f2 := klogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			f2.Value.Set(value)
		}
	})

	cfg := config.GetConfigOrDie()

	// Setup a Manager
	syncPeriod := 10 * time.Minute
	mgr, err := manager.New(cfg, manager.Options{
		SyncPeriod: &syncPeriod,
	})
	if err != nil {
		klog.Fatalf("Failed to set up overall controller manager: %v", err)
	}

	if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	if err := machine.AddWithActuator(mgr, initActuator(mgr)); err != nil {
		klog.Fatalf("Error adding actuator: %v", err)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		klog.Fatalf("Failed to run manager: %v", err)
	}
}

func initActuator(mgr manager.Manager) *machineactuator.Actuator {
	params := machineactuator.ActuatorParams{
		Client:        mgr.GetClient(),
		EventRecorder: mgr.GetRecorder("azure-controller"),
	}

	return machineactuator.NewActuator(params)
}
