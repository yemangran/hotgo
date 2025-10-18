// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BsReportDetail is the golang structure of table hg_bs_report_detail for DAO operations like Where/Data.
type BsReportDetail struct {
	g.Meta      `orm:"table:hg_bs_report_detail, do:true"`
	Id          any // 主键
	ServiceId   any // 服务id
	WorkOrderId any // 工单id
	Remark      any // 备注
	HandleType  any // 处理方式
	Price       any // 金额
}
