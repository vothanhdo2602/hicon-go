package grpcapi

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/pkg/service"
	"google.golang.org/protobuf/types/known/anypb"
)

type SQLExecutor struct {
	sqlexecutor.UnimplementedSQLExecutorServer
}

func (SQLExecutor) UpsertConfig(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.UpsertConfig]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	r := svc.UpsertConfig(ctx, req.Body)

	return grpctil.NewResponse(r), nil
}

func (SQLExecutor) FindByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.FindByPK]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.FindByPK(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) FindOne(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.FindOne]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.FindOne(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) FindAll(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.FindAll]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.FindAll(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) Exec(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.Exec]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.Exec(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkInsert(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.BulkInsert]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.BulkInsert(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) UpdateByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.UpdateByPK]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.UpdateByPK(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) UpdateAll(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.UpdateAll]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.UpdateAll(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkUpdateByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.BulkUpdateByPK]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.BulkUpdateByPK(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) DeleteByPK(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.DeleteByPK]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.DeleteByPK(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}

func (SQLExecutor) BulkWriteWithTx(ctx context.Context, data *anypb.Any) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req requestmodel.BaseRequestWithType[requestmodel.BulkWriteWithTx]
		svc = service.SQLExecutor()
	)

	err := json.Unmarshal(data.Value, &req)
	if err != nil {
		return nil, err
	}

	ctx = log.GetContext(ctx, req.Headers)
	resp := svc.BulkWriteWithTx(ctx, req.Body)

	return grpctil.NewResponse(resp), nil
}
