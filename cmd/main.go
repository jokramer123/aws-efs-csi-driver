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

package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/klog"

	"github.com/kubernetes-sigs/aws-efs-csi-driver/pkg/driver"
)

func main() {
	var (
		endpoint                = flag.String("endpoint", "unix://csi/csi.sock", "CSI Endpoint")
		version                 = flag.Bool("version", false, "Print the version and exit")
		efsUtilsCfgDirPath      = flag.String("efs-utils-config-dir-path", "/etc/amazon/efs/", "The path to efs-utils config directory")
		efsUtilsStaticFilesPath = flag.String("efs-utils-static-files-path", "/etc/amazon/efs-static-files/", "The path to efs-utils static files directory")
		volMetricsOptIn         = flag.Bool("vol-metrics-opt-in", false, "Opt in to emit volume metrics")
		volMetricsRefreshPeriod = flag.Float64("vol-metrics-refresh-period", 240, "Refresh period for volume metrics in minutes")
		volMetricsFsRateLimit   = flag.Int("vol-metrics-fs-rate-limit", 5, "Volume metrics routines rate limiter per file system")
	)
	klog.InitFlags(nil)
	flag.Parse()

	if *version {
		info, err := driver.GetVersionJSON()
		if err != nil {
			klog.Fatalln(err)
		}
		fmt.Println(info)
		os.Exit(0)
	}

	drv := driver.NewDriver(*endpoint, *efsUtilsCfgDirPath, *efsUtilsStaticFilesPath, *volMetricsOptIn, *volMetricsRefreshPeriod, *volMetricsFsRateLimit)
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}
}
