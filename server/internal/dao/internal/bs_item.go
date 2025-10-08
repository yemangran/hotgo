// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BsItemDao is the data access object for the table hg_bs_item.
type BsItemDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BsItemColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BsItemColumns defines and stores column names for the table hg_bs_item.
type BsItemColumns struct {
	Id       string // 主键
	Name     string // 物品名称
	Category string // 物品类别
	Spec     string // 规格型号
	Price    string // 单价
}

// bsItemColumns holds the columns for the table hg_bs_item.
var bsItemColumns = BsItemColumns{
	Id:       "id",
	Name:     "name",
	Category: "category",
	Spec:     "spec",
	Price:    "price",
}

// NewBsItemDao creates and returns a new DAO object for table data access.
func NewBsItemDao(handlers ...gdb.ModelHandler) *BsItemDao {
	return &BsItemDao{
		group:    "default",
		table:    "hg_bs_item",
		columns:  bsItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BsItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BsItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BsItemDao) Columns() BsItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BsItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BsItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BsItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
