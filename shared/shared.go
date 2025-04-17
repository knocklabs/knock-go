// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/knock-go/internal/apijson"
	"github.com/stainless-sdks/knock-go/internal/param"
)

// A condition to be evaluated.
type Condition struct {
	// The argument value to compare against in the condition.
	Argument string `json:"argument,required,nullable"`
	// The operator to use in the condition evaluation.
	Operator ConditionOperator `json:"operator,required"`
	// The variable to be evaluated in the condition.
	Variable string        `json:"variable,required"`
	JSON     conditionJSON `json:"-"`
}

// conditionJSON contains the JSON metadata for the struct [Condition]
type conditionJSON struct {
	Argument    apijson.Field
	Operator    apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Condition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conditionJSON) RawJSON() string {
	return r.raw
}

// The operator to use in the condition evaluation.
type ConditionOperator string

const (
	ConditionOperatorEqualTo              ConditionOperator = "equal_to"
	ConditionOperatorNotEqualTo           ConditionOperator = "not_equal_to"
	ConditionOperatorGreaterThan          ConditionOperator = "greater_than"
	ConditionOperatorLessThan             ConditionOperator = "less_than"
	ConditionOperatorGreaterThanOrEqualTo ConditionOperator = "greater_than_or_equal_to"
	ConditionOperatorLessThanOrEqualTo    ConditionOperator = "less_than_or_equal_to"
	ConditionOperatorContains             ConditionOperator = "contains"
	ConditionOperatorNotContains          ConditionOperator = "not_contains"
	ConditionOperatorEmpty                ConditionOperator = "empty"
	ConditionOperatorNotEmpty             ConditionOperator = "not_empty"
	ConditionOperatorContainsAll          ConditionOperator = "contains_all"
	ConditionOperatorIsTimestamp          ConditionOperator = "is_timestamp"
	ConditionOperatorIsNotTimestamp       ConditionOperator = "is_not_timestamp"
	ConditionOperatorIsTimestampAfter     ConditionOperator = "is_timestamp_after"
	ConditionOperatorIsTimestampBefore    ConditionOperator = "is_timestamp_before"
	ConditionOperatorIsTimestampBetween   ConditionOperator = "is_timestamp_between"
	ConditionOperatorIsAudienceMember     ConditionOperator = "is_audience_member"
	ConditionOperatorIsNotAudienceMember  ConditionOperator = "is_not_audience_member"
)

func (r ConditionOperator) IsKnown() bool {
	switch r {
	case ConditionOperatorEqualTo, ConditionOperatorNotEqualTo, ConditionOperatorGreaterThan, ConditionOperatorLessThan, ConditionOperatorGreaterThanOrEqualTo, ConditionOperatorLessThanOrEqualTo, ConditionOperatorContains, ConditionOperatorNotContains, ConditionOperatorEmpty, ConditionOperatorNotEmpty, ConditionOperatorContainsAll, ConditionOperatorIsTimestamp, ConditionOperatorIsNotTimestamp, ConditionOperatorIsTimestampAfter, ConditionOperatorIsTimestampBefore, ConditionOperatorIsTimestampBetween, ConditionOperatorIsAudienceMember, ConditionOperatorIsNotAudienceMember:
		return true
	}
	return false
}

// A condition to be evaluated.
type ConditionParam struct {
	// The argument value to compare against in the condition.
	Argument param.Field[string] `json:"argument,required"`
	// The operator to use in the condition evaluation.
	Operator param.Field[ConditionOperator] `json:"operator,required"`
	// The variable to be evaluated in the condition.
	Variable param.Field[string] `json:"variable,required"`
}

func (r ConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
