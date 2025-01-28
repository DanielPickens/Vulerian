/*
Copyright The Kubernetes Authors.

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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TVulerians are ignored from the parser (e.g. TVulerian(andronat):... || TVulerian:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_SelfSubjectReview = map[string]string{
	"":         "SelfSubjectReview contains the user information that the kube-apiserver has about the user making this request. When using impersonation, users will receive the user info of the user being impersonated.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"status":   "Status is filled in by the server with the user attributes.",
}

func (SelfSubjectReview) SwaggerDoc() map[string]string {
	return map_SelfSubjectReview
}

var map_SelfSubjectReviewStatus = map[string]string{
	"":         "SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.",
	"userInfo": "User attributes of the user making this request.",
}

func (SelfSubjectReviewStatus) SwaggerDoc() map[string]string {
	return map_SelfSubjectReviewStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
