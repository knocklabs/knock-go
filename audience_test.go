// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/knock-go"
	"github.com/stainless-sdks/knock-go/internal/testutil"
	"github.com/stainless-sdks/knock-go/option"
	"github.com/stainless-sdks/knock-go/shared"
)

func TestAudienceAddMembers(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Audiences.AddMembers(
		context.TODO(),
		"key",
		knock.AudienceAddMembersParams{
			Members: knock.F([]knock.AudienceAddMembersParamsMember{{
				User: knock.F(knock.AudienceAddMembersParamsMembersUser{
					ID: knock.F("user_1"),
					ChannelData: knock.F(map[string]knock.AudienceAddMembersParamsMembersUserChannelData{
						"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
							Data: knock.F[knock.AudienceAddMembersParamsMembersUserChannelDataDataUnion](knock.AudienceAddMembersParamsMembersUserChannelDataDataPushChannelData{
								Tokens: knock.F([]string{"push_token_xxx"}),
							}),
						},
					}),
					CreatedAt: knock.F(time.Now()),
					Preferences: knock.F(map[string]knock.AudienceAddMembersParamsMembersUserPreferences{
						"default": {
							Categories: knock.F(map[string]knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesUnion{
								"transactional": knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
										HTTP:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.AudienceAddMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
							ChannelTypes: knock.F(knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypes{
								Chat:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Workflows: knock.F(map[string]knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsUnion{
								"dinosaurs-loose": knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
										HTTP:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.AudienceAddMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
						},
					}),
				}),
				Tenant: knock.Null[string](),
			}}),
		},
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAudienceListMembers(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Audiences.ListMembers(context.TODO(), "key")
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAudienceRemoveMembers(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knock.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Audiences.RemoveMembers(
		context.TODO(),
		"key",
		knock.AudienceRemoveMembersParams{
			Members: knock.F([]knock.AudienceRemoveMembersParamsMember{{
				User: knock.F(knock.AudienceRemoveMembersParamsMembersUser{
					ID: knock.F("user_1"),
					ChannelData: knock.F(map[string]knock.AudienceRemoveMembersParamsMembersUserChannelData{
						"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
							Data: knock.F[knock.AudienceRemoveMembersParamsMembersUserChannelDataDataUnion](knock.AudienceRemoveMembersParamsMembersUserChannelDataDataPushChannelData{
								Tokens: knock.F([]string{"push_token_xxx"}),
							}),
						},
					}),
					CreatedAt: knock.F(time.Now()),
					Preferences: knock.F(map[string]knock.AudienceRemoveMembersParamsMembersUserPreferences{
						"default": {
							Categories: knock.F(map[string]knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesUnion{
								"transactional": knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(false)),
										HTTP:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.AudienceRemoveMembersParamsMembersUserPreferencesCategoriesPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
							ChannelTypes: knock.F(knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypes{
								Chat:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesChatUnion](shared.UnionBool(true)),
								Email:     knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesEmailUnion](shared.UnionBool(true)),
								HTTP:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesHTTPUnion](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesInAppFeedUnion](shared.UnionBool(true)),
								Push:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesPushUnion](shared.UnionBool(true)),
								SMS:       knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesChannelTypesSMSUnion](shared.UnionBool(true)),
							}),
							Workflows: knock.F(map[string]knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsUnion{
								"dinosaurs-loose": knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObject{
									ChannelTypes: knock.F(knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypes{
										Chat:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesChatUnion](shared.UnionBool(true)),
										Email:     knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesEmailUnion](shared.UnionBool(true)),
										HTTP:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesHTTPUnion](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesInAppFeedUnion](shared.UnionBool(true)),
										Push:      knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesPushUnion](shared.UnionBool(true)),
										SMS:       knock.F[knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectChannelTypesSMSUnion](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectCondition{{
										Argument: knock.F("some_property"),
										Operator: knock.F(knock.AudienceRemoveMembersParamsMembersUserPreferencesWorkflowsPreferenceSetWorkflowCategorySettingObjectConditionsOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
						},
					}),
				}),
				Tenant: knock.Null[string](),
			}}),
		},
	)
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
