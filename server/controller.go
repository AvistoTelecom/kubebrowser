package main

import (
	"context"
	"errors"
	"os"
	"time"

	clientset "github.com/brawdunoir/kubebrowser/pkg/client/clientset/versioned"
	informers "github.com/brawdunoir/kubebrowser/pkg/client/informers/externalversions"
	v1 "github.com/brawdunoir/kubebrowser/pkg/client/listers/kubeconfig/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

var (
	namespace = os.Getenv("POD_NAMESPACE")
)

// Setup the Kubernetes client and the SharedInformerFactory
// Returns a KubeconfigLister
func setupKubeconfigLister(ctx context.Context) (kubeconfigLister v1.KubeconfigLister, err error) {
	// creates the in-cluster config
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	// Create the namespace-scoped informer factory
	kubeInformerFactory := informers.NewSharedInformerFactoryWithOptions(
		exampleClient,
		time.Second*30,
		informers.WithNamespace(namespace),
	)

	// Get the lister for Kubeconfigs
	kubeconfigLister = kubeInformerFactory.Kubeconfig().V1().Kubeconfigs().Lister()

	// Start the informer factory
	kubeInformerFactory.Start(ctx.Done())

	// Wait for the caches to sync
	if !cache.WaitForCacheSync(ctx.Done(), kubeInformerFactory.Kubeconfig().V1().Kubeconfigs().Informer().HasSynced) {
		return nil, errors.New("failed to sync caches")
	}

	return kubeconfigLister, nil
}
