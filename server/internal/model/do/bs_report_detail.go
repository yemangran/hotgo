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
	Id          interface{} // 主键
	ServiceId   interface{} // 服务id
	WorkOrderId interface{} // 工单id
	Remark      interface{} // 备注
	HandleType  interface{} // 处理方式
}
