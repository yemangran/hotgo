// Package bsstacklog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package bsstacklog

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询库存流水列表
type ListReq struct {
	g.Meta `path:"/bsStackLog/list" method:"get" tags:"库存流水" summary:"获取库存流水列表"`
	sysin.BsStackLogListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BsStackLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出库存流水列表
type ExportReq struct {
	g.Meta `path:"/bsStackLog/export" method:"get" tags:"库存流水" summary:"导出库存流水列表"`
	sysin.BsStackLogListInp
}

type ExportRes struct{}

// ViewReq 获取库存流水指定信息
type ViewReq struct {
	g.Meta `path:"/bsStackLog/view" method:"get" tags:"库存流水" summary:"获取库存流水指定信息"`
	sysin.BsStackLogViewInp
}

type ViewRes struct {
	*sysin.BsStackLogViewModel
}

// EditReq 修改/新增库存流水
type EditReq struct {
	g.Meta `path:"/bsStackLog/edit" method:"post" tags:"库存流水" summary:"修改/新增库存流水"`
	sysin.BsStackLogEditInp
}

type EditRes struct{}

// DeleteReq 删除库存流水
type DeleteReq struct {
	g.Meta `path:"/bsStackLog/delete" method:"post" tags:"库存流水" summary:"删除库存流水"`
	sysin.BsStackLogDeleteInp
}

type DeleteRes struct{}