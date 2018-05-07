package v1

import (
	time "time"

	versioned "github.com/kubermatic/kubermatic/api/pkg/crd/client/clientset/versioned"
	internalinterfaces "github.com/kubermatic/kubermatic/api/pkg/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/kubermatic/kubermatic/api/pkg/crd/client/listers/kubermatic/v1"
	kubermatic_v1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectInformer provides access to a shared informer and lister for
// Projects.
type ProjectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProjectLister
}

type projectInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewProjectInformer constructs a new informer for Project type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredProjectInformer constructs a new informer for Project type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().Projects().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().Projects().Watch(options)
			},
		},
		&kubermatic_v1.Project{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermatic_v1.Project{}, f.defaultInformer)
}

func (f *projectInformer) Lister() v1.ProjectLister {
	return v1.NewProjectLister(f.Informer().GetIndexer())
}
