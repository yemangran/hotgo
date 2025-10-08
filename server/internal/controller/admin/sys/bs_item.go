// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/bsitem"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	BsItem = cBsItem{}
)

type cBsItem struct{}

// List 查看仓库物品列表
func (c *cBsItem) List(ctx context.Context, req *bsitem.ListReq) (res *bsitem.ListRes, err error) {
	list, totalCount, err := service.SysBsItem().List(ctx, &req.BsItemListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.BsItemListModel{}
	}

	res = new(bsitem.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出仓库物品列表
func (c *cBsItem) Export(ctx context.Context, req *bsitem.ExportReq) (res *bsitem.ExportRes, err error) {
	err = service.SysBsItem().Export(ctx, &req.BsItemListInp)
	return
}

// Edit 更新仓库物品
func (c *cBsItem) Edit(ctx context.Context, req *bsitem.EditReq) (res *bsitem.EditRes, err error) {
	err = service.SysBsItem().Edit(ctx, &req.BsItemEditInp)
	return
}

// View 获取指定仓库物品信息
func (c *cBsItem) View(ctx context.Context, req *bsitem.ViewReq) (res *bsitem.ViewRes, err error) {
	data, err := service.SysBsItem().View(ctx, &req.BsItemViewInp)
	if err != nil {
		return
	}

	res = new(bsitem.ViewRes)
	res.BsItemViewModel = data
	return
}

// Delete 删除仓库物品
func (c *cBsItem) Delete(ctx context.Context, req *bsitem.DeleteReq) (res *bsitem.DeleteRes, err error) {
	err = service.SysBsItem().Delete(ctx, &req.BsItemDeleteInp)
	return
}