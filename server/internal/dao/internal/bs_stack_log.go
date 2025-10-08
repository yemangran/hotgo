// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BsStackLogDao is the data access object for the table hg_bs_stack_log.
type BsStackLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BsStackLogColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BsStackLogColumns defines and stores column names for the table hg_bs_stack_log.
type BsStackLogColumns struct {
	Id          string // 主键
	ItemId      string // 物品Id
	Qty         string // 数量
	WorkOrderId string // 工单Id
	CreateTime  string // 创建时间
	Remark      string // 备注
}

// bsStackLogColumns holds the columns for the table hg_bs_stack_log.
var bsStackLogColumns = BsStackLogColumns{
	Id:          "id",
	ItemId:      "item_id",
	Qty:         "qty",
	WorkOrderId: "work_order_id",
	CreateTime:  "create_time",
	Remark:      "remark",
}

// NewBsStackLogDao creates and returns a new DAO object for table data access.
func NewBsStackLogDao(handlers ...gdb.ModelHandler) *BsStackLogDao {
	return &BsStackLogDao{
		group:    "default",
		table:    "hg_bs_stack_log",
		columns:  bsStackLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BsStackLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BsStackLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BsStackLogDao) Columns() BsStackLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BsStackLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BsStackLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BsStackLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
