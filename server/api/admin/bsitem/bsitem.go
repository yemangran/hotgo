// Package bsitem
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package bsitem

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询仓库物品列表
type ListReq struct {
	g.Meta `path:"/bsItem/list" method:"get" tags:"仓库物品" summary:"获取仓库物品列表"`
	sysin.BsItemListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BsItemListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出仓库物品列表
type ExportReq struct {
	g.Meta `path:"/bsItem/export" method:"get" tags:"仓库物品" summary:"导出仓库物品列表"`
	sysin.BsItemListInp
}

type ExportRes struct{}

// ViewReq 获取仓库物品指定信息
type ViewReq struct {
	g.Meta `path:"/bsItem/view" method:"get" tags:"仓库物品" summary:"获取仓库物品指定信息"`
	sysin.BsItemViewInp
}

type ViewRes struct {
	*sysin.BsItemViewModel
}

// EditReq 修改/新增仓库物品
type EditReq struct {
	g.Meta `path:"/bsItem/edit" method:"post" tags:"仓库物品" summary:"修改/新增仓库物品"`
	sysin.BsItemEditInp
}

type EditRes struct{}

// DeleteReq 删除仓库物品
type DeleteReq struct {
	g.Meta `path:"/bsItem/delete" method:"post" tags:"仓库物品" summary:"删除仓库物品"`
	sysin.BsItemDeleteInp
}

type DeleteRes struct{}