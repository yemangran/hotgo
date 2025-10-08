// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BsStackLog is the golang structure for table bs_stack_log.
type BsStackLog struct {
	Id          uint64      `json:"id"          orm:"id"            description:"主键"`
	ItemId      int         `json:"itemId"      orm:"item_id"       description:"物品Id"`
	Qty         int         `json:"qty"         orm:"qty"           description:"数量"`
	WorkOrderId int         `json:"workOrderId" orm:"work_order_id" description:"工单Id"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"        description:"备注"`
}
