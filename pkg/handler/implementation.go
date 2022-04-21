package handler

import (
	"github.com/amrchnk/account_service/pkg/service"
	pb "github.com/amrchnk/account_service/proto"
	"sync"
)

type Implementation struct {
	pb.UnimplementedAccountServiceServer
	*service.Service
	mu sync.Mutex
}

func NewService(s *service.Service) *Implementation {
	return &Implementation{
		Service: s,
	}
}
