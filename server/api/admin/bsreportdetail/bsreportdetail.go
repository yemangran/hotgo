package bsreportdetail

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询工单明细列表
type ListReq struct {
	g.Meta `path:"/bsReportDetail/list" method:"get" tags:"工单管理" summary:"获取工单明细列表"`
	sysin.BsReportDetailListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BsReportDetailListModel `json:"list"   dc:"数据列表"`
}

// DeleteReq 删除工单明细
type DeleteReq struct {
	g.Meta `path:"/bsReportDetail/delete" method:"post" tags:"工单管理" summary:"删除工单明细"`
	sysin.BsReportDetailDeleteInp
}

type DeleteRes struct{}
