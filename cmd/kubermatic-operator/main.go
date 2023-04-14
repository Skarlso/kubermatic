/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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
	"context"
	"flag"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	seedctrl "k8c.io/kubermatic/v3/pkg/controller/operator/seed"
	kubermaticlog "k8c.io/kubermatic/v3/pkg/log"
	"k8c.io/kubermatic/v3/pkg/pprof"
	kubernetesprovider "k8c.io/kubermatic/v3/pkg/provider/kubernetes"
	"k8c.io/kubermatic/v3/pkg/resources/reconciling"
	"k8c.io/kubermatic/v3/pkg/util/edition"
	"k8c.io/kubermatic/v3/pkg/version/kubermatic"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/klog/v2"
	ctrlruntime "sigs.k8s.io/controller-runtime"
	ctrlruntimelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

type controllerRunOptions struct {
	namespace            string
	internalAddr         string
	workerCount          int
	workerName           string
	enableLeaderElection bool
}

func main() {
	ctx := signals.SetupSignalHandler()

	klog.InitFlags(nil)

	pprofOpts := &pprof.Opts{}
	pprofOpts.AddFlags(flag.CommandLine)
	logOpts := kubermaticlog.NewDefaultOptions()
	logOpts.AddFlags(flag.CommandLine)

	opt := &controllerRunOptions{}
	flag.StringVar(&opt.namespace, "namespace", "", "The namespace the operator runs in, uses to determine where to look for KubermaticConfigurations.")
	flag.IntVar(&opt.workerCount, "worker-count", 4, "Number of workers which process reconcilings in parallel.")
	flag.StringVar(&opt.internalAddr, "internal-address", "127.0.0.1:8085", "The address on which the /metrics endpoint will be served")
	flag.StringVar(&opt.workerName, "worker-name", "", "The name of the worker that will only processes resources with label=worker-name.")
	flag.BoolVar(&opt.enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	rawLog := kubermaticlog.New(logOpts.Debug, logOpts.Format).Named(opt.workerName)
	log := rawLog.Sugar()

	// update global logger instance
	kubermaticlog.Logger = log
	reconciling.Configure(log)

	// set the logger used by sigs.k8s.io/controller-runtime
	ctrlruntimelog.SetLogger(zapr.NewLogger(rawLog.WithOptions(zap.AddCallerSkip(1))))

	if len(opt.namespace) == 0 {
		log.Fatal("-namespace is a mandatory flag")
	}

	v := kubermatic.NewDefaultVersions(edition.CommunityEdition)
	log.With("kubermatic", v.Kubermatic, "ui", v.UI).Infof("Moin, moin, I'm the Kubermatic %s Operator and these are the versions I work with.", v.KubermaticEdition)

	mgr, err := manager.New(ctrlruntime.GetConfigOrDie(), manager.Options{
		BaseContext: func() context.Context {
			return ctx
		},
		MetricsBindAddress: opt.internalAddr,
		LeaderElection:     opt.enableLeaderElection,
		LeaderElectionID:   "operator.kubermatic.k8c.io",
	})
	if err != nil {
		log.Fatalw("Failed to create Controller Manager instance", zap.Error(err))
	}

	if err := mgr.Add(pprofOpts); err != nil {
		log.Fatalw("Failed to add pprof endpoint", zap.Error(err))
	}

	if err := kubermaticv1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatalw("Failed to register scheme", zap.Stringer("api", kubermaticv1.SchemeGroupVersion), zap.Error(err))
	}

	if err := apiextensionsv1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatalw("Failed to register scheme", zap.Stringer("api", apiextensionsv1.SchemeGroupVersion), zap.Error(err))
	}

	configGetter, err := kubernetesprovider.DynamicKubermaticConfigurationGetterFactory(mgr.GetClient(), opt.namespace)
	if err != nil {
		log.Fatalw("Failed to construct configGetter", zap.Error(err))
	}

	if err := seedctrl.Add(ctx, log, opt.namespace, mgr, configGetter, opt.workerCount, opt.workerName, v); err != nil {
		log.Fatalw("Failed to add seed-operator controller", zap.Error(err))
	}

	if err := mgr.Start(ctx); err != nil {
		log.Fatalw("Cannot start manager", zap.Error(err))
	}
}
