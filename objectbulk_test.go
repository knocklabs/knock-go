// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knock_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/knocklabs/knock-go"
	"github.com/knocklabs/knock-go/internal/testutil"
	"github.com/knocklabs/knock-go/option"
	"github.com/knocklabs/knock-go/shared"
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objects.Bulk.Delete(
		context.TODO(),
		"collection",
		knock.ObjectBulkDeleteParams{
			ObjectIDs: knock.F([]string{"obj_123", "obj_456", "obj_789"}),
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objects.Bulk.AddSubscriptions(
		context.TODO(),
		"projects",
		knock.ObjectBulkAddSubscriptionsParams{
			Subscriptions: knock.F([]knock.ObjectBulkAddSubscriptionsParamsSubscription{{
				Recipients: knock.F([]knock.RecipientRequestUnionParam{knock.InlineIdentifyUserRequestParam{
					ID: knock.F("user_1"),
					ChannelData: knock.F(knock.InlineChannelDataRequestParam{
						"97c5837d-c65c-4d54-aa39-080eeb81c69d": knock.ChannelDataRequestParam{
							Data: knock.F[knock.ChannelDataRequestDataUnionParam](knock.PushChannelDataParam{
								Tokens: knock.F([]string{"push_token_xxx"}),
							}),
						},
					}),
					CreatedAt: knock.F(time.Now()),
					Email:     knock.F("email"),
					Name:      knock.F("name"),
					Preferences: knock.F(knock.InlinePreferenceSetRequestParam{
						"default": knock.PreferenceSetRequestParam{
							Categories: knock.F(map[string]knock.PreferenceSetRequestCategoriesUnionParam{
								"transactional": knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
									ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
										Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
										Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
										HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
										Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
										SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.ConditionParam{{
										Argument: knock.F("frog_genome"),
										Operator: knock.F(knock.ConditionOperatorContains),
										Variable: knock.F("specimen.dna_sequence"),
									}}),
								},
							}),
							ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
								Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
								HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
							}),
							Workflows: knock.F(map[string]knock.PreferenceSetRequestWorkflowsUnionParam{
								"dinosaurs-loose": knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
									ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
										Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
										Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
										HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
										InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
										Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
										SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
									}),
									Conditions: knock.F([]knock.ConditionParam{{
										Argument: knock.F("frog_genome"),
										Operator: knock.F(knock.ConditionOperatorContains),
										Variable: knock.F("specimen.dna_sequence"),
									}}),
								},
								"welcome-sequence": shared.UnionBool(true),
							}),
						},
					}),
					Timezone: knock.F("timezone"),
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objects.Bulk.Set(
		context.TODO(),
		"collection",
		knock.ObjectBulkSetParams{
			Objects: knock.F([]knock.InlineObjectRequestParam{{
				ID:         knock.F("project_1"),
				Collection: knock.F("projects"),
				ChannelData: knock.F(knock.InlineChannelDataRequestParam{
					"97c5837d-c65c-4d54-aa39-080eeb81c69d": knock.ChannelDataRequestParam{
						Data: knock.F[knock.ChannelDataRequestDataUnionParam](knock.PushChannelDataParam{
							Tokens: knock.F([]string{"push_token_xxx"}),
						}),
					},
				}),
				CreatedAt: knock.F(time.Now()),
				Preferences: knock.F(knock.InlinePreferenceSetRequestParam{
					"default": knock.PreferenceSetRequestParam{
						Categories: knock.F(map[string]knock.PreferenceSetRequestCategoriesUnionParam{
							"transactional": knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
									Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
									HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ConditionParam{{
									Argument: knock.F("frog_genome"),
									Operator: knock.F(knock.ConditionOperatorContains),
									Variable: knock.F("specimen.dna_sequence"),
								}}),
							},
						}),
						ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
							Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
							Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
							HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
							InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
							Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
							SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
						}),
						Workflows: knock.F(map[string]knock.PreferenceSetRequestWorkflowsUnionParam{
							"dinosaurs-loose": knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
								ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
									Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
									Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
									HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
									InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
									Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
									SMS:       knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](shared.UnionBool(true)),
								}),
								Conditions: knock.F([]knock.ConditionParam{{
									Argument: knock.F("frog_genome"),
									Operator: knock.F(knock.ConditionOperatorContains),
									Variable: knock.F("specimen.dna_sequence"),
								}}),
							},
							"welcome-sequence": shared.UnionBool(true),
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
