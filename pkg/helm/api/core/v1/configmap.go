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

var ConfigMapKind = resource.Kind{
	Group:   "core",
	Version: "v1",
	Kind:    "ConfigMap",
}

var ConfigMapResource = resource.Type{
	Kind: ConfigMapKind,
	Name: "configmaps",
}

func NewConfigMap(configMap *corev1.ConfigMap, client resource.Client) *ConfigMap {
	return &ConfigMap{
		Resource: resource.NewResource(configMap.ObjectMeta, ConfigMapKind, client),
		Object:   configMap,
	}
}

type ConfigMap struct {
	*resource.Resource
	Object *corev1.ConfigMap
}

func (r *ConfigMap) Delete() error {
	return r.Clientset().
		CoreV1().
		RESTClient().
		Delete().
		Namespace(r.Namespace).
		Resource(ConfigMapResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
