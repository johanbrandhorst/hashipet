package server

import (
	"context"
	"strings"
	"sync"

	hashipetv1 "github.com/johanbrandhorst/hashipet/proto/hashipet/v1"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// HashiPet implements the HashiPet API
type HashiPet struct {
	mu   *sync.RWMutex
	pets map[string]*hashipetv1.Pet
}

// New initializes a new HashiPet struct.
func New() *HashiPet {
	return &HashiPet{
		mu:   &sync.RWMutex{},
		pets: map[string]*hashipetv1.Pet{},
	}
}

func (s *HashiPet) CreatePet(_ context.Context, req *hashipetv1.CreatePetRequest) (*hashipetv1.CreatePetResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if req.GetPet().GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "a pet must have a name")
	}
	if req.GetPet().GetOwner() == "" {
		return nil, status.Error(codes.InvalidArgument, "a pet must have an owner")
	}
	switch req.GetPet().GetSpecies() {
	case hashipetv1.Species_SPECIES_CAT, hashipetv1.Species_SPECIES_DOG, hashipetv1.Species_SPECIES_PIG:
	default:
		return nil, status.Error(codes.InvalidArgument, "unregonized species")
	}
	name := strings.ToLower(req.GetPet().GetName())
	if _, ok := s.pets[name]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "pet %q already exists", req.GetPet().GetName())
	}
	s.pets[name] = req.GetPet()
	return &hashipetv1.CreatePetResponse{
		Pet: req.GetPet(),
	}, nil
}

func (s *HashiPet) GetPet(_ context.Context, req *hashipetv1.GetPetRequest) (*hashipetv1.GetPetResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	pet, ok := s.pets[strings.ToLower(req.GetName())]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "pet %q not found 😔. Maybe you should create it?", req.GetName())
	}
	return &hashipetv1.GetPetResponse{
		Pet: pet,
	}, nil
}

func (s *HashiPet) ListPets(_ context.Context, req *hashipetv1.ListPetsRequest) (*hashipetv1.ListPetsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var pets []*hashipetv1.Pet
	for _, pet := range s.pets {
		pets = append(pets, proto.Clone(pet).(*hashipetv1.Pet))
	}
	// Sort for deterministic results
	slices.SortFunc(pets, func(a, b *hashipetv1.Pet) bool {
		return a.GetName() < b.GetName()
	})
	return &hashipetv1.ListPetsResponse{
		Pets: pets,
	}, nil
}

func (s *HashiPet) DeletePet(_ context.Context, req *hashipetv1.DeletePetRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := strings.ToLower(req.GetName())
	if _, ok := s.pets[name]; !ok {
		return nil, status.Errorf(codes.NotFound, "pet %q not found 😔", req.GetName())
	}
	delete(s.pets, name)
	return &emptypb.Empty{}, nil
}
