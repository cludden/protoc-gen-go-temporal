package codec

import (
	"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoJSONCodec implements a converter.PayloadCodec that provides conversion between
// binary/protobuf and json/protobuf encodings. It can be used in conjunction
// with converter.NewPayloadCodecHTTPHandler to implement a Remote Codec Server.
type ProtoJSONCodec struct {
	scheme *scheme.Scheme
}

// NewProtoJSONCodec initializes a new Codec value from one or more Scheme values
func NewProtoJSONCodec(schemes ...*scheme.Scheme) converter.PayloadCodec {
	c := &ProtoJSONCodec{
		scheme: scheme.New(),
	}
	for _, scheme := range schemes {
		c.scheme.Merge(scheme)
	}
	return c
}

// Decode converts all binary/protobuf encoded payloads with registered
// message types to json/protobuf encoding to support Temporal UI automatic
// decoding. Any payloads with other encoding, or unregistered types are
// left unmodified.
func (c *ProtoJSONCodec) Decode(payloads []*common.Payload) ([]*common.Payload, error) {
	results := make([]*common.Payload, len(payloads))
	for i, p := range payloads {
		// skip non binary/protobuf payloads
		if encoding, ok := p.GetMetadata()[converter.MetadataEncoding]; !ok || string(encoding) != converter.MetadataEncodingProto {
			results[i] = p
			continue
		}
		// skip payloads missing a message type header
		t, ok := p.GetMetadata()[converter.MetadataMessageType]
		if !ok {
			results[i] = p
			continue
		}
		// skip payloads with unregistered types
		out, err := c.scheme.New(string(t))
		if err != nil || out == nil {
			results[i] = p
			continue
		}
		// skip payloads with invalid binary payload
		if err := proto.Unmarshal(p.GetData(), out); err != nil {
			results[i] = p
			continue
		}
		// skip payloads that fail to serialize to json
		b, err := protojson.Marshal(out)
		if err != nil {
			results[i] = p
			continue
		}
		results[i] = &common.Payload{
			Data: b,
			Metadata: map[string][]byte{
				converter.MetadataEncoding:    []byte(converter.MetadataEncodingProtoJSON),
				converter.MetadataMessageType: t,
			},
		}
	}
	return results, nil
}

// Encode converts all json/protobuf encoded payloads with registered message
// types to binary/protobuf encoding prior to being forwarded to Temporal.
func (c *ProtoJSONCodec) Encode(payloads []*common.Payload) ([]*common.Payload, error) {
	results := make([]*common.Payload, len(payloads))
	for i, p := range payloads {
		// skip non binary/protobuf payloads
		if encoding, ok := p.GetMetadata()[converter.MetadataEncoding]; !ok || string(encoding) != converter.MetadataEncodingProtoJSON {
			results[i] = p
			continue
		}
		// skip payloads missing a message type header
		t, ok := p.GetMetadata()[converter.MetadataMessageType]
		if !ok {
			results[i] = p
			continue
		}
		// skip payloads with unregistered types
		out, err := c.scheme.New(string(t))
		if err != nil || out == nil {
			results[i] = p
			continue
		}
		// skip payloads with invalid binary payload
		if err := protojson.Unmarshal(p.GetData(), out); err != nil {
			results[i] = p
			continue
		}
		// skip payloads that fail to serialize to json
		b, err := proto.Marshal(out)
		if err != nil {
			results[i] = p
			continue
		}
		results[i] = &common.Payload{
			Data: b,
			Metadata: map[string][]byte{
				converter.MetadataEncoding:    []byte(converter.MetadataEncodingProto),
				converter.MetadataMessageType: t,
			},
		}
	}
	return results, nil
}

// JSONCodec implements a converter.PayloadCodec that provides conversion between
// binary/protobuf and json/plain encodings. It can be used in conjunction
// with converter.NewPayloadCodecHTTPHandler to implement a Remote Codec Server.
type JSONCodec struct {
	scheme *scheme.Scheme
}

// NewJSONCodec initializes a new Codec value from one or more Scheme values
func NewJSONCodec(schemes ...*scheme.Scheme) converter.PayloadCodec {
	c := &JSONCodec{
		scheme: scheme.New(),
	}
	for _, scheme := range schemes {
		c.scheme.Merge(scheme)
	}
	return c
}

// Decode converts all binary/protobuf encoded payloads with registered
// message types to json/protobuf encoding to support Temporal UI automatic
// decoding. Any payloads with other encoding, or unregistered types are
// left unmodified.
func (c *JSONCodec) Decode(payloads []*common.Payload) ([]*common.Payload, error) {
	results := make([]*common.Payload, len(payloads))
	for i, p := range payloads {
		// skip non binary/protobuf payloads
		if encoding, ok := p.GetMetadata()[converter.MetadataEncoding]; !ok || string(encoding) != converter.MetadataEncodingProto {
			results[i] = p
			continue
		}
		// skip payloads missing a message type header
		t, ok := p.GetMetadata()[converter.MetadataMessageType]
		if !ok {
			results[i] = p
			continue
		}
		// skip payloads with unregistered types
		out, err := c.scheme.New(string(t))
		if err != nil || out == nil {
			results[i] = p
			continue
		}
		// skip payloads with invalid binary payload
		if err := proto.Unmarshal(p.GetData(), out); err != nil {
			results[i] = p
			continue
		}
		// skip payloads that fail to serialize to json
		b, err := protojson.Marshal(out)
		if err != nil {
			results[i] = p
			continue
		}
		results[i] = &common.Payload{
			Data: b,
			Metadata: map[string][]byte{
				converter.MetadataEncoding:    []byte(converter.MetadataEncodingJSON),
				converter.MetadataMessageType: t,
			},
		}
	}
	return results, nil
}

// Encode converts all json/plain encoded payloads with registered message
// types to binary/protobuf encoding prior to being forwarded to Temporal.
func (c *JSONCodec) Encode(payloads []*common.Payload) ([]*common.Payload, error) {
	results := make([]*common.Payload, len(payloads))
	for i, p := range payloads {
		// skip non binary/protobuf payloads
		if encoding, ok := p.GetMetadata()[converter.MetadataEncoding]; !ok || string(encoding) != converter.MetadataEncodingJSON {
			results[i] = p
			continue
		}
		// skip payloads missing a message type header
		t, ok := p.GetMetadata()[converter.MetadataMessageType]
		if !ok {
			results[i] = p
			continue
		}
		// skip payloads with unregistered types
		out, err := c.scheme.New(string(t))
		if err != nil || out == nil {
			results[i] = p
			continue
		}
		// skip payloads with invalid binary payload
		if err := protojson.Unmarshal(p.GetData(), out); err != nil {
			results[i] = p
			continue
		}
		// skip payloads that fail to serialize to json
		b, err := proto.Marshal(out)
		if err != nil {
			results[i] = p
			continue
		}
		results[i] = &common.Payload{
			Data: b,
			Metadata: map[string][]byte{
				converter.MetadataEncoding:    []byte(converter.MetadataEncodingProto),
				converter.MetadataMessageType: t,
			},
		}
	}
	return results, nil
}
