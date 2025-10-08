// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/bsstacklog"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	BsStackLog = cBsStackLog{}
)

type cBsStackLog struct{}

// List 查看库存流水列表
func (c *cBsStackLog) List(ctx context.Context, req *bsstacklog.ListReq) (res *bsstacklog.ListRes, err error) {
	list, totalCount, err := service.SysBsStackLog().List(ctx, &req.BsStackLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.BsStackLogListModel{}
	}

	res = new(bsstacklog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出库存流水列表
func (c *cBsStackLog) Export(ctx context.Context, req *bsstacklog.ExportReq) (res *bsstacklog.ExportRes, err error) {
	err = service.SysBsStackLog().Export(ctx, &req.BsStackLogListInp)
	return
}

// Edit 更新库存流水
func (c *cBsStackLog) Edit(ctx context.Context, req *bsstacklog.EditReq) (res *bsstacklog.EditRes, err error) {
	err = service.SysBsStackLog().Edit(ctx, &req.BsStackLogEditInp)
	return
}

// View 获取指定库存流水信息
func (c *cBsStackLog) View(ctx context.Context, req *bsstacklog.ViewReq) (res *bsstacklog.ViewRes, err error) {
	data, err := service.SysBsStackLog().View(ctx, &req.BsStackLogViewInp)
	if err != nil {
		return
	}

	res = new(bsstacklog.ViewRes)
	res.BsStackLogViewModel = data
	return
}

// Delete 删除库存流水
func (c *cBsStackLog) Delete(ctx context.Context, req *bsstacklog.DeleteReq) (res *bsstacklog.DeleteRes, err error) {
	err = service.SysBsStackLog().Delete(ctx, &req.BsStackLogDeleteInp)
	return
}