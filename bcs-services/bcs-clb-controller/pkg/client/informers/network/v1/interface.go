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

package v1

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-services/bcs-clb-controller/pkg/client/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CloudListeners returns a CloudListenerInformer.
	CloudListeners() CloudListenerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CloudListeners returns a CloudListenerInformer.
func (v *version) CloudListeners() CloudListenerInformer {
	return &cloudListenerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
