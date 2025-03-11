// Copyright 2024 Google LLC All Rights Reserved.
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

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "agones.dev/agones/pkg/apis/autoscaling/v1"
	autoscalingv1 "agones.dev/agones/pkg/client/applyconfiguration/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFleetAutoscalers implements FleetAutoscalerInterface
type FakeFleetAutoscalers struct {
	Fake *FakeAutoscalingV1
	ns   string
}

var fleetautoscalersResource = v1.SchemeGroupVersion.WithResource("fleetautoscalers")

var fleetautoscalersKind = v1.SchemeGroupVersion.WithKind("FleetAutoscaler")

// Get takes name of the fleetAutoscaler, and returns the corresponding fleetAutoscaler object, and an error if there is any.
func (c *FakeFleetAutoscalers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.FleetAutoscaler, err error) {
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(fleetautoscalersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// List takes label and field selectors, and returns the list of FleetAutoscalers that match those selectors.
func (c *FakeFleetAutoscalers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FleetAutoscalerList, err error) {
	emptyResult := &v1.FleetAutoscalerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(fleetautoscalersResource, fleetautoscalersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.FleetAutoscalerList{ListMeta: obj.(*v1.FleetAutoscalerList).ListMeta}
	for _, item := range obj.(*v1.FleetAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fleetAutoscalers.
func (c *FakeFleetAutoscalers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(fleetautoscalersResource, c.ns, opts))

}

// Create takes the representation of a fleetAutoscaler and creates it.  Returns the server's representation of the fleetAutoscaler, and an error, if there is any.
func (c *FakeFleetAutoscalers) Create(ctx context.Context, fleetAutoscaler *v1.FleetAutoscaler, opts metav1.CreateOptions) (result *v1.FleetAutoscaler, err error) {
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(fleetautoscalersResource, c.ns, fleetAutoscaler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// Update takes the representation of a fleetAutoscaler and updates it. Returns the server's representation of the fleetAutoscaler, and an error, if there is any.
func (c *FakeFleetAutoscalers) Update(ctx context.Context, fleetAutoscaler *v1.FleetAutoscaler, opts metav1.UpdateOptions) (result *v1.FleetAutoscaler, err error) {
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(fleetautoscalersResource, c.ns, fleetAutoscaler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFleetAutoscalers) UpdateStatus(ctx context.Context, fleetAutoscaler *v1.FleetAutoscaler, opts metav1.UpdateOptions) (result *v1.FleetAutoscaler, err error) {
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(fleetautoscalersResource, "status", c.ns, fleetAutoscaler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// Delete takes name of the fleetAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeFleetAutoscalers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fleetautoscalersResource, c.ns, name, opts), &v1.FleetAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFleetAutoscalers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(fleetautoscalersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.FleetAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched fleetAutoscaler.
func (c *FakeFleetAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FleetAutoscaler, err error) {
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fleetautoscalersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fleetAutoscaler.
func (c *FakeFleetAutoscalers) Apply(ctx context.Context, fleetAutoscaler *autoscalingv1.FleetAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FleetAutoscaler, err error) {
	if fleetAutoscaler == nil {
		return nil, fmt.Errorf("fleetAutoscaler provided to Apply must not be nil")
	}
	data, err := json.Marshal(fleetAutoscaler)
	if err != nil {
		return nil, err
	}
	name := fleetAutoscaler.Name
	if name == nil {
		return nil, fmt.Errorf("fleetAutoscaler.Name must be provided to Apply")
	}
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fleetautoscalersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFleetAutoscalers) ApplyStatus(ctx context.Context, fleetAutoscaler *autoscalingv1.FleetAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FleetAutoscaler, err error) {
	if fleetAutoscaler == nil {
		return nil, fmt.Errorf("fleetAutoscaler provided to Apply must not be nil")
	}
	data, err := json.Marshal(fleetAutoscaler)
	if err != nil {
		return nil, err
	}
	name := fleetAutoscaler.Name
	if name == nil {
		return nil, fmt.Errorf("fleetAutoscaler.Name must be provided to Apply")
	}
	emptyResult := &v1.FleetAutoscaler{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(fleetautoscalersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.FleetAutoscaler), err
}
