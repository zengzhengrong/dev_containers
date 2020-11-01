// Code generated by protoc-gen-goclay. DO NOT EDIT.
// source: sum.proto

/*
Package sum is a self-registering gRPC and JSON+Swagger service definition.

It conforms to the github.com/utrack/clay/v2/transport Service interface.
*/
package sum

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
	"github.com/utrack/clay/v2/transport/httpclient"
	"github.com/utrack/clay/v2/transport/httpruntime"
	"github.com/utrack/clay/v2/transport/httpruntime/httpmw"
	"github.com/utrack/clay/v2/transport/httptransport"
	"github.com/utrack/clay/v2/transport/swagger"
	"google.golang.org/grpc"
)

// Update your shared lib or downgrade generator to v1 if there's an error
var _ = transport.IsVersion2

var _ = ioutil.Discard
var _ chi.Router
var _ runtime.Marshaler
var _ bytes.Buffer
var _ context.Context
var _ fmt.Formatter
var _ strings.Reader
var _ errors.Frame
var _ httpruntime.Marshaler
var _ http.Handler
var _ url.Values
var _ base64.Encoding
var _ httptransport.MarshalerError
var _ utilities.DoubleArray

// SummatorDesc is a descriptor/registrator for the SummatorServer.
type SummatorDesc struct {
	svc  SummatorServer
	opts httptransport.DescOptions
}

// NewSummatorServiceDesc creates new registrator for the SummatorServer.
// It implements httptransport.ConfigurableServiceDesc as well.
func NewSummatorServiceDesc(svc SummatorServer) *SummatorDesc {
	return &SummatorDesc{
		svc: svc,
	}
}

// RegisterGRPC implements service registrator interface.
func (d *SummatorDesc) RegisterGRPC(s *grpc.Server) {
	RegisterSummatorServer(s, d.svc)
}

// Apply applies passed options.
func (d *SummatorDesc) Apply(oo ...transport.DescOption) {
	for _, o := range oo {
		o.Apply(&d.opts)
	}
}

// SwaggerDef returns this file's Swagger definition.
func (d *SummatorDesc) SwaggerDef(options ...swagger.Option) (result []byte) {
	if len(options) > 0 || len(d.opts.SwaggerDefaultOpts) > 0 {
		var err error
		var s = &spec.Swagger{}
		if err = s.UnmarshalJSON(_swaggerDef_sum_proto); err != nil {
			panic("Bad swagger definition: " + err.Error())
		}

		for _, o := range d.opts.SwaggerDefaultOpts {
			o(s)
		}
		for _, o := range options {
			o(s)
		}
		if result, err = s.MarshalJSON(); err != nil {
			panic("Failed marshal spec.Swagger definition: " + err.Error())
		}
	} else {
		result = _swaggerDef_sum_proto
	}
	return result
}

// RegisterHTTP registers this service's HTTP handlers/bindings.
func (d *SummatorDesc) RegisterHTTP(mux transport.Router) {
	chiMux, isChi := mux.(chi.Router)

	{
		// Handler for Sum, binding: POST /v1/example/sum/{a}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Summator_Sum_0(r)
			rsp, err := _Summator_Sum_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Summator_Sum_0, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}
	{
		// Handler for Sum, binding: POST /v1/example/sum/{a}
		var h http.HandlerFunc
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			unmFunc := unmarshaler_goclay_Summator_Sum_1(r)
			rsp, err := _Summator_Sum_Handler(d.svc, r.Context(), unmFunc, d.opts.UnaryInterceptor)

			if err != nil {
				if err, ok := err.(httptransport.MarshalerError); ok {
					httpruntime.SetError(r.Context(), r, w, errors.Wrap(err.Err, "couldn't parse request"))
					return
				}
				httpruntime.SetError(r.Context(), r, w, err)
				return
			}

			if ctxErr := r.Context().Err(); ctxErr != nil && ctxErr == context.Canceled {
				w.WriteHeader(499) // Client Closed Request
				return
			}

			_, outbound := httpruntime.MarshalerForRequest(r)
			w.Header().Set("Content-Type", outbound.ContentType())
			err = outbound.Marshal(w, rsp)
			if err != nil {
				httpruntime.SetError(r.Context(), r, w, errors.Wrap(err, "couldn't write response"))
				return
			}
		})

		h = httpmw.DefaultChain(h)

		if isChi {
			chiMux.Method("POST", pattern_goclay_Summator_Sum_1, h)
		} else {
			panic("query URI params supported only for chi.Router")
		}
	}

}

type Summator_httpClient struct {
	c    *http.Client
	host string
}

// NewSummatorHTTPClient creates new HTTP client for SummatorServer.
// Pass addr in format "http://host[:port]".
func NewSummatorHTTPClient(c *http.Client, addr string) *Summator_httpClient {
	if strings.HasSuffix(addr, "/") {
		addr = addr[:len(addr)-1]
	}
	return &Summator_httpClient{c: c, host: addr}
}

func (c *Summator_httpClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	mw, err := httpclient.NewMiddlewareGRPC(opts)
	if err != nil {
		return nil, err
	}

	path := pattern_goclay_Summator_Sum_0_builder(in)

	buf := bytes.NewBuffer(nil)

	m := httpruntime.DefaultMarshaler(nil)

	if err = m.Marshal(buf, in.B); err != nil {
		return nil, errors.Wrap(err, "can't marshal request")
	}

	req, err := http.NewRequest("POST", c.host+path, buf)
	if err != nil {
		return nil, errors.Wrap(err, "can't initiate HTTP request")
	}
	req = req.WithContext(ctx)

	req.Header.Add("Accept", m.ContentType())

	req, err = mw.ProcessRequest(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error from client")
	}
	defer rsp.Body.Close()

	rsp, err = mw.ProcessResponse(rsp)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.Errorf("%v %v: server returned HTTP %v: '%v'", req.Method, req.URL.String(), rsp.StatusCode, string(b))
	}

	ret := SumResponse{}

	err = m.Unmarshal(rsp.Body, &ret)

	return &ret, errors.Wrap(err, "can't unmarshal response")
}

// patterns for Summator
var (
	pattern_goclay_Summator_Sum_0 = "/v1/example/sum/{a}"

	pattern_goclay_Summator_Sum_0_builder = func(in *SumRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/example/sum/%v", in.A),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Summator_Sum_0_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"b": 0, "a": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}

	pattern_goclay_Summator_Sum_1 = "/v1/example/sum/{a}"

	pattern_goclay_Summator_Sum_1_builder = func(in *SumRequest) string {
		values := url.Values{}

		u := url.URL{
			Path:     fmt.Sprintf("/v1/example/sum/%v", in.A),
			RawQuery: values.Encode(),
		}
		return u.String()
	}

	unmarshaler_goclay_Summator_Sum_1_boundParams = &utilities.DoubleArray{Encoding: map[string]int{"b": 0, "a": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}
)

// marshalers for Summator
var (
	unmarshaler_goclay_Summator_Sum_0 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*SumRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Summator_Sum_0_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			req.B = &NestedB{}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req.B), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}

	unmarshaler_goclay_Summator_Sum_1 = func(r *http.Request) func(interface{}) error {
		return func(rif interface{}) error {
			req := rif.(*SumRequest)

			if err := errors.Wrap(runtime.PopulateQueryParameters(req, r.URL.Query(), unmarshaler_goclay_Summator_Sum_1_boundParams), "couldn't populate query parameters"); err != nil {
				return httpruntime.TransformUnmarshalerError(err)
			}

			req.B = &NestedB{}

			inbound, _ := httpruntime.MarshalerForRequest(r)
			if err := errors.Wrap(inbound.Unmarshal(r.Body, &req.B), "couldn't read request JSON"); err != nil {
				return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
			}
			rctx := chi.RouteContext(r.Context())
			if rctx == nil {
				panic("Only chi router is supported for GETs atm")
			}
			for pos, k := range rctx.URLParams.Keys {
				if err := errors.Wrapf(runtime.PopulateFieldFromPath(req, k, rctx.URLParams.Values[pos]), "can't read '%v' from path", k); err != nil {
					return httptransport.NewMarshalerError(httpruntime.TransformUnmarshalerError(err))
				}
			}

			return nil
		}
	}
)

var _swaggerDef_sum_proto = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "sum.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/sum/{a}": {
      "post": {
        "operationId": "Sum2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/sumpbSumResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "a",
            "description": "A is the number we're adding to. Can't be zero for the sake of example.",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "description": "B is the number we're adding.",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/sumpbNestedB"
            }
          }
        ],
        "tags": [
          "Summator"
        ]
      }
    }
  },
  "definitions": {
    "sumpbNestedB": {
      "type": "object",
      "properties": {
        "b": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "sumpbSumResponse": {
      "type": "object",
      "properties": {
        "sum": {
          "type": "string",
          "format": "int64"
        },
        "error": {
          "type": "string"
        }
      }
    }
  }
}

`)
