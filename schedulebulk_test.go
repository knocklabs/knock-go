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

func TestScheduleBulkNew(t *testing.T) {
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
	_, err := client.Schedules.Bulk.New(context.TODO(), knock.ScheduleBulkNewParams{
		Schedules: knock.F([]knock.ScheduleBulkNewParamsSchedule{{
			Workflow: knock.F("comment-created"),
			Actor: knock.F[knock.RecipientRequestUnionParam](knock.InlineIdentifyUserRequestParam{
				ID: knock.F("user_1"),
				ChannelData: knock.F(knock.InlineChannelDataRequestParam{knock.InlineChannelDataRequestItemParam{
					ChannelID: knock.F("97c5837d-c65c-4d54-aa39-080eeb81c69d"),
					Data: knock.F[knock.InlineChannelDataRequestItemDataUnionParam](knock.PushChannelDataParam{
						Tokens:   knock.F([]string{"push_token_xxx"}),
						Type:     knock.F(knock.PushChannelDataTypePushFcm),
						Typename: knock.F(knock.PushChannelData_TypenamePushChannelData),
					}),
					Provider: knock.F("push_fcm"),
				}}),
				CreatedAt:   knock.F(time.Now()),
				Preferences: knock.F[any](map[string]interface{}{}),
			}),
			Data: knock.F(map[string]interface{}{
				"key": "bar",
			}),
			EndingAt:  knock.Null[time.Time](),
			Recipient: knock.F[knock.RecipientRequestUnionParam](shared.UnionString("dnedry")),
			Repeats: knock.F([]knock.ScheduleRepeatRuleParam{{
				Typename:   knock.F("ScheduleRepeat"),
				Frequency:  knock.F(knock.ScheduleRepeatRuleFrequencyDaily),
				DayOfMonth: knock.Null[int64](),
				Days:       knock.F([]knock.ScheduleRepeatRuleDay{knock.ScheduleRepeatRuleDayMon, knock.ScheduleRepeatRuleDayTue, knock.ScheduleRepeatRuleDayWed, knock.ScheduleRepeatRuleDayThu, knock.ScheduleRepeatRuleDayFri, knock.ScheduleRepeatRuleDaySat, knock.ScheduleRepeatRuleDaySun}),
				Hours:      knock.Null[int64](),
				Interval:   knock.F(int64(1)),
				Minutes:    knock.Null[int64](),
			}}),
			ScheduledAt: knock.Null[time.Time](),
			Tenant:      knock.F[knock.InlineTenantRequestUnionParam](shared.UnionString("acme_corp")),
		}, {
			Workflow: knock.F("comment-created"),
			Actor: knock.F[knock.RecipientRequestUnionParam](knock.InlineIdentifyUserRequestParam{
				ID: knock.F("user_1"),
				ChannelData: knock.F(knock.InlineChannelDataRequestParam{knock.InlineChannelDataRequestItemParam{
					ChannelID: knock.F("97c5837d-c65c-4d54-aa39-080eeb81c69d"),
					Data: knock.F[knock.InlineChannelDataRequestItemDataUnionParam](knock.PushChannelDataParam{
						Tokens:   knock.F([]string{"push_token_xxx"}),
						Type:     knock.F(knock.PushChannelDataTypePushFcm),
						Typename: knock.F(knock.PushChannelData_TypenamePushChannelData),
					}),
					Provider: knock.F("push_fcm"),
				}}),
				CreatedAt:   knock.F(time.Now()),
				Preferences: knock.F[any](map[string]interface{}{}),
			}),
			Data: knock.F(map[string]interface{}{
				"key": "bar",
			}),
			EndingAt:  knock.Null[time.Time](),
			Recipient: knock.F[knock.RecipientRequestUnionParam](shared.UnionString("esattler")),
			Repeats: knock.F([]knock.ScheduleRepeatRuleParam{{
				Typename:   knock.F("ScheduleRepeat"),
				Frequency:  knock.F(knock.ScheduleRepeatRuleFrequencyDaily),
				DayOfMonth: knock.Null[int64](),
				Days:       knock.F([]knock.ScheduleRepeatRuleDay{knock.ScheduleRepeatRuleDayMon, knock.ScheduleRepeatRuleDayTue, knock.ScheduleRepeatRuleDayWed, knock.ScheduleRepeatRuleDayThu, knock.ScheduleRepeatRuleDayFri, knock.ScheduleRepeatRuleDaySat, knock.ScheduleRepeatRuleDaySun}),
				Hours:      knock.Null[int64](),
				Interval:   knock.F(int64(1)),
				Minutes:    knock.Null[int64](),
			}}),
			ScheduledAt: knock.Null[time.Time](),
			Tenant:      knock.F[knock.InlineTenantRequestUnionParam](shared.UnionString("acme_corp")),
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
