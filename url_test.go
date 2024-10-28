// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssturlshortener_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Dizzzmas/sst-url-shortener-go-sdk"
	"github.com/Dizzzmas/sst-url-shortener-go-sdk/internal/testutil"
	"github.com/Dizzzmas/sst-url-shortener-go-sdk/option"
)

func TestURLNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.New(context.TODO(), ssturlshortener.URLNewParams{
		OriginalURL: ssturlshortener.F("https://example.com"),
		ExpiredAt:   ssturlshortener.F(time.Now()),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLDeleteByOriginalURL(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.DeleteByOriginalURL(context.TODO(), ssturlshortener.URLDeleteByOriginalURLParams{
		OriginalURL: ssturlshortener.F("https://example.com"),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLDeleteByShortID(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.DeleteByShortID(context.TODO(), ssturlshortener.URLDeleteByShortIDParams{
		ShortID: ssturlshortener.F("shortId"),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLFromOriginalURL(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.FromOriginalURL(context.TODO(), ssturlshortener.URLFromOriginalURLParams{
		OriginalURL: ssturlshortener.F("https://example.com"),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLFromShortID(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.FromShortID(context.TODO(), ssturlshortener.URLFromShortIDParams{
		ShortID: ssturlshortener.F("xxx"),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLQuickCount(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.QuickCount(context.TODO())
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLSearchWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.Search(context.TODO(), ssturlshortener.URLSearchParams{
		Cursor:                ssturlshortener.F("cursor"),
		ExpiredAtLte:          ssturlshortener.F(time.Now()),
		Limit:                 ssturlshortener.F(1.000000),
		OriginalURLBeginsWith: ssturlshortener.F("originalUrlBeginsWith"),
	})
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestURLSlowCount(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ssturlshortener.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.URLs.SlowCount(context.TODO())
	if err != nil {
		var apierr *ssturlshortener.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
