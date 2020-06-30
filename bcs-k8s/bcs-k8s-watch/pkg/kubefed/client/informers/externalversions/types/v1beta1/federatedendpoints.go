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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	types_v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	versioned "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/listers/types/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FederatedEndpointsInformer provides access to a shared informer and lister for
// FederatedEndpointses.
type FederatedEndpointsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FederatedEndpointsLister
}

type federatedEndpointsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedEndpointsInformer constructs a new informer for FederatedEndpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedEndpointsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedEndpointsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedEndpointsInformer constructs a new informer for FederatedEndpoints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedEndpointsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedEndpointses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedEndpointses(namespace).Watch(options)
			},
		},
		&types_v1beta1.FederatedEndpoints{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedEndpointsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedEndpointsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedEndpointsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&types_v1beta1.FederatedEndpoints{}, f.defaultInformer)
}

func (f *federatedEndpointsInformer) Lister() v1beta1.FederatedEndpointsLister {
	return v1beta1.NewFederatedEndpointsLister(f.Informer().GetIndexer())
}
