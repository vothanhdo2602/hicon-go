package grpcapi

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/pkg/service"
	"google.golang.org/protobuf/types/known/anypb"
)

type SQLExecutor struct {
	sqlexecutor.UnimplementedSQLExecutorServer
}

func (SQLExecutor) UpsertConfiguration(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.UpsertConfiguration
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	r := svc.UpsertConfiguration(ctx, &req)

	return grpctil.NewResponse(r), nil
}

func (SQLExecutor) FindByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.FindByPK
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.FindByPK(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) FindOne(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.FindOne
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.FindOne(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) FindAll(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.FindAll
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.FindAll(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) Exec(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.Exec
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.Exec(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkInsert(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BulkInsert
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.BulkInsert(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) UpdateByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.UpdateByPK
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.UpdateByPK(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) UpdateAll(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.UpdateAll
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.UpdateAll(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkUpdateByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BulkUpdateByPK
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.BulkUpdateByPK(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) DeleteByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.DeleteByPK
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.DeleteByPK(ctx, &req)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkWriteWithTx(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BulkWriteWithTx
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	resp := svc.BulkWriteWithTx(ctx, &req)

	return grpctil.NewResponse(resp), nil
}
