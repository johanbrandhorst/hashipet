package server

import (
	"context"
	"sync"

	hashipetv1 "github.com/johanbrandhorst/hashipet/proto/hashipet/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HashiPet implements the HashiPet API
type HashiPet struct {
	mu   *sync.RWMutex
	pets map[string]*hashipetv1.Pet
}

// New initializes a new Backend struct.
func New() *HashiPet {
	return &HashiPet{
		mu: &sync.RWMutex{},
	}
}

func (b *HashiPet) GetPet(_ context.Context, req *hashipetv1.GetPetRequest) (*hashipetv1.GetPetResponse, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	pet, ok := b.pets[req.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "pet %q not found 😔. Maybe you should create it?", req.GetName())
	}
	return &hashipetv1.GetPetResponse{
		Pet: pet,
	}, nil
}
