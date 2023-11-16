package fga

import (
	"context"
	"os"
	"testing"

	"go.uber.org/zap"
)

const (
	modelFile     = "testdata/datum.json"
	defaultFGAURL = "localhost:8080"
)

func TestFGAClient(t *testing.T) {
	url := os.Getenv("TEST_FGA_URL")
	if url == "" {
		url = defaultFGAURL
	}

	// TODO: create test tuples and run through model
	_ = newTestFGAClient(t, url)
}

func newTestFGAClient(t testing.TB, url string) *Client {
	// create FGA client for test suites
	c, err := NewClient(url,
		WithScheme("http"),
		WithLogger(zap.NewNop().Sugar()))
	if err != nil {
		t.Fatal(err)
		return nil
	}

	// Create new store
	if _, err := c.CreateStore(context.Background(), "datum_test"); err != nil {
		t.Fatal(err)
	}

	// Create model
	if _, err := c.CreateModel(context.Background(), modelFile); err != nil {
		t.Fatal(err)
	}

	// for _, tk := range testTuples {
	// 	if err := c.WriteTuple(context.Background(), tk.TupleKey()); err != nil {
	// 		t.Fatal(err)
	// 	}
	// }
	return c
}
