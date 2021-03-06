/*
Copyright 2020 argoproj/gitops-engine Authors.

Modifications Copyright 2020 cakehappens/seaworthy Authors.

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

package v1beta1

import (
	"fmt"

	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/cakehappens/seaworthy/pkg/kubernetes/health"
)

const (
	// IngressKind for Ingress
	IngressKind = "Ingress"
)

// IngressHealth checks the health of an Ingress
func IngressHealth(obj unstructured.Unstructured) (health.Status, error) {
	ingress := &extv1beta1.Ingress{}
	err := scheme.Scheme.Convert(&obj, ingress, nil)
	if err != nil {
		err = fmt.Errorf("failed to convert %T to %T: %v", obj, ingress, err)
		return health.Status{
			Code:    health.Unknown,
			Message: err.Error(),
		}, err
	}

	if len(ingress.Status.LoadBalancer.Ingress) == 0 {
		return health.Status{
			Code:    health.Progressing,
			Message: "Working on it...",
		}, nil
	}

	return health.NewHealthyHealthStatus(), nil
}
