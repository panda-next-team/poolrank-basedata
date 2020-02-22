package app

import (
	"context"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/panda-next-team/poolrank-basedata/api/internal/app/model"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

const (
	DEFAULT_LIST_POOLS_LIMIT        = 100;
	DEFAULT_LIST_POOLS_MAX_LIMIT    = 500;
	DEFAULT_LIST_POOLS_SKIP         = 0;
	DEFAULT_LIST_POOLS_SORT         = "id ASC";
	PREFIX_LIST_POOLS_QUERY_FIELD   = "Q_"
	PREFIX_LIST_POOLS_SORT_FIELD    = "S_"
	LIST_POOLS_REQ_QUERY_BOOL_VALUE = "QueryBoolValue"
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

func (s *PoolService) AddPoolCoinbaseAddress(ctx context.Context, in *pb.AddPoolCoinbaseAddressRequest) (*pb.AddPoolCoinbaseAddressResponse, error) {
	if in.CoinId <= 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument coin id")
		return nil, st.Err()
	}

	if in.PoolId <= 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument pool id")
		return nil, st.Err()
	}

	if in.Address == "" {
		st := status.New(codes.InvalidArgument, "Invalid argument address")
		return nil, st.Err()
	}

	poolAddress := new(model.PoolAddress)
	has, err := s.Engine.Where("coin_id =? and pool_id =? and type = ? and address = ?", in.CoinId, in.PoolId,
		in.Type, in.Address).Get(poolAddress)

	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}

	if has {
		st := status.New(codes.AlreadyExists, "Already exists entity")
		return nil, st.Err()
	}

	poolAddress.CoinId = in.CoinId
	poolAddress.PoolId = in.PoolId
	poolAddress.Type = int32(in.Type)
	poolAddress.Address = in.Address
	poolAddress.UpdatedAtTs = int32(time.Now().Unix())
	poolAddress.CreatedAtTs = int32(time.Now().Unix())

	_, err = s.Engine.Insert(poolAddress)
	if err != nil {
		st := status.New(codes.Internal, "Server internal error")
		return nil, st.Err()
	}
	return &pb.AddPoolCoinbaseAddressResponse{Result: true}, nil
}

func (s *PoolService) ListPools(ctx context.Context, in *pb.ListPoolsRequest) (*pb.ListPoolsResponse, error) {
	if in.Limit < 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument limit")
		return nil, st.Err()
	}

	if in.Skip < 0 {
		st := status.New(codes.InvalidArgument, "Invalid argument skip")
		return nil, st.Err()
	}

	var listPoolsRequestQueryOperators = map[pb.ListPoolsRequest_QueryOperator]string{
		pb.ListPoolsRequest_EQ:  "=",
		pb.ListPoolsRequest_NE:  "!=",
		pb.ListPoolsRequest_GTE: ">=",
		pb.ListPoolsRequest_LTE: "<=",
		pb.ListPoolsRequest_LT:  "<",
		pb.ListPoolsRequest_GT:  ">",
	}

	var queryString string
	for _, query := range in.Queries {
		field := strings.ToLower(strings.TrimLeft(query.Field.String(), PREFIX_LIST_POOLS_QUERY_FIELD))
		operator := listPoolsRequestQueryOperators[query.Operator]

		typeUrlData := strings.Split(query.Value.GetTypeUrl(), ".")
		typeUrlSuffix := typeUrlData[len(typeUrlData)-1]

		if typeUrlSuffix == LIST_POOLS_REQ_QUERY_BOOL_VALUE {
			queryVal := &pb.ListPoolsRequest_QueryBoolValue{}
			err := ptypes.UnmarshalAny(query.Value, queryVal)
			if err != nil {
				st := status.New(codes.Internal, fmt.Sprintf("Unmarshal query value error:%s", err))
				return nil, st.Err()
			}
			queryString = fmt.Sprintf("%s %s ?", field, operator)

			if queryVal.Value {
				s.Engine.Where(queryString, 1)
			} else {
				s.Engine.Where(queryString, 0)
			}
		}
	}

	total, err := s.Engine.Count(new(model.Pool))
	if err != nil {
		st := status.New(codes.Internal, fmt.Sprintf("get total error:%s", err))
		return nil, st.Err()
	}

	if in.OnlyTotal {
		return &pb.ListPoolsResponse{Total: int32(total)}, nil
	}


	var limit, skip int32
	if in.Limit == 0 {
		limit = DEFAULT_LIST_POOLS_LIMIT
	} else if in.Limit > DEFAULT_LIST_POOLS_MAX_LIMIT {
		limit = DEFAULT_LIST_POOLS_MAX_LIMIT
	} else {
		limit = in.Limit
	}

	if in.Skip == 0 {
		skip = DEFAULT_LIST_POOLS_SKIP
	}

	var sorts []string
	if len(in.Sorts) == 0 {
		sorts = make([]string, 1)
		sorts[0] = DEFAULT_LIST_POOLS_SORT
	} else {
		sorts = make([]string, len(in.Sorts))
		for index, sort := range in.Sorts {
			field := strings.ToLower(strings.TrimLeft(sort.Field.String(), PREFIX_LIST_POOLS_SORT_FIELD))
			if sort.Direction == pb.ListPoolsRequest_ASC {
				sorts[index] = field
			} else {
				sorts[index] = fmt.Sprintf("%s DESC", field)
			}
		}
	}

	var pools []*model.Pool
	err = s.Engine.Limit(int(limit), int(skip)).OrderBy(strings.Join(sorts, ",")).Find(&pools)
	if err != nil {
		st := status.New(codes.Internal, fmt.Sprintf("get total error:%s", err))
		return nil, st.Err()
	}

	pbPools := make([]*pb.Pool, len(pools))

	for index, pool := range pools {
		pbPools[index] = loadPbPool(pool)
	}
	return &pb.ListPoolsResponse{Total: int32(total), Pools: pbPools}, nil
}

func loadPbPool(model *model.Pool) *pb.Pool {
	entity := new(pb.Pool)
	entity.Id = model.Id
	entity.Name = model.Name
	entity.WebsiteUrl = model.WebsiteUrl
	entity.Icon = model.Icon
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
