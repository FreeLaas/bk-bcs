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

package kube

import (
	ingressv1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/apis/clb/v1"
	"bk-bcs/bcs-services/bcs-clb-controller/pkg/clbingress"
	informerv1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/client/informers/clb/v1"
	ingressClientV1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/client/internalclientset/typed/clb/v1"
	listerv1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/client/lister/clb/v1"
	"bk-bcs/bcs-services/bcs-clb-controller/pkg/model"

	"fmt"

	"bk-bcs/bcs-common/common/blog"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

// KubeRegistry clb ingress kube registry
type KubeRegistry struct {
	clbName  string
	informer informerv1.ClbIngressInformer
	lister   listerv1.ClbIngressLister
	client   ingressClientV1.ClbV1Interface
}

// NewKubeRegistry create new registry for clb ingress
func NewKubeRegistry(clbname string, informer informerv1.ClbIngressInformer, lister listerv1.ClbIngressLister, client ingressClientV1.ClbV1Interface) (clbingress.Registry, error) {

	return &KubeRegistry{
		clbName:  clbname,
		informer: informer,
		lister:   lister,
		client:   client,
	}, nil
}

// AddIngressHandler add ingress handler
func (kr *KubeRegistry) AddIngressHandler(handler model.EventHandler) {
	kr.informer.Informer().AddEventHandler(handler)
}

// ListIngresses implements clbingress interface
func (kr *KubeRegistry) ListIngresses() ([]*ingressv1.ClbIngress, error) {

	// get ingresses for the certain clb name
	selector := labels.NewSelector()
	requirement, err := labels.NewRequirement("bmsf.tencent.com/clbname", selection.Equals, []string{kr.clbName})
	if err != nil {
		return nil, fmt.Errorf("create requirement failed, err %s", err.Error())
	}
	selector = selector.Add(*requirement)
	list, err := kr.informer.Lister().List(selector)
	if err != nil {
		return nil, err
	}
	for index, ingress := range list {
		blog.V(5).Infof("index: %d ingress for clb %s\n ingress: %v", index, kr.clbName, ingress)
	}
	// get ingresses for all clb
	requirementIngressForAll, err := labels.NewRequirement("bmsf.tencent.com/clbname", selection.Equals, []string{"all"})
	if err != nil {
		return nil, fmt.Errorf("create requirement of clb ingress for all clb failed, err %s", err.Error())
	}
	selectorIngressForAll := labels.NewSelector()
	selectorIngressForAll = selectorIngressForAll.Add(*requirementIngressForAll)
	listIngressForAll, err := kr.informer.Lister().List(selectorIngressForAll)
	if err != nil {
		return nil, err
	}
	for index, ingress := range listIngressForAll {
		blog.V(5).Infof("index: %d ingress for all clb\n ingress: %v", index, ingress)
	}

	list = append(list, listIngressForAll...)
	return list, nil
}

// GetIngress implement Registry interface
func (kr *KubeRegistry) GetIngress(name string) (*ingressv1.ClbIngress, error) {

	list, err := kr.ListIngresses()
	if err != nil {
		blog.Errorf("list ingresses failed, err %s", err.Error())
		return nil, fmt.Errorf("list ingresses failed, err %s", err.Error())
	}
	for _, ingress := range list {
		if ingress.GetName() == name {
			return ingress, nil
		}
	}
	return nil, fmt.Errorf("no found ingress with name %s", name)
}

// SetIngress implement Registry interface
func (kr *KubeRegistry) SetIngress(ingress *ingressv1.ClbIngress) error {
	// ingress with labels ("bmsf.tencent.com/clbname": "all") should be skip,
	// which has effects on multiple clb controller
	if ingress.GetLabels()["bmsf.tencent.com/clbname"] == "all" {
		return nil
	}

	old, err := kr.lister.ClbIngresses(ingress.GetNamespace()).Get(ingress.GetName())
	if err != nil {
		blog.Infof("get old ingress %s %s failed, err %s", ingress.GetNamespace(), ingress.GetName())
		return fmt.Errorf("get old ingress %s %s failed, err %s",
			ingress.GetNamespace(), ingress.GetName(), err.Error())
	}
	ingress.SetResourceVersion(old.GetResourceVersion())
	_, err = kr.client.ClbIngresses(ingress.GetNamespace()).Update(ingress)
	return err
}
