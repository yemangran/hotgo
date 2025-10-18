// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BsReportDetail is the golang structure for table bs_report_detail.
type BsReportDetail struct {
	Id          uint64  `json:"id"          orm:"id"            description:"主键"`
	ServiceId   int     `json:"serviceId"   orm:"service_id"    description:"服务id"`
	WorkOrderId int     `json:"workOrderId" orm:"work_order_id" description:"工单id"`
	Remark      string  `json:"remark"      orm:"remark"        description:"备注"`
	HandleType  string  `json:"handleType"  orm:"handle_type"   description:"处理方式"`
	Price       float64 `json:"price"       orm:"price"         description:"金额"`
}
