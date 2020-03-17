// Code generated by onit-generate. DO NOT EDIT.

package v1beta1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var CronJobKind = resource.Kind{
	Group:   "batch",
	Version: "v1beta1",
	Kind:    "CronJob",
	Scoped:  true,
}

var CronJobResource = resource.Type{
	Kind: CronJobKind,
	Name: "cronjobs",
}

func NewCronJob(cronJob *batchv1beta1.CronJob, client resource.Client) *CronJob {
	return &CronJob{
		Resource: resource.NewResource(cronJob.ObjectMeta, CronJobKind, client),
		Object:   cronJob,
	}
}

type CronJob struct {
	*resource.Resource
	Object *batchv1beta1.CronJob
}

func (r *CronJob) Delete() error {
	return r.Clientset().
		BatchV1beta1().
		RESTClient().
		Delete().
		Namespace(r.Namespace).
		Resource(CronJobResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
