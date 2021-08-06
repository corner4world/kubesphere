/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/api/storage/v1alpha1"
)

// FakeStorageClassCapabilities implements StorageClassCapabilityInterface
type FakeStorageClassCapabilities struct {
	Fake *FakeStorageV1alpha1
}

var storageclasscapabilitiesResource = schema.GroupVersionResource{Group: "storage.kubesphere.io", Version: "v1alpha1", Resource: "storageclasscapabilities"}

var storageclasscapabilitiesKind = schema.GroupVersionKind{Group: "storage.kubesphere.io", Version: "v1alpha1", Kind: "StorageClassCapability"}

// Get takes name of the storageClassCapability, and returns the corresponding storageClassCapability object, and an error if there is any.
func (c *FakeStorageClassCapabilities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageClassCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storageclasscapabilitiesResource, name), &v1alpha1.StorageClassCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageClassCapability), err
}

// List takes label and field selectors, and returns the list of StorageClassCapabilities that match those selectors.
func (c *FakeStorageClassCapabilities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageClassCapabilityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storageclasscapabilitiesResource, storageclasscapabilitiesKind, opts), &v1alpha1.StorageClassCapabilityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageClassCapabilityList{ListMeta: obj.(*v1alpha1.StorageClassCapabilityList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageClassCapabilityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageClassCapabilities.
func (c *FakeStorageClassCapabilities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storageclasscapabilitiesResource, opts))
}

// Create takes the representation of a storageClassCapability and creates it.  Returns the server's representation of the storageClassCapability, and an error, if there is any.
func (c *FakeStorageClassCapabilities) Create(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.CreateOptions) (result *v1alpha1.StorageClassCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storageclasscapabilitiesResource, storageClassCapability), &v1alpha1.StorageClassCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageClassCapability), err
}

// Update takes the representation of a storageClassCapability and updates it. Returns the server's representation of the storageClassCapability, and an error, if there is any.
func (c *FakeStorageClassCapabilities) Update(ctx context.Context, storageClassCapability *v1alpha1.StorageClassCapability, opts v1.UpdateOptions) (result *v1alpha1.StorageClassCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storageclasscapabilitiesResource, storageClassCapability), &v1alpha1.StorageClassCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageClassCapability), err
}

// Delete takes name of the storageClassCapability and deletes it. Returns an error if one occurs.
func (c *FakeStorageClassCapabilities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storageclasscapabilitiesResource, name), &v1alpha1.StorageClassCapability{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageClassCapabilities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storageclasscapabilitiesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageClassCapabilityList{})
	return err
}

// Patch applies the patch and returns the patched storageClassCapability.
func (c *FakeStorageClassCapabilities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageClassCapability, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageclasscapabilitiesResource, name, pt, data, subresources...), &v1alpha1.StorageClassCapability{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageClassCapability), err
}
