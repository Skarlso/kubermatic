// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterTemplateLister helps list ClusterTemplates.
// All objects returned here must be treated as read-only.
type ClusterTemplateLister interface {
	// List lists all ClusterTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTemplate, err error)
	// Get retrieves the ClusterTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterTemplate, error)
	ClusterTemplateListerExpansion
}

// clusterTemplateLister implements the ClusterTemplateLister interface.
type clusterTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterTemplateLister returns a new ClusterTemplateLister.
func NewClusterTemplateLister(indexer cache.Indexer) ClusterTemplateLister {
	return &clusterTemplateLister{indexer: indexer}
}

// List lists all ClusterTemplates in the indexer.
func (s *clusterTemplateLister) List(selector labels.Selector) (ret []*v1.ClusterTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterTemplate from the index for a given name.
func (s *clusterTemplateLister) Get(name string) (*v1.ClusterTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustertemplate"), name)
	}
	return obj.(*v1.ClusterTemplate), nil
}
