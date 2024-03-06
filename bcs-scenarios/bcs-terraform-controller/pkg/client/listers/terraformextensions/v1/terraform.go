/*
Copyright 2023.

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

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-scenarios/bcs-terraform-controller/pkg/apis/terraformextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TerraformLister helps list Terraforms.
// All objects returned here must be treated as read-only.
type TerraformLister interface {
	// List lists all Terraforms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Terraform, err error)
	// Terraforms returns an object that can list and get Terraforms.
	Terraforms(namespace string) TerraformNamespaceLister
	TerraformListerExpansion
}

// terraformLister implements the TerraformLister interface.
type terraformLister struct {
	indexer cache.Indexer
}

// NewTerraformLister returns a new TerraformLister.
func NewTerraformLister(indexer cache.Indexer) TerraformLister {
	return &terraformLister{indexer: indexer}
}

// List lists all Terraforms in the indexer.
func (s *terraformLister) List(selector labels.Selector) (ret []*v1.Terraform, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Terraform))
	})
	return ret, err
}

// Terraforms returns an object that can list and get Terraforms.
func (s *terraformLister) Terraforms(namespace string) TerraformNamespaceLister {
	return terraformNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TerraformNamespaceLister helps list and get Terraforms.
// All objects returned here must be treated as read-only.
type TerraformNamespaceLister interface {
	// List lists all Terraforms in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Terraform, err error)
	// Get retrieves the Terraform from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Terraform, error)
	TerraformNamespaceListerExpansion
}

// terraformNamespaceLister implements the TerraformNamespaceLister
// interface.
type terraformNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Terraforms in the indexer for a given namespace.
func (s terraformNamespaceLister) List(selector labels.Selector) (ret []*v1.Terraform, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Terraform))
	})
	return ret, err
}

// Get retrieves the Terraform from the indexer for a given namespace and name.
func (s terraformNamespaceLister) Get(name string) (*v1.Terraform, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("terraform"), name)
	}
	return obj.(*v1.Terraform), nil
}
