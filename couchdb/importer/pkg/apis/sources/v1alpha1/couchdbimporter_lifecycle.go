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

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	"knative.dev/eventing/pkg/apis/duck"
	"knative.dev/pkg/apis"
)

const (
	// CouchDbConditionReady has status True when the CouchDbImporter is ready to send events.
	CouchDbConditionReady = apis.ConditionReady

	// CouchDbConditionSinkProvided has status True when the CouchDbImporter has been configured with a sink target.
	CouchDbConditionSinkProvided apis.ConditionType = "SinkProvided"

	// CouchDbConditionDeployed has status True when the CouchDbImporter has had it's deployment created.
	CouchDbConditionDeployed apis.ConditionType = "Deployed"

	// CouchDbConditionEventTypeProvided has status True when the CouchDbImporter has been configured with its event types.
	CouchDbConditionEventTypeProvided apis.ConditionType = "EventTypesProvided"
)

var CouchDbCondSet = apis.NewLivingConditionSet(
	CouchDbConditionSinkProvided,
	CouchDbConditionDeployed,
)

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *CouchDbImporterStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return CouchDbCondSet.Manage(s).GetCondition(t)
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *CouchDbImporterStatus) InitializeConditions() {
	CouchDbCondSet.Manage(s).InitializeConditions()
}

// MarkSink sets the condition that the source has a sink configured.
func (s *CouchDbImporterStatus) MarkSink(uri string) {
	s.SinkURI = uri
	if len(uri) > 0 {
		CouchDbCondSet.Manage(s).MarkTrue(CouchDbConditionSinkProvided)
	} else {
		CouchDbCondSet.Manage(s).MarkUnknown(CouchDbConditionSinkProvided, "SinkEmpty", "Sink has resolved to empty.%s", "")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (s *CouchDbImporterStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	CouchDbCondSet.Manage(s).MarkFalse(CouchDbConditionSinkProvided, reason, messageFormat, messageA...)
}

// PropagateDeploymentAvailability uses the availability of the provided Deployment to determine if
// CouchDbConditionDeployed should be marked as true or false.
func (s *CouchDbImporterStatus) PropagateDeploymentAvailability(d *appsv1.Deployment) {
	if duck.DeploymentIsAvailable(&d.Status, false) {
		CouchDbCondSet.Manage(s).MarkTrue(CouchDbConditionDeployed)
	} else {
		// I don't know how to propagate the status well, so just give the name of the Deployment
		// for now.
		CouchDbCondSet.Manage(s).MarkFalse(CouchDbConditionDeployed, "DeploymentUnavailable", "The Deployment '%s' is unavailable.", d.Name)
	}
}

// MarkEventTypes sets the condition that the source has set its event type.
func (s *CouchDbImporterStatus) MarkEventTypes() {
	CouchDbCondSet.Manage(s).MarkTrue(CouchDbConditionEventTypeProvided)
}

// MarkNoEventTypes sets the condition that the source does not its event type configured.
func (s *CouchDbImporterStatus) MarkNoEventTypes(reason, messageFormat string, messageA ...interface{}) {
	CouchDbCondSet.Manage(s).MarkFalse(CouchDbConditionEventTypeProvided, reason, messageFormat, messageA...)
}

// IsReady returns true if the resource is ready overall.
func (s *CouchDbImporterStatus) IsReady() bool {
	return CouchDbCondSet.Manage(s).IsHappy()
}
