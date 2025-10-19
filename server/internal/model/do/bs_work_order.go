// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BsWorkOrder is the golang structure of table hg_bs_work_order for DAO operations like Where/Data.
type BsWorkOrder struct {
	g.Meta             `orm:"table:hg_bs_work_order, do:true"`
	Id                 any         // 主键
	UserId             any         // 用户Id
	CreateTime         *gtime.Time // 创建时间
	CustomerName       any         // 客户名称
	CustomerAddress    any         // 客户地址
	CustomerContact    any         // 客户联系方式
	CustomerPerson     any         // 客户对接人
	ProblemDescription any         // 问题描述
	DefectDescription  any         // 检测描述
	DispatchId         any         // 分配处理人Id
	SendType           any         // 发起类型
	AcceptType         any         // 收取方式
	ProductType        any         // 产品类型
	Status             any         // 工单状态
	SuggestSolution    any         // 建议处理方案(多选)
	ActualSolution     any         // 实际处理方案(多选)
	TotalMoney         any         // 总金额
	InvoiceType        any         // 开票类型
	Remark             any         // 备注
	DeviceModel        any         // 设备型号
	ComponentName      any         // 部件名称
}
