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

func TestObjectBulkDelete(t *testing.T) {
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
	_, err := client.Objects.Bulk.Delete(
		context.TODO(),
		"collection",
		knock.ObjectBulkDeleteParams{
			ObjectIDs: knock.F([]string{"string"}),
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

func TestObjectBulkAddSubscriptions(t *testing.T) {
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
	_, err := client.Objects.Bulk.AddSubscriptions(
		context.TODO(),
		"projects",
		knock.ObjectBulkAddSubscriptionsParams{
			Subscriptions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscription{{
				ID: knock.F("project-1"),
				Recipients: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientUnion{knock.ObjectBulkAddSubscriptionsParamsSubscriptionsRecipientsInlineIdentifyUserRequest{
					ID: knock.F("user_1"),
					ChannelData: knock.F(map[string]shared.ChannelDataRequestParam{
						"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
							Data: knock.F[shared.ChannelDataRequestDataUnionParam](shared.PushChannelDataParam{
								Tokens: knock.F([]string{"push_token_xxx"}),
							}),
						},
					}),
					CreatedAt: knock.F(time.Now()),
					Preferences: knock.F(map[string]shared.PreferenceSetRequestParam{
						"default": {
							Categories: knock.F(map[string]shared.PreferenceSetRequestCategoriesUnionParam{
								"transactional": shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
									ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
										Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
										Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
										HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
										InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
										Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
										SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("some_property"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
							ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
								Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
								HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
							}),
							Workflows: knock.F(map[string]shared.PreferenceSetRequestWorkflowsUnionParam{
								"dinosaurs-loose": shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
									ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
										Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
										Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
										HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
										InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
										Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
										SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("some_property"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.property"),
									}}),
								},
							}),
						},
					}),
				}}),
				Properties: knock.F(map[string]interface{}{
					"foo": "bar",
				}),
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

func TestObjectBulkSet(t *testing.T) {
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
	_, err := client.Objects.Bulk.Set(
		context.TODO(),
		"collection",
		knock.ObjectBulkSetParams{
			Objects: knock.F([]knock.ObjectBulkSetParamsObject{{
				ID:         knock.F("project_1"),
				Collection: knock.F("projects"),
				ChannelData: knock.F(map[string]shared.ChannelDataRequestParam{
					"97c5837d-c65c-4d54-aa39-080eeb81c69d": {
						Data: knock.F[shared.ChannelDataRequestDataUnionParam](shared.PushChannelDataParam{
							Tokens: knock.F([]string{"push_token_xxx"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(map[string]shared.PreferenceSetRequestParam{
					"default": {
						Categories: knock.F(map[string]shared.PreferenceSetRequestCategoriesUnionParam{
							"transactional": shared.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
									Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
									HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]shared.ConditionParam{{
									Argument: knock.F("some_property"),
									Operator: knock.F(shared.ConditionOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
						ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
							Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
							Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
							HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
							InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
							Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
							SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
						}),
						Workflows: knock.F(map[string]shared.PreferenceSetRequestWorkflowsUnionParam{
							"dinosaurs-loose": shared.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(shared.PreferenceSetChannelTypesParam{
									Chat:      knock.F[shared.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[shared.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
									HTTP:      knock.F[shared.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[shared.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[shared.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[shared.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]shared.ConditionParam{{
									Argument: knock.F("some_property"),
									Operator: knock.F(shared.ConditionOperatorEqualTo),
									Variable: knock.F("recipient.property"),
								}}),
							},
						}),
					},
				}),
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
