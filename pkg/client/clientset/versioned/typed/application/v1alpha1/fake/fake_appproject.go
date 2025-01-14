// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppProjects implements AppProjectInterface
type FakeAppProjects struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var appprojectsResource = v1alpha1.SchemeGroupVersion.WithResource("appprojects")

var appprojectsKind = v1alpha1.SchemeGroupVersion.WithKind("AppProject")

// Get takes name of the appProject, and returns the corresponding appProject object, and an error if there is any.
func (c *FakeAppProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppProject, err error) {
	emptyResult := &v1alpha1.AppProject{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(appprojectsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// List takes label and field selectors, and returns the list of AppProjects that match those selectors.
func (c *FakeAppProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppProjectList, err error) {
	emptyResult := &v1alpha1.AppProjectList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(appprojectsResource, appprojectsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppProjectList{ListMeta: obj.(*v1alpha1.AppProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appProjects.
func (c *FakeAppProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(appprojectsResource, c.ns, opts))

}

// Create takes the representation of a appProject and creates it.  Returns the server's representation of the appProject, and an error, if there is any.
func (c *FakeAppProjects) Create(ctx context.Context, appProject *v1alpha1.AppProject, opts v1.CreateOptions) (result *v1alpha1.AppProject, err error) {
	emptyResult := &v1alpha1.AppProject{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(appprojectsResource, c.ns, appProject, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// Update takes the representation of a appProject and updates it. Returns the server's representation of the appProject, and an error, if there is any.
func (c *FakeAppProjects) Update(ctx context.Context, appProject *v1alpha1.AppProject, opts v1.UpdateOptions) (result *v1alpha1.AppProject, err error) {
	emptyResult := &v1alpha1.AppProject{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(appprojectsResource, c.ns, appProject, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AppProject), err
}

// Delete takes name of the appProject and deletes it. Returns an error if one occurs.
func (c *FakeAppProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(appprojectsResource, c.ns, name, opts), &v1alpha1.AppProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(appprojectsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppProjectList{})
	return err
}

// Patch applies the patch and returns the patched appProject.
func (c *FakeAppProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppProject, err error) {
	emptyResult := &v1alpha1.AppProject{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(appprojectsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AppProject), err
}
