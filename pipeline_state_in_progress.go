/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

type PipelineStateInProgress struct {

	Type_ string `json:"type"`

	// The name of pipeline state (IN_PROGRESS).
	Name string `json:"name,omitempty"`

	// A stage of an in progress state of a pipeline.
	Stage *PipelineStateInProgressStage `json:"stage,omitempty"`
}
