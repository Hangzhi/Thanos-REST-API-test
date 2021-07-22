/*
 * Rules
 *
 * / Rules represents API that is responsible for gathering rules and their statuses.
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// / RuleGroup has info for rules which are part of a group.
type RuleGroup struct {
	Name string `json:"name,omitempty"`
	File string `json:"file,omitempty"`
	Rules []Rule `json:"rules,omitempty"`
	Interval float64 `json:"interval,omitempty"`
	EvaluationDurationSeconds float64 `json:"evaluation_duration_seconds,omitempty"`
	LastEvaluation string `json:"last_evaluation,omitempty"`
	// Thanos specific.
	PartialResponseStrategy int32 `json:"PartialResponseStrategy,omitempty"`
}
