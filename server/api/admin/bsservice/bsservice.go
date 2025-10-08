// Package bsservice
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package bsservice

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询服务项目列表
type ListReq struct {
	g.Meta `path:"/bsService/list" method:"get" tags:"服务项目" summary:"获取服务项目列表"`
	sysin.BsServiceListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.BsServiceListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取服务项目指定信息
type ViewReq struct {
	g.Meta `path:"/bsService/view" method:"get" tags:"服务项目" summary:"获取服务项目指定信息"`
	sysin.BsServiceViewInp
}

type ViewRes struct {
	*sysin.BsServiceViewModel
}

// EditReq 修改/新增服务项目
type EditReq struct {
	g.Meta `path:"/bsService/edit" method:"post" tags:"服务项目" summary:"修改/新增服务项目"`
	sysin.BsServiceEditInp
}

type EditRes struct{}

// DeleteReq 删除服务项目
type DeleteReq struct {
	g.Meta `path:"/bsService/delete" method:"post" tags:"服务项目" summary:"删除服务项目"`
	sysin.BsServiceDeleteInp
}

type DeleteRes struct{}

// TreeOptionReq 获取服务项目关系树选项
type TreeOptionReq struct {
	g.Meta `path:"/bsService/treeOption" method:"get" tags:"服务项目" summary:"获取服务项目关系树选项"`
}

type TreeOptionRes []tree.Node