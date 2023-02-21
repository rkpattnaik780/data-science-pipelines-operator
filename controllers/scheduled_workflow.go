/*
Copyright 2023.

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

package controllers

import (
	dspipelinesiov1alpha1 "github.com/opendatahub-io/data-science-pipelines-operator/api/v1alpha1"
)

var scheduledWorkflowTemplates = []string{
	"scheduled-workflow/deployment.yaml.tmpl",
	"scheduled-workflow/role.yaml.tmpl",
	"scheduled-workflow/rolebinding.yaml.tmpl",
	"scheduled-workflow/sa.yaml.tmpl",
	"scheduled-workflow/role.yaml.tmpl",
	"scheduled-workflow/rolebinding.yaml.tmpl",
}

func (r *DSPipelineReconciler) ReconcileScheduledWorkflow(dsp *dspipelinesiov1alpha1.DSPipeline,
	params *DSPipelineParams) error {

	r.Log.Info("Applying ScheduledWorkflow ResourceRequirements")

	for _, template := range scheduledWorkflowTemplates {
		err := r.Apply(dsp, params, template)
		if err != nil {
			return err
		}
	}

	r.Log.Info("Finished applying ScheduledWorkflow ResourceRequirements")
	return nil
}
