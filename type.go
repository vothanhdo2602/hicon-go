// Package hicon provides a client SDK for optimizing database queries through hicon query proxy.
//
// MIT License - see LICENSE file for details.
package hicon

import "github.com/vothanhdo2602/hicon-go/hicon-sm/constant"

type UpsertConfig struct {
	DBConfig     *DBConfig
	Redis        *Redis
	TableConfigs []*TableConfig
	Debug        bool
	DisableCache bool
}

type DBConfig struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	MaxCons  int
	TLS      *TLS
}

type Redis struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
	PoolSize int
}

type TLS struct {
	CertPEM       string
	PrivateKeyPEM string
	RootCAPEM     string
}

type TableConfig struct {
	Name            string
	Columns         []*Column
	RelationColumns []*RelationColumn
}

type Column struct {
	Name         string
	Type         string
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

type RelationColumn struct {
	Name     string
	RefTable string
	Type     string
	Join     string
}

type FindByPK struct {
	table               string
	disableCache        bool
	selects             []string
	data                interface{}
	whereAllWithDeleted bool
}

type FindOne struct {
	table               string
	disableCache        bool
	selects             []string
	where               []*QueryWithArgs
	relations           []string
	joins               []*Join
	offset              int
	orderBy             []string
	whereAllWithDeleted bool
}

type QueryWithArgs struct {
	Query string
	Args  []interface{}
}

type Join struct {
	Join string
	Args []interface{}
}

type Exec struct {
	lockKey string
	sql     string
	args    []interface{}
}

func (s *Exec) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationExec,
		data: s,
	}
}

type FindAll struct {
	table               string
	disableCache        bool
	selects             []string
	where               []*QueryWithArgs
	relations           []string
	joins               []*Join
	limit               int
	offset              int
	orderBy             []string
	whereAllWithDeleted bool
}

type BulkInsert struct {
	lockKey      string
	table        string
	data         []interface{}
	disableCache bool
}

func (s *BulkInsert) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationBulkInsert,
		data: s,
	}
}

type UpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	lockKey      string
	table        string
	data         interface{}
	where        []*QueryWithArgs
	disableCache bool
}

func (s *UpdateByPK) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationUpdateByPK,
		data: s,
	}
}

type UpdateAll struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	lockKey             string
	table               string
	where               []*QueryWithArgs
	set                 []*QueryWithArgs
	whereAllWithDeleted bool
	disableCache        bool
}

func (s *UpdateAll) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationUpdateAll,
		data: s,
	}
}

type BulkUpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	lockKey      string
	table        string
	set          []string
	where        []string
	data         interface{}
	disableCache bool
}

func (s *BulkUpdateByPK) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationBulkUpdateByPK,
		data: s,
	}
}

type DeleteByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	lockKey      string
	table        string
	data         interface{}
	where        []*QueryWithArgs
	disableCache bool
	forceDelete  bool // if enable soft delete in table
}

func (s *DeleteByPK) ToOperation() *Operation {
	return &Operation{
		name: constant.BWOperationDeleteByPK,
		data: s,
	}
}

type BulkWriteWithTx struct {
	lockKey    string
	operations []*Operation
}
type Operation struct {
	name string
	data interface{}
}

type BaseRequest struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}
}
