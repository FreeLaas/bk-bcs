/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedStatefulSetLister helps list FederatedStatefulSets.
type FederatedStatefulSetLister interface {
	// List lists all FederatedStatefulSets in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedStatefulSet, err error)
	// FederatedStatefulSets returns an object that can list and get FederatedStatefulSets.
	FederatedStatefulSets(namespace string) FederatedStatefulSetNamespaceLister
	FederatedStatefulSetListerExpansion
}

// federatedStatefulSetLister implements the FederatedStatefulSetLister interface.
type federatedStatefulSetLister struct {
	indexer cache.Indexer
}

// NewFederatedStatefulSetLister returns a new FederatedStatefulSetLister.
func NewFederatedStatefulSetLister(indexer cache.Indexer) FederatedStatefulSetLister {
	return &federatedStatefulSetLister{indexer: indexer}
}

// List lists all FederatedStatefulSets in the indexer.
func (s *federatedStatefulSetLister) List(selector labels.Selector) (ret []*v1beta1.FederatedStatefulSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedStatefulSet))
	})
	return ret, err
}

// FederatedStatefulSets returns an object that can list and get FederatedStatefulSets.
func (s *federatedStatefulSetLister) FederatedStatefulSets(namespace string) FederatedStatefulSetNamespaceLister {
	return federatedStatefulSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedStatefulSetNamespaceLister helps list and get FederatedStatefulSets.
type FederatedStatefulSetNamespaceLister interface {
	// List lists all FederatedStatefulSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedStatefulSet, err error)
	// Get retrieves the FederatedStatefulSet from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedStatefulSet, error)
	FederatedStatefulSetNamespaceListerExpansion
}

// federatedStatefulSetNamespaceLister implements the FederatedStatefulSetNamespaceLister
// interface.
type federatedStatefulSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedStatefulSets in the indexer for a given namespace.
func (s federatedStatefulSetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedStatefulSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedStatefulSet))
	})
	return ret, err
}

// Get retrieves the FederatedStatefulSet from the indexer for a given namespace and name.
func (s federatedStatefulSetNamespaceLister) Get(name string) (*v1beta1.FederatedStatefulSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedstatefulset"), name)
	}
	return obj.(*v1beta1.FederatedStatefulSet), nil
}
