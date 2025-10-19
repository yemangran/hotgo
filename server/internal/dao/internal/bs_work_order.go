// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BsWorkOrderDao is the data access object for the table hg_bs_work_order.
type BsWorkOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BsWorkOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BsWorkOrderColumns defines and stores column names for the table hg_bs_work_order.
type BsWorkOrderColumns struct {
	Id                 string // 主键
	UserId             string // 用户Id
	CreateTime         string // 创建时间
	CustomerName       string // 客户名称
	CustomerAddress    string // 客户地址
	CustomerContact    string // 客户联系方式
	CustomerPerson     string // 客户对接人
	ProblemDescription string // 问题描述
	DefectDescription  string // 检测描述
	DispatchId         string // 分配处理人Id
	SendType           string // 发起类型
	AcceptType         string // 收取方式
	ProductType        string // 产品类型
	Status             string // 工单状态
	SuggestSolution    string // 建议处理方案(多选)
	ActualSolution     string // 实际处理方案(多选)
	TotalMoney         string // 总金额
	InvoiceType        string // 开票类型
	Remark             string // 备注
	DeviceModel        string // 设备型号
	ComponentName      string // 部件名称
}

// bsWorkOrderColumns holds the columns for the table hg_bs_work_order.
var bsWorkOrderColumns = BsWorkOrderColumns{
	Id:                 "id",
	UserId:             "user_id",
	CreateTime:         "create_time",
	CustomerName:       "customer_name",
	CustomerAddress:    "customer_address",
	CustomerContact:    "customer_contact",
	CustomerPerson:     "customer_person",
	ProblemDescription: "problem_description",
	DefectDescription:  "defect_description",
	DispatchId:         "dispatch_id",
	SendType:           "send_type",
	AcceptType:         "accept_type",
	ProductType:        "product_type",
	Status:             "status",
	SuggestSolution:    "suggest_solution",
	ActualSolution:     "actual_solution",
	TotalMoney:         "total_money",
	InvoiceType:        "invoice_type",
	Remark:             "remark",
	DeviceModel:        "device_model",
	ComponentName:      "component_name",
}

// NewBsWorkOrderDao creates and returns a new DAO object for table data access.
func NewBsWorkOrderDao(handlers ...gdb.ModelHandler) *BsWorkOrderDao {
	return &BsWorkOrderDao{
		group:    "default",
		table:    "hg_bs_work_order",
		columns:  bsWorkOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BsWorkOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BsWorkOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BsWorkOrderDao) Columns() BsWorkOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BsWorkOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BsWorkOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BsWorkOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
