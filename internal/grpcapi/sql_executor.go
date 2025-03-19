package grpcapi

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SQLExecutor struct {
	sqlexecutor.UnimplementedSQLExecutorServer
}

func (SQLExecutor) UpsertConfiguration(ctx context.Context, data *sqlexecutor.UpsertConfiguration) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	if data.DbConfiguration == nil {
		return NewResponse(responsemodel.R400("field DbConfiguration is required")), nil
	}

	var (
		svc = service.SQLExecutor()
		req = &requestmodel.UpsertConfiguration{
			DBConfiguration: &requestmodel.DBConfiguration{
				Type:     data.DbConfiguration.Type,
				Host:     data.DbConfiguration.Host,
				Port:     int(data.DbConfiguration.Port),
				Username: data.DbConfiguration.Username,
				Password: data.DbConfiguration.Password,
				Database: data.DbConfiguration.Database,
				MaxCons:  int(data.DbConfiguration.MaxCons),
			},
			Redis: &requestmodel.Redis{
				Host:     data.Redis.Host,
				Port:     int(data.Redis.Port),
				Username: data.Redis.Username,
				Password: data.Redis.Password,
				DB:       int(data.Redis.Db),
				PoolSize: int(data.Redis.PoolSize),
			},
			TableConfigurations: []*requestmodel.TableConfiguration{},
			Debug:               data.Debug,
			DisableCache:        data.DisableCache,
		}
	)

	if data.DbConfiguration.TLS != nil {
		req.DBConfiguration.TLS = &requestmodel.TLS{
			CertPEM:       data.DbConfiguration.TLS.CertPem,
			PrivateKeyPEM: data.DbConfiguration.TLS.PrivateKeyPem,
			RootCAPEM:     data.DbConfiguration.TLS.RootCaPem,
		}
	}

	for _, t := range data.TableConfigurations {
		var (
			tbl = &requestmodel.TableConfiguration{
				Name: t.Name,
			}
		)

		for _, c := range t.ColumnConfigs {
			col := &requestmodel.ColumnConfig{
				Name:         c.Name,
				Type:         c.Type,
				Nullable:     c.Nullable,
				IsPrimaryKey: c.IsPrimaryKey,
			}
			tbl.ColumnConfigs = append(tbl.ColumnConfigs, col)
		}

		for _, c := range t.RelationColumns {
			col := &requestmodel.RelationColumns{
				Name:     c.Name,
				RefTable: c.RefTable,
				Type:     c.Type,
				Join:     c.Join,
			}
			tbl.RelationColumns = append(tbl.RelationColumns, col)
		}

		req.TableConfigurations = append(req.TableConfigurations, tbl)
	}

	r := svc.UpsertConfiguration(ctx, req)

	return NewResponse(r), nil
}

func (SQLExecutor) FindByPK(ctx context.Context, data *sqlexecutor.FindByPK) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.FindByPK{
			Table:        data.Table,
			DisableCache: data.DisableCache,
			Data:         AnyMapToInterfaceMap(data.Data),
			Select:       data.Select,
		}
		svc = service.SQLExecutor()
	)

	resp := svc.FindByPK(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) FindOne(ctx context.Context, data *sqlexecutor.FindOne) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.FindOne{
			Table:        data.Table,
			DisableCache: data.DisableCache,
			Select:       data.Select,
			Offset:       int(data.Offset),
			OrderBy:      data.OrderBy,
		}
		svc = service.SQLExecutor()
	)

	for _, protoWhere := range data.Where {
		w, err := ConvertWhereProtoToGo(protoWhere)
		if err != nil {
			return NewResponse(responsemodel.R400(err.Error())), nil
		}

		req.Where = append(req.Where, w)
	}

	for _, rel := range data.Relations {
		req.Relations = append(req.Relations, rel)
	}

	resp := svc.FindOne(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) FindAll(ctx context.Context, data *sqlexecutor.FindAll) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.FindAll{
			Table:        data.Table,
			DisableCache: data.DisableCache,
			Select:       data.Select,
			Offset:       int(data.Offset),
			OrderBy:      data.OrderBy,
		}
		svc = service.SQLExecutor()
	)

	for _, protoWhere := range data.Where {
		w, err := ConvertWhereProtoToGo(protoWhere)
		if err != nil {
			return NewResponse(responsemodel.R400(err.Error())), nil
		}

		req.Where = append(req.Where, w)
	}

	for _, rel := range data.Relations {
		req.Relations = append(req.Relations, rel)
	}

	for _, rel := range data.Joins {
		j := &requestmodel.Join{
			Join: rel.Join,
		}

		for _, arg := range rel.Args {
			val, _ := ConvertAnyToInterface(arg)
			j.Args = append(j.Args, val)
		}

		req.Joins = append(req.Joins, j)
	}

	resp := svc.FindAll(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) Exec(ctx context.Context, data *sqlexecutor.Exec) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.Exec{
			SQL:     data.Sql,
			LockKey: data.LockKey,
		}
		svc = service.SQLExecutor()
	)

	for _, arg := range data.Args {
		val, _ := ConvertAnyToInterface(arg)
		req.Args = append(req.Args, val)
	}

	resp := svc.Exec(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) BulkInsert(ctx context.Context, data *sqlexecutor.BulkInsert) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.BulkInsert{
			LockKey:      data.LockKey,
			Table:        data.Table,
			DisableCache: data.DisableCache,
		}
		svc = service.SQLExecutor()
	)

	r, err := SliceConvertAnyToInterface(data.Data)
	if err != nil {
		return nil, err
	}
	req.Data = r

	resp := svc.BulkInsert(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) UpdateByPK(ctx context.Context, data *sqlexecutor.UpdateByPK) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.UpdateByPK{
			LockKey:      data.LockKey,
			Table:        data.Table,
			DisableCache: data.DisableCache,
		}
		svc = service.SQLExecutor()
	)

	r, err := ConvertAnyToInterface(data.Data)
	if err != nil {
		return nil, err
	}
	req.Data = r

	resp := svc.UpdateByPK(ctx, req)

	return NewResponse(resp), nil
}

func (SQLExecutor) BulkUpdateByPK(ctx context.Context, data *sqlexecutor.BulkUpdateByPK) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	return nil, status.Errorf(codes.Unimplemented, "method BulkUpdateByPK not implemented")
}
