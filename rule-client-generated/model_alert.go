/*
 * Rules
 *
 * / Rules represents API that is responsible for gathering rules and their statuses.
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Alert struct {
	// / state returns the maximum state of alert instances for this rule.
	State int32 `json:"state,omitempty"`
	Name string `json:"name,omitempty"`
	Query string `json:"query,omitempty"`
	DurationSeconds float64 `json:"duration_seconds,omitempty"`
	Labels *ZLabelSet `json:"labels,omitempty"`
	Annotations *ZLabelSet `json:"annotations,omitempty"`
	Alerts []AlertInstance `json:"alerts,omitempty"`
	Health string `json:"health,omitempty"`
	LastError string `json:"last_error,omitempty"`
	EvaluationDurationSeconds float64 `json:"evaluation_duration_seconds,omitempty"`
	LastEvaluation string `json:"last_evaluation,omitempty"`
}
