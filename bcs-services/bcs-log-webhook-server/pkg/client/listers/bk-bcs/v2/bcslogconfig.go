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

package v2

import (
	v2 "bk-bcs/bcs-services/bcs-log-webhook-server/pkg/apis/bk-bcs/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BcsLogConfigLister helps list BcsLogConfigs.
type BcsLogConfigLister interface {
	// List lists all BcsLogConfigs in the indexer.
	List(selector labels.Selector) (ret []*v2.BcsLogConfig, err error)
	// BcsLogConfigs returns an object that can list and get BcsLogConfigs.
	BcsLogConfigs(namespace string) BcsLogConfigNamespaceLister
	BcsLogConfigListerExpansion
}

// bcsLogConfigLister implements the BcsLogConfigLister interface.
type bcsLogConfigLister struct {
	indexer cache.Indexer
}

// NewBcsLogConfigLister returns a new BcsLogConfigLister.
func NewBcsLogConfigLister(indexer cache.Indexer) BcsLogConfigLister {
	return &bcsLogConfigLister{indexer: indexer}
}

// List lists all BcsLogConfigs in the indexer.
func (s *bcsLogConfigLister) List(selector labels.Selector) (ret []*v2.BcsLogConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsLogConfig))
	})
	return ret, err
}

// BcsLogConfigs returns an object that can list and get BcsLogConfigs.
func (s *bcsLogConfigLister) BcsLogConfigs(namespace string) BcsLogConfigNamespaceLister {
	return bcsLogConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BcsLogConfigNamespaceLister helps list and get BcsLogConfigs.
type BcsLogConfigNamespaceLister interface {
	// List lists all BcsLogConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.BcsLogConfig, err error)
	// Get retrieves the BcsLogConfig from the indexer for a given namespace and name.
	Get(name string) (*v2.BcsLogConfig, error)
	BcsLogConfigNamespaceListerExpansion
}

// bcsLogConfigNamespaceLister implements the BcsLogConfigNamespaceLister
// interface.
type bcsLogConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BcsLogConfigs in the indexer for a given namespace.
func (s bcsLogConfigNamespaceLister) List(selector labels.Selector) (ret []*v2.BcsLogConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsLogConfig))
	})
	return ret, err
}

// Get retrieves the BcsLogConfig from the indexer for a given namespace and name.
func (s bcsLogConfigNamespaceLister) Get(name string) (*v2.BcsLogConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("bcslogconfig"), name)
	}
	return obj.(*v2.BcsLogConfig), nil
}