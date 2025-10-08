// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BsServiceDao is the data access object for the table hg_bs_service.
type BsServiceDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BsServiceColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BsServiceColumns defines and stores column names for the table hg_bs_service.
type BsServiceColumns struct {
	Id       string // 主键
	Pid      string // 父键
	Level    string // 关系树级别
	Tree     string // 关系树
	Name     string // 服务名称
	Price    string // 价格
	OrderNum string // 序号
}

// bsServiceColumns holds the columns for the table hg_bs_service.
var bsServiceColumns = BsServiceColumns{
	Id:       "id",
	Pid:      "pid",
	Level:    "level",
	Tree:     "tree",
	Name:     "name",
	Price:    "price",
	OrderNum: "order_num",
}

// NewBsServiceDao creates and returns a new DAO object for table data access.
func NewBsServiceDao(handlers ...gdb.ModelHandler) *BsServiceDao {
	return &BsServiceDao{
		group:    "default",
		table:    "hg_bs_service",
		columns:  bsServiceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BsServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BsServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BsServiceDao) Columns() BsServiceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BsServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BsServiceDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BsServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
