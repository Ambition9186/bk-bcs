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
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "bk-bcs/bcs-services/bcs-network/bcs-cloudnetwork/pkg/apis/cloud/v1"
	"bk-bcs/bcs-services/bcs-network/bcs-cloudnetwork/pkg/client/internalclientset/scheme"

	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CloudV1Interface interface {
	RESTClient() rest.Interface
	NodeNetworksGetter
}

// CloudV1Client is used to interact with features provided by the cloud.bkbcs.tencent.com group.
type CloudV1Client struct {
	restClient rest.Interface
}

func (c *CloudV1Client) NodeNetworks(namespace string) NodeNetworkInterface {
	return newNodeNetworks(c, namespace)
}

// NewForConfig creates a new CloudV1Client for the given config.
func NewForConfig(c *rest.Config) (*CloudV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CloudV1Client{client}, nil
}

// NewForConfigOrDie creates a new CloudV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CloudV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CloudV1Client for the given RESTClient.
func New(c rest.Interface) *CloudV1Client {
	return &CloudV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CloudV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
