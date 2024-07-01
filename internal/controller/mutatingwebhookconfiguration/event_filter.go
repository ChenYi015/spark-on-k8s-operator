/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mutatingwebhookconfiguration

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MutatingWebhookConfigurationEventFilter filters events for MutatingWebhookConfiguration.
type MutatingWebhookConfigurationEventFilter struct {
	name string
}

func NewMutatingWebhookConfigurationEventFilter(name string) *MutatingWebhookConfigurationEventFilter {
	return &MutatingWebhookConfigurationEventFilter{
		name: name,
	}
}

// MutatingWebhookConfigurationEventFilter implements predicate.Predicate.
var _ predicate.Predicate = &MutatingWebhookConfigurationEventFilter{}

// Create implements predicate.Predicate.
func (m *MutatingWebhookConfigurationEventFilter) Create(event.CreateEvent) bool {
	return true
}

// Update implements predicate.Predicate.
func (m *MutatingWebhookConfigurationEventFilter) Update(event.UpdateEvent) bool {
	return true
}

// Delete implements predicate.Predicate.
func (m *MutatingWebhookConfigurationEventFilter) Delete(event.DeleteEvent) bool {
	return false
}

// Generic implements predicate.Predicate.
func (m *MutatingWebhookConfigurationEventFilter) Generic(event.GenericEvent) bool {
	return true
}
