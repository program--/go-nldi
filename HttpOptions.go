package nldi

import (
	"context"

	"github.com/program--/nldi/internal"
)

// `client` is defined by the following interface:
//
//	type HttpClient interface {
//		Do(*http.Request) (*http.Response, error)
//	}
//
// This interface is not exported due to the
// internal http client.
func WithClient(client internal.HttpClient) {
	internal.NLDIHttpClient.WithClient(client)
}

func WithContext(ctx context.Context) {
	internal.NLDIHttpClient.WithContext(ctx)
}
