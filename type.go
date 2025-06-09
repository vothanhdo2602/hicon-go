package hicon

import "github.com/vothanhdo2602/hicon-sm/constant"

type UpsertConfig struct {
	DBConfig     *DBConfig
	Redis        *Redis
	TableConfigs []*TableConfig
	Debug        bool
	DisableCache bool
}

type DBConfig struct {
	Type     constant.DBType
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
	Type         constant.ColumnType
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

type RelationColumn struct {
	Name     string
	RefTable string
	Type     constant.RelationType
	Join     string
}

type FindByPK struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Data                interface{}
	WhereAllWithDeleted bool
}

type FindOne struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Where               []*QueryWithArgs
	Relations           []string
	Joins               []*Join
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
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
	LockKey string
	Sql     string
	Args    []interface{}
}

func (s *Exec) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationExec,
		Data: s,
	}
}

type FindAll struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Where               []*QueryWithArgs
	Relations           []string
	Joins               []*Join
	Limit               int
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
}

type BulkInsert struct {
	LockKey      string
	Table        string
	Data         []interface{}
	DisableCache bool
}

func (s *BulkInsert) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationBulkInsert,
		Data: s,
	}
}

type UpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string
	Table        string
	Data         interface{}
	Where        []*QueryWithArgs
	DisableCache bool
}

func (s *UpdateByPK) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationUpdateByPK,
		Data: s,
	}
}

type UpdateAll struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey             string
	Table               string
	Where               []*QueryWithArgs
	Set                 []*QueryWithArgs
	WhereAllWithDeleted bool
	DisableCache        bool
}

func (s *UpdateAll) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationUpdateAll,
		Data: s,
	}
}

type BulkUpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string
	Table        string
	Set          []string
	Where        []string
	Data         interface{}
	DisableCache bool
}

func (s *BulkUpdateByPK) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationBulkUpdateByPK,
		Data: s,
	}
}

type DeleteByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string
	Table        string
	Data         interface{}
	Where        []*QueryWithArgs
	DisableCache bool
	ForceDelete  bool // if enable soft delete in table
}

func (s *DeleteByPK) ToOperation() *Operation {
	return &Operation{
		Name: constant.BWOperationDeleteByPK,
		Data: s,
	}
}

type BulkWriteWithTx struct {
	LockKey    string
	Operations []*Operation
}
type Operation struct {
	Name string
	Data interface{}
}

type BaseRequest struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Shared  bool        `json:"shared"`
	Success bool        `json:"success"`
}
