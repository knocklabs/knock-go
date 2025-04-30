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

func TestScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Schedules.New(context.TODO(), knock.ScheduleNewParams{
		Recipients: knock.F([]knock.RecipientRequestUnionParam{shared.UnionString("user_123")}),
		Repeats: knock.F([]knock.ScheduleRepeatRuleParam{{
			Typename:   knock.F("ScheduleRepeat"),
			Frequency:  knock.F(knock.ScheduleRepeatRuleFrequencyDaily),
			DayOfMonth: knock.Null[int64](),
			Days:       knock.F([]knock.ScheduleRepeatRuleDay{knock.ScheduleRepeatRuleDayMon, knock.ScheduleRepeatRuleDayTue, knock.ScheduleRepeatRuleDayWed, knock.ScheduleRepeatRuleDayThu, knock.ScheduleRepeatRuleDayFri, knock.ScheduleRepeatRuleDaySat, knock.ScheduleRepeatRuleDaySun}),
			Hours:      knock.Null[int64](),
			Interval:   knock.F(int64(1)),
			Minutes:    knock.Null[int64](),
		}}),
		Workflow: knock.F("comment-created"),
		Data: knock.F(map[string]interface{}{
			"key": "bar",
		}),
		EndingAt:    knock.Null[time.Time](),
		ScheduledAt: knock.Null[time.Time](),
		Tenant:      knock.F[knock.InlineTenantRequestUnionParam](shared.UnionString("acme_corp")),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Schedules.Update(context.TODO(), knock.ScheduleUpdateParams{
		ScheduleIDs: knock.F([]string{"123e4567-e89b-12d3-a456-426614174000"}),
		Actor:       knock.F[knock.RecipientReferenceUnionParam](shared.UnionString("string")),
		Data: knock.F(map[string]interface{}{
			"key": "bar",
		}),
		EndingAt: knock.Null[time.Time](),
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
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScheduleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Schedules.List(context.TODO(), knock.ScheduleListParams{
		Workflow:   knock.F("workflow"),
		After:      knock.F("after"),
		Before:     knock.F("before"),
		PageSize:   knock.F(int64(0)),
		Recipients: knock.F([]string{"string"}),
		Tenant:     knock.F("tenant"),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScheduleDelete(t *testing.T) {
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
	_, err := client.Schedules.Delete(context.TODO(), knock.ScheduleDeleteParams{
		ScheduleIDs: knock.F([]string{"123e4567-e89b-12d3-a456-426614174000"}),
	})
	if err != nil {
		var apierr *knock.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
