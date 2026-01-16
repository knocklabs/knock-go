// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock

import (
	"github.com/knocklabs/knock-go/internal/apierror"
	"github.com/knocklabs/knock-go/shared"
)

type Error = apierror.Error

// A condition to be evaluated.
//
// This is an alias to an internal type.
type Condition = shared.Condition

// The operator to use in the condition evaluation.
//
// This is an alias to an internal type.
type ConditionOperator = shared.ConditionOperator

// This is an alias to an internal value.
const ConditionOperatorEqualTo = shared.ConditionOperatorEqualTo

// This is an alias to an internal value.
const ConditionOperatorNotEqualTo = shared.ConditionOperatorNotEqualTo

// This is an alias to an internal value.
const ConditionOperatorGreaterThan = shared.ConditionOperatorGreaterThan

// This is an alias to an internal value.
const ConditionOperatorLessThan = shared.ConditionOperatorLessThan

// This is an alias to an internal value.
const ConditionOperatorGreaterThanOrEqualTo = shared.ConditionOperatorGreaterThanOrEqualTo

// This is an alias to an internal value.
const ConditionOperatorLessThanOrEqualTo = shared.ConditionOperatorLessThanOrEqualTo

// This is an alias to an internal value.
const ConditionOperatorContains = shared.ConditionOperatorContains

// This is an alias to an internal value.
const ConditionOperatorNotContains = shared.ConditionOperatorNotContains

// This is an alias to an internal value.
const ConditionOperatorEmpty = shared.ConditionOperatorEmpty

// This is an alias to an internal value.
const ConditionOperatorNotEmpty = shared.ConditionOperatorNotEmpty

// This is an alias to an internal value.
const ConditionOperatorContainsAll = shared.ConditionOperatorContainsAll

// This is an alias to an internal value.
const ConditionOperatorIsTimestamp = shared.ConditionOperatorIsTimestamp

// This is an alias to an internal value.
const ConditionOperatorIsNotTimestamp = shared.ConditionOperatorIsNotTimestamp

// This is an alias to an internal value.
const ConditionOperatorIsTimestampOnOrAfter = shared.ConditionOperatorIsTimestampOnOrAfter

// This is an alias to an internal value.
const ConditionOperatorIsTimestampBefore = shared.ConditionOperatorIsTimestampBefore

// This is an alias to an internal value.
const ConditionOperatorIsTimestampBetween = shared.ConditionOperatorIsTimestampBetween

// This is an alias to an internal value.
const ConditionOperatorIsAudienceMember = shared.ConditionOperatorIsAudienceMember

// This is an alias to an internal value.
const ConditionOperatorIsNotAudienceMember = shared.ConditionOperatorIsNotAudienceMember

// A condition to be evaluated.
//
// This is an alias to an internal type.
type ConditionParam = shared.ConditionParam

// Pagination information for a list of resources.
//
// This is an alias to an internal type.
type PageInfo = shared.PageInfo
