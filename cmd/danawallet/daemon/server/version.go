package server

import (
	"context"
	"github.com/danannet/danad/cmd/danawallet/daemon/pb"
	"github.com/danannet/danad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
