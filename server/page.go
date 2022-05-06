package server

import (
	"bytes"
	"context"
	"embed"
	"html/template"

	hashipetv1 "github.com/johanbrandhorst/hashipet/proto/hashipet/v1"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:embed pets.html.tmpl
var files embed.FS

var petsTemplate = template.Must(template.New("pets").ParseFS(files, "*"))

func (s *HashiPet) ServePets(ctx context.Context, req *emptypb.Empty) (*httpbody.HttpBody, error) {
	resp, err := s.ListPets(ctx, &hashipetv1.ListPetsRequest{})
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := petsTemplate.Execute(&buf, resp.Pets); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to render pets template: %v", err)
	}
	return &httpbody.HttpBody{
		Data:        buf.Bytes(),
		ContentType: "text/html",
	}, nil
}
