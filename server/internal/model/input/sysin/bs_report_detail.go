package sysin

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
)

type BsReportDetailListInp struct {
	form.PageReq
	entity.BsReportDetail
}

type BsReportDetailListModel struct {
	Id          uint64  `json:"id" dc:"主键"`
	ServiceId   int     `json:"serviceId" dc:"服务id"`
	WorkOrderId int     `json:"workOrderId" dc:"工单id"`
	Remark      string  `json:"remark" dc:"备注"`
	HandleType  string  `json:"handleType" dc:"处理方式"`
	Price       float64 `json:"price" dc:"金额"`
}

type BsReportDetailDeleteInp struct {
	Id any `json:"id" v:"required#主键不能为空" dc:"主键"`
}
