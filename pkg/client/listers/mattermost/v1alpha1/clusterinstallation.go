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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterInstallationLister helps list ClusterInstallations.
type ClusterInstallationLister interface {
	// List lists all ClusterInstallations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterInstallation, err error)
	// ClusterInstallations returns an object that can list and get ClusterInstallations.
	ClusterInstallations(namespace string) ClusterInstallationNamespaceLister
	ClusterInstallationListerExpansion
}

// clusterInstallationLister implements the ClusterInstallationLister interface.
type clusterInstallationLister struct {
	indexer cache.Indexer
}

// NewClusterInstallationLister returns a new ClusterInstallationLister.
func NewClusterInstallationLister(indexer cache.Indexer) ClusterInstallationLister {
	return &clusterInstallationLister{indexer: indexer}
}

// List lists all ClusterInstallations in the indexer.
func (s *clusterInstallationLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterInstallation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterInstallation))
	})
	return ret, err
}

// ClusterInstallations returns an object that can list and get ClusterInstallations.
func (s *clusterInstallationLister) ClusterInstallations(namespace string) ClusterInstallationNamespaceLister {
	return clusterInstallationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterInstallationNamespaceLister helps list and get ClusterInstallations.
type ClusterInstallationNamespaceLister interface {
	// List lists all ClusterInstallations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterInstallation, err error)
	// Get retrieves the ClusterInstallation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ClusterInstallation, error)
	ClusterInstallationNamespaceListerExpansion
}

// clusterInstallationNamespaceLister implements the ClusterInstallationNamespaceLister
// interface.
type clusterInstallationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterInstallations in the indexer for a given namespace.
func (s clusterInstallationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterInstallation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterInstallation))
	})
	return ret, err
}

// Get retrieves the ClusterInstallation from the indexer for a given namespace and name.
func (s clusterInstallationNamespaceLister) Get(name string) (*v1alpha1.ClusterInstallation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterinstallation"), name)
	}
	return obj.(*v1alpha1.ClusterInstallation), nil
}
