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

func TestUserBulkDelete(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Users.Bulk.Delete(context.TODO(), knock.UserBulkDeleteParams{
		UserIDs: knock.F([]string{"user_1", "user_2"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserBulkIdentify(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Users.Bulk.Identify(context.TODO(), knock.UserBulkIdentifyParams{
		Users: knock.F([]knock.InlineIdentifyUserRequestParam{{
			ID:     knock.F("user_1"),
			Avatar: knock.F("avatar"),
			ChannelData: knock.F(knock.InlineChannelDataRequestParam{
				"97c5837d-c65c-4d54-aa39-080eeb81c69d": knock.InlineChannelDataRequestItemPushChannelDataTokensOnlyParam{
					Tokens: knock.F([]string{"push_token_123"}),
				},
			}),
			CreatedAt:   knock.F(time.Now()),
			Email:       knock.F("jane@ingen.net"),
			Locale:      knock.F("locale"),
			Name:        knock.F("Jane Doe"),
			PhoneNumber: knock.F("phone_number"),
			Preferences: knock.F(knock.InlinePreferenceSetRequestParam{
				"default": knock.PreferenceSetRequestParam{
					PersistenceStrategy: knock.F(knock.PreferenceSetRequest_PersistenceStrategyMerge),
					Categories: knock.F(map[string]knock.PreferenceSetRequestCategoriesUnionParam{
						"marketing": shared.UnionBool(false),
						"transactional": knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
								Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
								HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("US"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.country_code"),
									}}),
								}),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("frog_genome"),
								Operator: knock.F(shared.ConditionOperatorContains),
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
						SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("US"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.country_code"),
							}}),
						}),
					}),
					CommercialSubscribed: knock.F(true),
					Workflows: knock.F(map[string]knock.PreferenceSetRequestWorkflowsUnionParam{
						"dinosaurs-loose": knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
							ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
								Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
								Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(true)),
								HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
								InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
								Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
								SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
									Conditions: knock.F([]shared.ConditionParam{{
										Argument: knock.F("US"),
										Operator: knock.F(shared.ConditionOperatorEqualTo),
										Variable: knock.F("recipient.country_code"),
									}}),
								}),
							}),
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("frog_genome"),
								Operator: knock.F(shared.ConditionOperatorContains),
								Variable: knock.F("specimen.dna_sequence"),
							}}),
						},
					}),
				},
			}),
			Timezone: knock.F("America/New_York"),
		}}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserBulkSetPreferencesWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Users.Bulk.SetPreferences(context.TODO(), knock.UserBulkSetPreferencesParams{
		Preferences: knock.F(knock.PreferenceSetRequestParam{
			PersistenceStrategy: knock.F(knock.PreferenceSetRequest_PersistenceStrategyMerge),
			Categories: knock.F(map[string]knock.PreferenceSetRequestCategoriesUnionParam{
				"marketing": shared.UnionBool(false),
				"transactional": knock.PreferenceSetRequestCategoriesPreferenceSetWorkflowCategorySettingObjectParam{
					ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
						Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
						Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
						HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
						Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
						SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("US"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.country_code"),
							}}),
						}),
					}),
					Conditions: knock.F([]shared.ConditionParam{{
						Argument: knock.F("frog_genome"),
						Operator: knock.F(shared.ConditionOperatorContains),
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
				SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
					Conditions: knock.F([]shared.ConditionParam{{
						Argument: knock.F("US"),
						Operator: knock.F(shared.ConditionOperatorEqualTo),
						Variable: knock.F("recipient.country_code"),
					}}),
				}),
			}),
			CommercialSubscribed: knock.F(true),
			Workflows: knock.F(map[string]knock.PreferenceSetRequestWorkflowsUnionParam{
				"dinosaurs-loose": knock.PreferenceSetRequestWorkflowsPreferenceSetWorkflowCategorySettingObjectParam{
					ChannelTypes: knock.F(knock.PreferenceSetChannelTypesParam{
						Chat:      knock.F[knock.PreferenceSetChannelTypesChatUnionParam](shared.UnionBool(true)),
						Email:     knock.F[knock.PreferenceSetChannelTypesEmailUnionParam](shared.UnionBool(false)),
						HTTP:      knock.F[knock.PreferenceSetChannelTypesHTTPUnionParam](shared.UnionBool(true)),
						InAppFeed: knock.F[knock.PreferenceSetChannelTypesInAppFeedUnionParam](shared.UnionBool(true)),
						Push:      knock.F[knock.PreferenceSetChannelTypesPushUnionParam](shared.UnionBool(true)),
						SMS: knock.F[knock.PreferenceSetChannelTypesSMSUnionParam](knock.PreferenceSetChannelTypeSettingParam{
							Conditions: knock.F([]shared.ConditionParam{{
								Argument: knock.F("US"),
								Operator: knock.F(shared.ConditionOperatorEqualTo),
								Variable: knock.F("recipient.country_code"),
							}}),
						}),
					}),
					Conditions: knock.F([]shared.ConditionParam{{
						Argument: knock.F("frog_genome"),
						Operator: knock.F(shared.ConditionOperatorContains),
						Variable: knock.F("specimen.dna_sequence"),
					}}),
				},
			}),
		}),
		UserIDs: knock.F([]string{"user_1", "user_2"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
