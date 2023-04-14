/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package validation

import (
	"context"
	"errors"

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v3/pkg/validation"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// validator for validating Kubermatic UserSSHKey CRD.
type validator struct{}

// NewValidator returns a new user SSH key validator.
func NewValidator() *validator {
	return &validator{}
}

var _ admission.CustomValidator = &validator{}

func (v *validator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	key, ok := obj.(*kubermaticv1.UserSSHKey)
	if !ok {
		return errors.New("object is not a UserSSHKey")
	}

	errs := validation.ValidateUserSSHKeyCreate(key)

	return errs.ToAggregate()
}

func (v *validator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	oldKey, ok := oldObj.(*kubermaticv1.UserSSHKey)
	if !ok {
		return errors.New("old object is not a UserSSHKey")
	}

	newKey, ok := newObj.(*kubermaticv1.UserSSHKey)
	if !ok {
		return errors.New("new object is not a UserSSHKey")
	}

	errs := validation.ValidateUserSSHKeyUpdate(oldKey, newKey)

	return errs.ToAggregate()
}

func (v *validator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
