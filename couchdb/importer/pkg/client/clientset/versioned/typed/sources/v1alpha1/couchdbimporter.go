/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-contrib/couchdb/importer/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-contrib/couchdb/importer/pkg/client/clientset/versioned/scheme"
)

// CouchDbImportersGetter has a method to return a CouchDbImporterInterface.
// A group's client should implement this interface.
type CouchDbImportersGetter interface {
	CouchDbImporters(namespace string) CouchDbImporterInterface
}

// CouchDbImporterInterface has methods to work with CouchDbImporter resources.
type CouchDbImporterInterface interface {
	Create(*v1alpha1.CouchDbImporter) (*v1alpha1.CouchDbImporter, error)
	Update(*v1alpha1.CouchDbImporter) (*v1alpha1.CouchDbImporter, error)
	UpdateStatus(*v1alpha1.CouchDbImporter) (*v1alpha1.CouchDbImporter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CouchDbImporter, error)
	List(opts v1.ListOptions) (*v1alpha1.CouchDbImporterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CouchDbImporter, err error)
	CouchDbImporterExpansion
}

// couchDbImporters implements CouchDbImporterInterface
type couchDbImporters struct {
	client rest.Interface
	ns     string
}

// newCouchDbImporters returns a CouchDbImporters
func newCouchDbImporters(c *SourcesV1alpha1Client, namespace string) *couchDbImporters {
	return &couchDbImporters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the couchDbImporter, and returns the corresponding couchDbImporter object, and an error if there is any.
func (c *couchDbImporters) Get(name string, options v1.GetOptions) (result *v1alpha1.CouchDbImporter, err error) {
	result = &v1alpha1.CouchDbImporter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("couchdbimporters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CouchDbImporters that match those selectors.
func (c *couchDbImporters) List(opts v1.ListOptions) (result *v1alpha1.CouchDbImporterList, err error) {
	result = &v1alpha1.CouchDbImporterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("couchdbimporters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested couchDbImporters.
func (c *couchDbImporters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("couchdbimporters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a couchDbImporter and creates it.  Returns the server's representation of the couchDbImporter, and an error, if there is any.
func (c *couchDbImporters) Create(couchDbImporter *v1alpha1.CouchDbImporter) (result *v1alpha1.CouchDbImporter, err error) {
	result = &v1alpha1.CouchDbImporter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("couchdbimporters").
		Body(couchDbImporter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a couchDbImporter and updates it. Returns the server's representation of the couchDbImporter, and an error, if there is any.
func (c *couchDbImporters) Update(couchDbImporter *v1alpha1.CouchDbImporter) (result *v1alpha1.CouchDbImporter, err error) {
	result = &v1alpha1.CouchDbImporter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("couchdbimporters").
		Name(couchDbImporter.Name).
		Body(couchDbImporter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *couchDbImporters) UpdateStatus(couchDbImporter *v1alpha1.CouchDbImporter) (result *v1alpha1.CouchDbImporter, err error) {
	result = &v1alpha1.CouchDbImporter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("couchdbimporters").
		Name(couchDbImporter.Name).
		SubResource("status").
		Body(couchDbImporter).
		Do().
		Into(result)
	return
}

// Delete takes name of the couchDbImporter and deletes it. Returns an error if one occurs.
func (c *couchDbImporters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("couchdbimporters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *couchDbImporters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("couchdbimporters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched couchDbImporter.
func (c *couchDbImporters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CouchDbImporter, err error) {
	result = &v1alpha1.CouchDbImporter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("couchdbimporters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
