// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BsReportDetailDao is the data access object for the table hg_bs_report_detail.
type BsReportDetailDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  BsReportDetailColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// BsReportDetailColumns defines and stores column names for the table hg_bs_report_detail.
type BsReportDetailColumns struct {
	Id          string // 主键
	ServiceId   string // 服务id
	WorkOrderId string // 工单id
	Remark      string // 备注
	HandleType  string // 处理方式
	Price       string // 金额
}

// bsReportDetailColumns holds the columns for the table hg_bs_report_detail.
var bsReportDetailColumns = BsReportDetailColumns{
	Id:          "id",
	ServiceId:   "service_id",
	WorkOrderId: "work_order_id",
	Remark:      "remark",
	HandleType:  "handle_type",
	Price:       "price",
}

// NewBsReportDetailDao creates and returns a new DAO object for table data access.
func NewBsReportDetailDao(handlers ...gdb.ModelHandler) *BsReportDetailDao {
	return &BsReportDetailDao{
		group:    "default",
		table:    "hg_bs_report_detail",
		columns:  bsReportDetailColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BsReportDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BsReportDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BsReportDetailDao) Columns() BsReportDetailColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BsReportDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BsReportDetailDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BsReportDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
