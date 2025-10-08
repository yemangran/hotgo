// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BsStackLog is the golang structure of table hg_bs_stack_log for DAO operations like Where/Data.
type BsStackLog struct {
	g.Meta      `orm:"table:hg_bs_stack_log, do:true"`
	Id          interface{} // 主键
	ItemId      interface{} // 物品Id
	Qty         interface{} // 数量
	WorkOrderId interface{} // 工单Id
	CreateTime  *gtime.Time // 创建时间
	Remark      interface{} // 备注
}
