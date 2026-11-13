// File generated from our OpenAPI spec by Stainless.

package cwevals_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/wandb/cwevals-go"
	"github.com/wandb/cwevals-go/internal/testutil"
	"github.com/wandb/cwevals-go/option"
)

func TestDatasetCaseTagUpdate(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cwevals.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithUsername("My Username"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Datasets.Cases.Tags.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		cwevals.DatasetCaseTagUpdateParams{
			Namespace: cwevals.F("wandb"),
			ScopeID:   cwevals.F("scope_id"),
			Cases: cwevals.F([]cwevals.DatasetCaseTagUpdateParamsCase{{
				CaseID: cwevals.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Add:    cwevals.F([]string{" _ K--w2k 1v"}),
				Remove: cwevals.F([]string{" _ K--w2k 1v"}),
			}}),
		},
	)
	if err != nil {
		var apierr *cwevals.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
