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

type PowCoinService struct {
	Engine *xorm.Engine
}

func (s *PowCoinService) CountPowCoins(ctx context.Context, in *empty.Empty) (*pb.CountPowCoinsResponse, error) {
	powCoin := new(model.PowCoin)
	session := s.Engine.NoCache()
	total,err := session.Count(powCoin)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	return &pb.CountPowCoinsResponse{Count: int32(total)}, nil
}

func (s *PowCoinService) GetPowCoin(ctx context.Context, in *pb.GetPowCoinRequest) (*pb.GetPowCoinResponse, error) {

	powCoin := new(model.PowCoin)
	has, err := s.Engine.Id(in.Id).Get(powCoin)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if !has {
		st := status.New(codes.NotFound, "Not found entity")
		return nil, st.Err()
	}

	pbPowCoin := loadPbPowCoin(powCoin)

	return &pb.GetPowCoinResponse{Coin: pbPowCoin}, nil
}

func loadPbPowCoin(model *model.PowCoin) *pb.Pow_Coin {
	entity := new(pb.Pow_Coin)
	entity.Id = int32(model.Id)
	entity.Name = model.Name
	entity.EnName = model.EnName
	entity.EnTag = model.EnTag
	entity.MaxSupply = model.MaxSupply
	entity.AlgorithmId = model.AlgorithmId
	entity.ReleaseDate = model.ReleaseDate.Format("2006-01-02")
	entity.BlockTime = model.BlockTime
	entity.Icon = model.Icon
	entity.GithubUrl = model.GithubUrl
	entity.WebsiteUrl = model.WebsiteUrl
	entity.Intro = model.Intro
	entity.Status = model.Status
	entity.ListOrder = model.ListOrder
	entity.CreatedAtTs = model.CreatedAtTs
	entity.CreatedAt = model.CreatedAt.Format("2006-01-02 15:04:05")
	entity.UpdatedAtTs = model.UpdatedAtTs
	entity.UpdatedAt = model.UpdatedAt.Format("2006-01-02 15:04:05")
	return entity
}
