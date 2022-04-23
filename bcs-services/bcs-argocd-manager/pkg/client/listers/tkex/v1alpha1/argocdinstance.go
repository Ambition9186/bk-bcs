/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-argocd-manager/pkg/apis/tkex/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArgocdInstanceLister helps list ArgocdInstances.
// All objects returned here must be treated as read-only.
type ArgocdInstanceLister interface {
	// List lists all ArgocdInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArgocdInstance, err error)
	// ArgocdInstances returns an object that can list and get ArgocdInstances.
	ArgocdInstances(namespace string) ArgocdInstanceNamespaceLister
	ArgocdInstanceListerExpansion
}

// argocdInstanceLister implements the ArgocdInstanceLister interface.
type argocdInstanceLister struct {
	indexer cache.Indexer
}

// NewArgocdInstanceLister returns a new ArgocdInstanceLister.
func NewArgocdInstanceLister(indexer cache.Indexer) ArgocdInstanceLister {
	return &argocdInstanceLister{indexer: indexer}
}

// List lists all ArgocdInstances in the indexer.
func (s *argocdInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ArgocdInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArgocdInstance))
	})
	return ret, err
}

// ArgocdInstances returns an object that can list and get ArgocdInstances.
func (s *argocdInstanceLister) ArgocdInstances(namespace string) ArgocdInstanceNamespaceLister {
	return argocdInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArgocdInstanceNamespaceLister helps list and get ArgocdInstances.
// All objects returned here must be treated as read-only.
type ArgocdInstanceNamespaceLister interface {
	// List lists all ArgocdInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArgocdInstance, err error)
	// Get retrieves the ArgocdInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ArgocdInstance, error)
	ArgocdInstanceNamespaceListerExpansion
}

// argocdInstanceNamespaceLister implements the ArgocdInstanceNamespaceLister
// interface.
type argocdInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArgocdInstances in the indexer for a given namespace.
func (s argocdInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ArgocdInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArgocdInstance))
	})
	return ret, err
}

// Get retrieves the ArgocdInstance from the indexer for a given namespace and name.
func (s argocdInstanceNamespaceLister) Get(name string) (*v1alpha1.ArgocdInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("argocdinstance"), name)
	}
	return obj.(*v1alpha1.ArgocdInstance), nil
}