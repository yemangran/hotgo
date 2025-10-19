// Package bsworkorder
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package bsworkorder

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询工单管理列表
type ListReq struct {
	g.Meta `path:"/bsWorkOrder/list" method:"get" tags:"工单管理" summary:"获取工单管理列表"`
	sysin.BsWorkOrderListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BsWorkOrderListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出工单管理列表
type ExportReq struct {
	g.Meta `path:"/bsWorkOrder/export" method:"get" tags:"工单管理" summary:"导出工单管理列表"`
	sysin.BsWorkOrderListInp
}

type ExportRes struct{}

// ViewReq 获取工单管理指定信息
type ViewReq struct {
	g.Meta `path:"/bsWorkOrder/view" method:"get" tags:"工单管理" summary:"获取工单管理指定信息"`
	sysin.BsWorkOrderViewInp
}

type ViewRes struct {
	*sysin.BsWorkOrderViewModel
}

// EditReq 修改/新增工单管理
type EditReq struct {
	g.Meta `path:"/bsWorkOrder/edit" method:"post" tags:"工单管理" summary:"修改/新增工单管理"`
	sysin.BsWorkOrderEditInp
}

type EditRes struct{}

// DeleteReq 删除工单管理
type DeleteReq struct {
	g.Meta `path:"/bsWorkOrder/delete" method:"post" tags:"工单管理" summary:"删除工单管理"`
	sysin.BsWorkOrderDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新工单管理状态
type StatusReq struct {
	g.Meta `path:"/bsWorkOrder/status" method:"post" tags:"工单管理" summary:"更新工单管理状态"`
	sysin.BsWorkOrderStatusInp
}

type StatusRes struct{}

// ProcessReq 处理工单
type ProcessReq struct {
	g.Meta `path:"/bsWorkOrder/process" method:"post" tags:"工单管理" summary:"处理工单"`
	sysin.BsWorkOrderProcessInp
}

type ProcessRes struct{}

// GenerateReportReq 生成工单报告
type GenerateReportReq struct {
	g.Meta `path:"/bsWorkOrder/generateReport" method:"get" tags:"工单管理" summary:"生成工单报告"`
	Id     int64 `json:"id" v:"required#工单ID不能为空" dc:"工单ID"`
}

type GenerateReportRes struct {
	Html string `json:"html" dc:"报告HTML内容"`
}
