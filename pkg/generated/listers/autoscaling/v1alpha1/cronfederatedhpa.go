/*
Copyright The Karmada Authors.

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
	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/autoscaling/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CronFederatedHPALister helps list CronFederatedHPAs.
// All objects returned here must be treated as read-only.
type CronFederatedHPALister interface {
	// List lists all CronFederatedHPAs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CronFederatedHPA, err error)
	// CronFederatedHPAs returns an object that can list and get CronFederatedHPAs.
	CronFederatedHPAs(namespace string) CronFederatedHPANamespaceLister
	CronFederatedHPAListerExpansion
}

// cronFederatedHPALister implements the CronFederatedHPALister interface.
type cronFederatedHPALister struct {
	indexer cache.Indexer
}

// NewCronFederatedHPALister returns a new CronFederatedHPALister.
func NewCronFederatedHPALister(indexer cache.Indexer) CronFederatedHPALister {
	return &cronFederatedHPALister{indexer: indexer}
}

// List lists all CronFederatedHPAs in the indexer.
func (s *cronFederatedHPALister) List(selector labels.Selector) (ret []*v1alpha1.CronFederatedHPA, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronFederatedHPA))
	})
	return ret, err
}

// CronFederatedHPAs returns an object that can list and get CronFederatedHPAs.
func (s *cronFederatedHPALister) CronFederatedHPAs(namespace string) CronFederatedHPANamespaceLister {
	return cronFederatedHPANamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronFederatedHPANamespaceLister helps list and get CronFederatedHPAs.
// All objects returned here must be treated as read-only.
type CronFederatedHPANamespaceLister interface {
	// List lists all CronFederatedHPAs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CronFederatedHPA, err error)
	// Get retrieves the CronFederatedHPA from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CronFederatedHPA, error)
	CronFederatedHPANamespaceListerExpansion
}

// cronFederatedHPANamespaceLister implements the CronFederatedHPANamespaceLister
// interface.
type cronFederatedHPANamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronFederatedHPAs in the indexer for a given namespace.
func (s cronFederatedHPANamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CronFederatedHPA, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CronFederatedHPA))
	})
	return ret, err
}

// Get retrieves the CronFederatedHPA from the indexer for a given namespace and name.
func (s cronFederatedHPANamespaceLister) Get(name string) (*v1alpha1.CronFederatedHPA, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cronfederatedhpa"), name)
	}
	return obj.(*v1alpha1.CronFederatedHPA), nil
}