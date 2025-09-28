package internal

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	registry "nerosoft.com/aone/registry-server/gen"
)

type RegistryService struct {
	reg *Registry
}

func NewRegistryService(reg *Registry) *RegistryService {
	return &RegistryService{reg: reg}
}

func (s *RegistryService) Register(
	ctx context.Context,
	req *connect.Request[registry.RegisterRequest],
) (*connect.Response[registry.RegisterReply], error) {
	err := s.reg.Register(req.Msg.Service, req.Msg.Address, int(req.Msg.TtlSeconds))
	if err != nil {
		return nil, fmt.Errorf("register error: %w", err)
	}
	return connect.NewResponse(&registry.RegisterReply{}), nil
}

func (s *RegistryService) Deregister(
	ctx context.Context,
	req *connect.Request[registry.DeregisterRequest],
) (*connect.Response[registry.DeregisterReply], error) {
	err := s.reg.Unregister(req.Msg.Service, req.Msg.Address)
	if err != nil {
		return nil, fmt.Errorf("unregister error: %w", err)
	}
	return connect.NewResponse(&registry.DeregisterReply{}), nil
}

func (s *RegistryService) GetNodes(
	ctx context.Context,
	req *connect.Request[registry.GetNodesRequest],
) (*connect.Response[registry.GetNodesReply], error) {
	nodes, err := s.reg.GetNodes(req.Msg.Service)
	if err != nil {
		return nil, fmt.Errorf("list error: %w", err)
	}
	return connect.NewResponse(&registry.GetNodesReply{
		Nodes: nodes,
	}), nil
}
