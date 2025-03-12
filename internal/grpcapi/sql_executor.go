package grpcapi

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/pkg/service"
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

		for _, c := range t.RelationColumnConfigs {
			col := &requestmodel.RelationColumnConfigs{
				Name:     c.Name,
				RefTable: c.RefTable,
				Type:     c.Type,
				Join:     c.Join,
			}
			tbl.RelationColumnConfigs = append(tbl.RelationColumnConfigs, col)
		}

		req.TableConfigurations = append(req.TableConfigurations, tbl)
	}

	r := svc.UpdateConfiguration(ctx, req)

	return NewResponse(r), nil
}

func (SQLExecutor) FindByPrimaryKeys(ctx context.Context, data *sqlexecutor.FindByPrimaryKeys) (*sqlexecutor.BaseResponse, error) {
	defer commontil.Recover(ctx)

	var (
		req = &requestmodel.FindByPrimaryKeys{
			Table:        data.Table,
			DisableCache: data.DisableCache,
			PrimaryKeys:  AnyMapToInterfaceMap(data.PrimaryKeys),
			Select:       data.Select,
		}
		svc = service.SQLExecutor()
	)

	resp := svc.FindByPrimaryKeys(ctx, req)

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
