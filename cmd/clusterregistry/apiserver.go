/*
Copyright 2014 The Kubernetes Authors.

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

// clusterregistry is the main API server that serves the Cluster Registry API.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/cluster-registry/cmd/clusterregistry/app"
	"k8s.io/cluster-registry/cmd/clusterregistry/app/options"

	"github.com/spf13/pflag"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)
	flag.InitFlags()

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := app.Run(s, wait.NeverStop); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
