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
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var JobKind = resource.Kind{
	Group:   "batch",
	Version: "v1",
	Kind:    "Job",
}

var JobResource = resource.Type{
	Kind: JobKind,
	Name: "jobs",
}

func NewJob(job *batchv1.Job, client resource.Client) *Job {
	return &Job{
		Resource: resource.NewResource(job.ObjectMeta, JobKind, client),
		Object:   job,
	}
}

type Job struct {
	*resource.Resource
	Object *batchv1.Job
}

func (r *Job) Delete() error {
	return r.Clientset().
		BatchV1().
		RESTClient().
		Delete().
		Namespace(r.Namespace).
		Resource(JobResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
