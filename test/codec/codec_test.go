package codec

import (
	"net/http/httptest"
	"testing"

	simplev1 "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/codec"
	"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestJSONCodec(t *testing.T) {
	require := require.New(t)

	// initialize protobuf message value
	msg := &simplev1.OtherWorkflowRequest{
		SomeVal: "example",
		ExampleOneof: &simplev1.OtherWorkflowRequest_ExampleEnum{
			ExampleEnum: simplev1.OtherEnum_OTHER_FOO,
		},
	}

	// initialize scheme
	scheme := scheme.New(
		simplev1.WithIgnoredSchemeTypes(),
		simplev1.WithOtherSchemeTypes(),
		simplev1.WithSimpleSchemeTypes(),
	)
	var payloads []*common.Payload

	// marshal json/protobuf payload
	protojsondc := converter.NewProtoJSONPayloadConverterWithOptions(converter.ProtoJSONPayloadConverterOptions{
		AllowUnknownFields: true,
	})
	protojsonp, err := protojsondc.ToPayload(msg)
	require.NoError(err)
	require.Equal(string(converter.MetadataEncodingProtoJSON), string(protojsonp.Metadata[converter.MetadataEncoding]))
	payloads = append(payloads, protojsonp)

	// marshal json/plain payload
	jsondc := converter.NewJSONPayloadConverter()
	jsonp, err := jsondc.ToPayload(msg)
	require.NoError(err)
	require.Equal(string(converter.MetadataEncodingJSON), string(jsonp.Metadata[converter.MetadataEncoding]))
	payloads = append(payloads, jsonp)

	// marshal binary/protobuf payload
	protodc := converter.NewProtoPayloadConverter()
	protop, err := protodc.ToPayload(msg)
	require.NoError(err)
	require.Equal(string(converter.MetadataEncodingProto), string(protop.Metadata[converter.MetadataEncoding]))
	payloads = append(payloads, protop)

	// verify decode successful
	c := codec.NewJSONCodec(scheme)
	results, err := c.Decode(payloads)
	require.NoError(err)
	require.Len(results, len(payloads))

	// verify json/protobuf payload unchanged
	require.Equal(protojsonp.Data, results[0].Data)
	require.Equal(string(protojsonp.Metadata[converter.MetadataEncoding]), string(results[0].Metadata[converter.MetadataEncoding]))
	require.Equal(string(protojsonp.Metadata[converter.MetadataMessageType]), string(results[0].Metadata[converter.MetadataMessageType]))

	// verify json/plain payload unchanged
	require.Equal(jsonp.Data, results[1].Data)
	require.Equal(string(jsonp.Metadata[converter.MetadataEncoding]), string(results[1].Metadata[converter.MetadataEncoding]))

	// verify binary/protobuf payload conversion to binary/protobuf
	require.NotEqual(protop.Data, results[2].Data)
	require.Equal(string(converter.MetadataEncodingJSON), string(results[2].Metadata[converter.MetadataEncoding]))
	require.Equal(string(protop.Metadata[converter.MetadataMessageType]), string(results[2].Metadata[converter.MetadataMessageType]))
	require.JSONEq(`{"someVal":"example","exampleEnum":"OTHER_FOO"}`, string(results[2].Data))

	// verify conversion back to proto message value
	var out simplev1.OtherWorkflowRequest
	require.NoError(protojson.Unmarshal(results[2].Data, &out))
	require.True(proto.Equal(msg, &out))

	handler := converter.NewPayloadCodecHTTPHandler(c)
	srv := httptest.NewServer(handler)
	defer srv.Close()
	dc := converter.NewRemoteDataConverter(converter.NewCompositeDataConverter(jsondc), converter.RemoteDataConverterOptions{
		Endpoint: srv.URL,
	})

	jsonout := make(map[string]any)
	require.NoError(dc.FromPayload(protop, &jsonout))
	require.Equal("example", jsonout["someVal"])
	require.Equal("OTHER_FOO", jsonout["exampleEnum"])
}

func TestProtoJSONCodec(t *testing.T) {
	require := require.New(t)

	// initialize protobuf message value
	msg := &simplev1.OtherWorkflowRequest{
		SomeVal: "example",
		ExampleOneof: &simplev1.OtherWorkflowRequest_ExampleEnum{
			ExampleEnum: simplev1.OtherEnum_OTHER_FOO,
		},
	}

	// initialize scheme
	scheme := scheme.New(
		simplev1.WithIgnoredSchemeTypes(),
		simplev1.WithOtherSchemeTypes(),
		simplev1.WithSimpleSchemeTypes(),
	)
	var payloads []*common.Payload

	// marshal binary/protobuf payload
	protodc := converter.NewProtoPayloadConverter()
	protop, err := protodc.ToPayload(msg)
	require.NoError(err)
	require.Equal(string(converter.MetadataEncodingProto), string(protop.Metadata[converter.MetadataEncoding]))
	payloads = append(payloads, protop)

	// verify decode successful
	c := codec.NewProtoJSONCodec(scheme)
	results, err := c.Decode(payloads)
	require.NoError(err)
	require.Len(results, len(payloads))

	// verify binary/protobuf payload conversion to binary/protobuf
	require.NotEqual(protop.Data, results[0].Data)
	require.Equal(string(converter.MetadataEncodingProtoJSON), string(results[0].Metadata[converter.MetadataEncoding]))
	require.Equal(string(protop.Metadata[converter.MetadataMessageType]), string(results[0].Metadata[converter.MetadataMessageType]))
	require.JSONEq(`{"someVal":"example","exampleEnum":"OTHER_FOO"}`, string(results[0].Data))

	// verify conversion back to proto message value
	var out simplev1.OtherWorkflowRequest
	require.NoError(protojson.Unmarshal(results[0].Data, &out))
	require.True(proto.Equal(msg, &out))

	handler := converter.NewPayloadCodecHTTPHandler(c)
	srv := httptest.NewServer(handler)
	defer srv.Close()
	dc := converter.NewRemoteDataConverter(
		converter.NewCompositeDataConverter(
			converter.NewProtoJSONPayloadConverterWithOptions(converter.ProtoJSONPayloadConverterOptions{
				AllowUnknownFields: true,
			}),
		),
		converter.RemoteDataConverterOptions{
			Endpoint: srv.URL,
		},
	)

	out = simplev1.OtherWorkflowRequest{}
	require.NoError(dc.FromPayload(protop, &out))
	require.Equal(msg.SomeVal, out.SomeVal)
	require.True(proto.Equal(msg, &out))
}
