package main

import (
	"context"
	"errors"
	"time"

	clientset "github.com/AvistoTelecom/kubebrowser/pkg/client/clientset/versioned"
	informers "github.com/AvistoTelecom/kubebrowser/pkg/client/informers/externalversions"
	v1alpha1 "github.com/AvistoTelecom/kubebrowser/pkg/client/listers/kubeconfig/v1alpha1"
	"github.com/spf13/viper"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

var kubecfg = &Kubecfg{}

type Kubecfg struct {
	lister v1alpha1.KubeconfigLister
}

// Setup the Kubernetes client and the SharedInformerFactory
// Returns a KubeconfigLister
func (k *Kubecfg) Init(ctx context.Context) error {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return err
	}

	// Create the namespace-scoped informer factory
	kubeInformerFactory := informers.NewSharedInformerFactoryWithOptions(
		exampleClient,
		time.Second*30,
		informers.WithNamespace(viper.GetString(podNamespaceKey)),
	)

	k.lister = kubeInformerFactory.Kubeconfig().V1alpha1().Kubeconfigs().Lister()

	kubeInformerFactory.Start(ctx.Done())

	if !cache.WaitForCacheSync(ctx.Done(), kubeInformerFactory.Kubeconfig().V1alpha1().Kubeconfigs().Informer().HasSynced) {
		return errors.New("failed to sync caches")
	}

	return nil
}
