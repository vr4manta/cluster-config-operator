// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOperators implements ClusterOperatorInterface
type FakeClusterOperators struct {
	Fake *FakeConfigV1
}

var clusteroperatorsResource = v1.SchemeGroupVersion.WithResource("clusteroperators")

var clusteroperatorsKind = v1.SchemeGroupVersion.WithKind("ClusterOperator")

// Get takes name of the clusterOperator, and returns the corresponding clusterOperator object, and an error if there is any.
func (c *FakeClusterOperators) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterOperator, err error) {
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clusteroperatorsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// List takes label and field selectors, and returns the list of ClusterOperators that match those selectors.
func (c *FakeClusterOperators) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterOperatorList, err error) {
	emptyResult := &v1.ClusterOperatorList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clusteroperatorsResource, clusteroperatorsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterOperatorList{ListMeta: obj.(*v1.ClusterOperatorList).ListMeta}
	for _, item := range obj.(*v1.ClusterOperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOperators.
func (c *FakeClusterOperators) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clusteroperatorsResource, opts))
}

// Create takes the representation of a clusterOperator and creates it.  Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Create(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.CreateOptions) (result *v1.ClusterOperator, err error) {
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clusteroperatorsResource, clusterOperator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// Update takes the representation of a clusterOperator and updates it. Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Update(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (result *v1.ClusterOperator, err error) {
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clusteroperatorsResource, clusterOperator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOperators) UpdateStatus(ctx context.Context, clusterOperator *v1.ClusterOperator, opts metav1.UpdateOptions) (result *v1.ClusterOperator, err error) {
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(clusteroperatorsResource, "status", clusterOperator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// Delete takes name of the clusterOperator and deletes it. Returns an error if one occurs.
func (c *FakeClusterOperators) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteroperatorsResource, name, opts), &v1.ClusterOperator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOperators) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clusteroperatorsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterOperatorList{})
	return err
}

// Patch applies the patch and returns the patched clusterOperator.
func (c *FakeClusterOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterOperator, err error) {
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteroperatorsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterOperator.
func (c *FakeClusterOperators) Apply(ctx context.Context, clusterOperator *configv1.ClusterOperatorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterOperator, err error) {
	if clusterOperator == nil {
		return nil, fmt.Errorf("clusterOperator provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterOperator)
	if err != nil {
		return nil, err
	}
	name := clusterOperator.Name
	if name == nil {
		return nil, fmt.Errorf("clusterOperator.Name must be provided to Apply")
	}
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteroperatorsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeClusterOperators) ApplyStatus(ctx context.Context, clusterOperator *configv1.ClusterOperatorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterOperator, err error) {
	if clusterOperator == nil {
		return nil, fmt.Errorf("clusterOperator provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterOperator)
	if err != nil {
		return nil, err
	}
	name := clusterOperator.Name
	if name == nil {
		return nil, fmt.Errorf("clusterOperator.Name must be provided to Apply")
	}
	emptyResult := &v1.ClusterOperator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteroperatorsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterOperator), err
}
