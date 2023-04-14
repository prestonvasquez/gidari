// Copyright 2023 The Gidari Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0

package gidari

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
	"google.golang.org/protobuf/types/known/structpb"
)

type mockServiceOptions struct {
	upsertStgCount int
	reqCount       int
	rateLimiter    *rate.Limiter
}

func newMockService(opts mockServiceOptions) *Service {
	svc, err := NewService(context.Background())
	if err != nil {
		panic(err)
	}

	reqs := newHTTPRequests(opts.reqCount)

	svc.HTTP.
		RateLimiter(opts.rateLimiter).
		Requests(reqs...)

	svc.HTTP.client = newMockHTTPClient(withMockHTTPClientRequests(reqs...))

	return svc
}

func newHTTPRequests(volume int) []*Request {
	requests := make([]*Request, volume)

	writer := newMockUpsertStorage()

	for i := 0; i < volume; i++ {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://example%d", i), nil)
		requests[i] = &Request{
			http:    req,
			writers: []ListWriter{writer},
		}
	}

	return requests
}

type mockHTTPClientResponseError struct {
	rsp *http.Response
	err error
}

type mockHTTPClient struct {
	mutex     sync.Mutex
	responses map[*http.Request]*mockHTTPClientResponseError
}

type mockHTTPClientOption func(*mockHTTPClient)

func newMockHTTPClient(opts ...mockHTTPClientOption) *mockHTTPClient {
	client := &mockHTTPClient{
		responses: make(map[*http.Request]*mockHTTPClientResponseError),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// withMockHTTPClientRequests will set the mockHTTPClient responses to the
// provided requests.
func withMockHTTPClientRequests(reqs ...*Request) mockHTTPClientOption {
	return func(client *mockHTTPClient) {
		for _, req := range reqs {
			body := io.NopCloser(bytes.NewBufferString(""))
			code := http.StatusOK

			rsp := &http.Response{
				Body:       body,
				StatusCode: code,
				Request:    req.http,
			}

			rspErr := &mockHTTPClientResponseError{rsp: rsp}

			// If the request has already been set, then just
			// update the response.
			if _, ok := client.responses[req.http]; ok {
				client.responses[req.http].rsp = rspErr.rsp

				continue
			}

			client.responses[req.http] = rspErr
		}
	}
}

func withMockHTTPClientResponseError(req *Request, err error) mockHTTPClientOption {
	return func(client *mockHTTPClient) {
		if req == nil {
			return
		}

		if err == nil {
			return
		}

		// If the request has already been set, then just
		// update the error.
		if _, ok := client.responses[req.http]; ok {
			client.responses[req.http].err = err

			return
		}

		client.responses[req.http] = &mockHTTPClientResponseError{
			err: err,
		}
	}
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	rsp := m.responses[req]

	// If the response has an error, return it.
	if rsp.err != nil {
		return nil, rsp.err
	}

	return rsp.rsp, nil
}

type mockUpsertWriter struct {
	count   int
	countMu sync.Mutex
}

func newMockUpsertStorage() *mockUpsertWriter {
	return &mockUpsertWriter{}
}

func (m *mockUpsertWriter) Write(context.Context, *structpb.ListValue) error {
	m.countMu.Lock()
	defer m.countMu.Unlock()

	m.count++

	return nil
}

type mockConn struct {
	readData [][]byte
	readIdx  int
}

func (c *mockConn) Read(b []byte) (int, error) {
	if c.readIdx >= len(c.readData) {
		return 0, io.EOF
	}

	n := copy(b, c.readData[c.readIdx])
	c.readIdx++

	return n, nil
}

func (c *mockConn) Write(b []byte) (int, error) {
	return len(b), nil
}

type mockListWriter struct {
	count   int
	countMu sync.Mutex
	data    [][]byte
}

func (m *mockListWriter) Write(_ context.Context, val *structpb.ListValue) error {
	m.countMu.Lock()
	defer m.countMu.Unlock()

	// Marshal the value to JSON.
	data, err := val.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	m.data = append(m.data, data)
	m.count++

	return nil
}

type mockConnBlocker struct{}

func (m *mockConnBlocker) Read(b []byte) (int, error) {
	ch := make(chan struct{})
	<-ch

	return 0, nil
}

func (m *mockConnBlocker) Write(b []byte) (int, error) {
	return len(b), nil
}
