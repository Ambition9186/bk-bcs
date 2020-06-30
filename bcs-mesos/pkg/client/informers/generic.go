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

package informers

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=bkbcs.tencent.com, Version=v2
	case v2.SchemeGroupVersion.WithResource("admissionwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().AdmissionWebhookConfigurations().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("agents"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Agents().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("agentschedinfos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().AgentSchedInfos().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Applications().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcsclusteragentsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsClusterAgentSettings().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcscommandinfos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsCommandInfos().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcsconfigmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsConfigMaps().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcsendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsEndpoints().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcssecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsSecrets().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("bcsservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().BcsServices().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("crds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Crds().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("crrs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Crrs().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Deployments().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("frameworks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Frameworks().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("tasks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Tasks().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("taskgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().TaskGroups().Informer()}, nil
	case v2.SchemeGroupVersion.WithResource("versions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bkbcs().V2().Versions().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
