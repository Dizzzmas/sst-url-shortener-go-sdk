// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssturlshortener_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/sst-url-shortener-go"
	"github.com/stainless-sdks/sst-url-shortener-go/internal/testutil"
	"github.com/stainless-sdks/sst-url-shortener-go/option"
)

func TestUsage(t *testing.T) {
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
	url, err := client.URLs.New(context.TODO(), ssturlshortener.URLNewParams{
		OriginalURL: ssturlshortener.F("REPLACE_ME"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", url.Result)
}
