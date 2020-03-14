// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var ServiceKind = resource.Kind{
	Group:   "core",
	Version: "v1",
	Kind:    "Service",
}

var ServiceResource = resource.Type{
	Kind: ServiceKind,
	Name: "services",
}

func NewService(service *corev1.Service, client resource.Client) *Service {
	return &Service{
		Resource:           resource.NewResource(service.ObjectMeta, ServiceKind, client),
		Object:             service,
		EndpointsReference: NewEndpointsReference(client, resource.NewUIDFilter(service.UID)),
	}
}

type Service struct {
	*resource.Resource
	Object *corev1.Service
	EndpointsReference
}

func (r *Service) Delete() error {
	return r.Clientset().
		CoreV1().
		RESTClient().
		Delete().
		Namespace(r.Namespace).
		Resource(ServiceResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}