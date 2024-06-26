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
 */

package gcp

const (
	// SystemNameInMetric system name in metric
	SystemNameInMetric = "gcp"
	// HandlerNameInMetricSDK handler name in metric
	HandlerNameInMetricSDK = "sdk"

	// ProtocolHTTP http protocol
	ProtocolHTTP = "HTTP"
	// ProtocolHTTPS https protocol
	ProtocolHTTPS = "HTTPS"
	// ProtocolTCP tcp protocol
	ProtocolTCP = "TCP"
	// ProtocolUDP udp protocol
	ProtocolUDP = "UDP"

	// AddressTypeInternal internal address
	AddressTypeInternal = "INTERNAL"
	// AddressTypeExternal external address
	AddressTypeExternal = "EXTERNAL"

	// AnnotationKeyLoadBalancerType service load-balancer-type annotation
	AnnotationKeyLoadBalancerType = "cloud.google.com/load-balancer-type"
	// AnnotationValueLoadBalancerTypeInternal internal value, need to be set when lb ip is internal
	AnnotationValueLoadBalancerTypeInternal = "Internal"
)
