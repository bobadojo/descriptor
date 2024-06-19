// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package gapic

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	storespb "github.com/bobadojo/go/stores/v1/storespb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newStoreLocatorClientHook clientHook

// StoreLocatorCallOptions contains the retry settings for each method of StoreLocatorClient.
type StoreLocatorCallOptions struct {
	FindStores []gax.CallOption
	GetStore []gax.CallOption
}

func defaultStoreLocatorGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint(":443"),
		internaloption.WithDefaultEndpointTemplate(":443"),
		internaloption.WithDefaultMTLSEndpoint(":443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultStoreLocatorCallOptions() *StoreLocatorCallOptions {
	return &StoreLocatorCallOptions{
		FindStores: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Aborted,
					codes.Canceled,
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    200 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetStore: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Aborted,
					codes.Canceled,
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    200 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalStoreLocatorClient is an interface that defines the methods available from .
type internalStoreLocatorClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	FindStores(context.Context, *storespb.FindStoresRequest, ...gax.CallOption) (*storespb.FindStoresResponse, error)
	GetStore(context.Context, *storespb.GetStoreRequest, ...gax.CallOption) (*storespb.Store, error)
}

// StoreLocatorClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A store locator API
type StoreLocatorClient struct {
	// The internal transport-dependent client.
	internalClient internalStoreLocatorClient

	// The call options for this service.
	CallOptions *StoreLocatorCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *StoreLocatorClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *StoreLocatorClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *StoreLocatorClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// FindStores returns a list of all stores in the store.
func (c *StoreLocatorClient) FindStores(ctx context.Context, req *storespb.FindStoresRequest, opts ...gax.CallOption) (*storespb.FindStoresResponse, error) {
	return c.internalClient.FindStores(ctx, req, opts...)
}

// GetStore returns a specific store.
func (c *StoreLocatorClient) GetStore(ctx context.Context, req *storespb.GetStoreRequest, opts ...gax.CallOption) (*storespb.Store, error) {
	return c.internalClient.GetStore(ctx, req, opts...)
}

// storeLocatorGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type storeLocatorGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing StoreLocatorClient
	CallOptions **StoreLocatorCallOptions

	// The gRPC API client.
	storeLocatorClient storespb.StoreLocatorClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewStoreLocatorClient creates a new store locator client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A store locator API
func NewStoreLocatorClient(ctx context.Context, opts ...option.ClientOption) (*StoreLocatorClient, error) {
	clientOpts := defaultStoreLocatorGRPCClientOptions()
	if newStoreLocatorClientHook != nil {
		hookOpts, err := newStoreLocatorClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := StoreLocatorClient{CallOptions: defaultStoreLocatorCallOptions()}

	c := &storeLocatorGRPCClient{
		connPool:    connPool,
		storeLocatorClient: storespb.NewStoreLocatorClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *storeLocatorGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *storeLocatorGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *storeLocatorGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *storeLocatorGRPCClient) FindStores(ctx context.Context, req *storespb.FindStoresRequest, opts ...gax.CallOption) (*storespb.FindStoresResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).FindStores[0:len((*c.CallOptions).FindStores):len((*c.CallOptions).FindStores)], opts...)
	var resp *storespb.FindStoresResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.storeLocatorClient.FindStores(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *storeLocatorGRPCClient) GetStore(ctx context.Context, req *storespb.GetStoreRequest, opts ...gax.CallOption) (*storespb.Store, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetStore[0:len((*c.CallOptions).GetStore):len((*c.CallOptions).GetStore)], opts...)
	var resp *storespb.Store
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.storeLocatorClient.GetStore(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
