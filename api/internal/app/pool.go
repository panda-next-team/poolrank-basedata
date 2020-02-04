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

type PoolService struct {
	Engine *xorm.Engine
}

func (s *PoolService) GetPoolAddresses(ctx context.Context, in *pb.GetPoolAddressesRequest) (*pb.GetPoolAddressesResponse, error) {
	if in.CoinId <= 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument coin_id")
		return nil, st.Err()
	}
	var err error
	poolAddresses := make([]*model.PoolAddress, 0)
	session := s.Engine.Where("coin_id = ?", in.CoinId)
	if in.PoolId > 0 {
		pool := new(model.Pool)
		has, err := s.Engine.Id(in.PoolId).Get(pool)

		if err != nil {
			st := status.New(codes.Internal, "Server internal error.")
			return nil, st.Err()
		}

		if !has {
			st := status.New(codes.NotFound, "Not found entity.")
			return nil, st.Err()
		}

		session.And("pool_id = ?", in.PoolId)
	}

	if in.Type > 0 {
		session.And("type = ?", in.Type)
	}

	err = session.Find(&poolAddresses)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error.")
		return nil, st.Err()
	}

	pbPoolAddresses := make([]*pb.PoolAddress, len(poolAddresses))
	for index, poolAddress := range poolAddresses {
		pbPoolAddresses[index] = loadPbPoolAddress(poolAddress)
	}

	return &pb.GetPoolAddressesResponse{Addresses: pbPoolAddresses}, nil
}

func (s *PoolService) GetPoolCoinbaseTags(ctx context.Context, in *pb.GetPoolCoinbaseTagsRequest) (*pb.GetPoolCoinbaseTagsResponse, error) {
	var err error
	poolTags := make([]*model.PoolCoinbaseTag, 0)

	if in.PoolId > 0 {
		pool := new(model.Pool)
		has, err := s.Engine.Id(in.PoolId).Get(pool)

		if err != nil {
			st := status.New(codes.Internal, "Server internal error.")
			return nil, st.Err()
		}

		if !has {
			st := status.New(codes.NotFound, "Not found entity.")
			return nil, st.Err()
		}

		err = s.Engine.Where("pool_id = ?", in.PoolId).Find(&poolTags)
	} else {
		err = s.Engine.Find(&poolTags)
	}

	if err != nil {
		st := status.New(codes.Internal, "Server internal error.")
		return nil, st.Err()
	}

	pbPoolTags := make([]*pb.PoolCoinbaseTag, len(poolTags))
	for index, poolTag := range poolTags {
		pbPoolTags[index] = loadPbPoolCoinbaseTag(poolTag)
	}

	return &pb.GetPoolCoinbaseTagsResponse{Tags: pbPoolTags}, nil
}

func (s *PoolService) CountPools(ctx context.Context, in *empty.Empty) (*pb.CountPoolsResponse, error) {
	pool := new(model.Pool)
	total, err := s.Engine.Count(pool)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error.")
		return nil, st.Err()
	}

	return &pb.CountPoolsResponse{Count: int32(total)}, nil
}

func (s *PoolService) GetPool(ctx context.Context, in *pb.GetPoolRequest) (*pb.GetPoolResponse, error) {
	if in.Id <= 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument id")
		return nil, st.Err()
	}

	pool := new(model.Pool)
	has, err := s.Engine.Id(in.Id).Get(pool)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if !has {
		st := status.New(codes.NotFound, "Not found entity")
		return nil, st.Err()
	}

	pbPool := loadPbPool(pool)

	return &pb.GetPoolResponse{Pool: pbPool}, nil
}

func loadPbPool(model *model.Pool) *pb.Pool {
	entity := new(pb.Pool)
	entity.Id = model.Id
	entity.Name = model.Name
	entity.WebsiteUrl = model.WebsiteUrl
	entity.Status = int32(model.Status)
	entity.ListOrder = model.ListOrder
	entity.CreatedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}

func loadPbPoolCoinbaseTag(model *model.PoolCoinbaseTag) *pb.PoolCoinbaseTag {
	entity := new(pb.PoolCoinbaseTag)
	entity.Id = model.Id
	entity.PoolId = model.PoolId
	entity.Tag = model.Tag
	entity.CreatedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}

func loadPbPoolAddress(model *model.PoolAddress) *pb.PoolAddress {
	entity := new(pb.PoolAddress)
	entity.Id = model.Id
	entity.PoolId = model.PoolId
	entity.CoinId = model.CoinId
	entity.Address = model.Address
	entity.Type = model.Type
	entity.CreatedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}
