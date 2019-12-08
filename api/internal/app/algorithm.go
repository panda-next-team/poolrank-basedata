package app

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/panda-next-team/poolrank-basedata/api/internal/app/model"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlgorithmService struct {
	Engine *xorm.Engine
}

func (s *AlgorithmService) CountAlgorithms(ctx context.Context, in *empty.Empty) (*pb.CountAlgorithmsResponse, error) {
	algorithm := new(model.Algorithm)
	total, err := s.Engine.Count(algorithm)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}
	return &pb.CountAlgorithmsResponse{Count: int32(total)}, nil
}

func (s *AlgorithmService) GetAlgorithm(ctx context.Context, in *pb.GetAlgorithmRequest) (*pb.GetAlgorithmResponse, error) {
	if in.Id <= 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument id")
		return nil, st.Err()
	}

	algorithm := new(model.Algorithm)
	has, err := s.Engine.Id(in.Id).Get(algorithm)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if !has {
		st := status.New(codes.NotFound, "Not found entity")
		return nil, st.Err()
	}

	pbAlgorithm := loadPbAlgorithm(algorithm)
	return &pb.GetAlgorithmResponse{Algorithm: pbAlgorithm}, nil
}

func loadPbAlgorithm(model *model.Algorithm) *pb.Algorithm {
	entity := new(pb.Algorithm)
	entity.Id = model.Id
	entity.Name = model.Name
	entity.CratedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}
