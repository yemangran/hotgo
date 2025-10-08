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
	Id                 interface{} // 主键
	UserId             interface{} // 用户Id
	CreateTime         *gtime.Time // 创建时间
	CustomerName       interface{} // 客户名称
	CustomerAddress    interface{} // 客户地址
	CustomerContact    interface{} // 客户联系方式
	CustomerPerson     interface{} // 客户对接人
	ProblemDescription interface{} // 问题描述
	DefectDescription  interface{} // 检测描述
	DispatchId         interface{} // 分配处理人Id
	SendType           interface{} // 发起类型
	AcceptType         interface{} // 收取方式
	ProductType        interface{} // 产品类型
	Status             interface{} // 工单状态
	SuggestSolution    interface{} // 建议处理方案(多选)
	ActualSolution     interface{} // 实际处理方案(多选)
	TotalMoney         interface{} // 总金额
	InvoiceType        interface{} // 开票类型
	Remark             interface{} // 备注
}
