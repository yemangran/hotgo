// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/bsworkorder"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	BsWorkOrder = cBsWorkOrder{}
)

type cBsWorkOrder struct{}

// List 查看工单管理列表
func (c *cBsWorkOrder) List(ctx context.Context, req *bsworkorder.ListReq) (res *bsworkorder.ListRes, err error) {
	list, totalCount, err := service.SysBsWorkOrder().List(ctx, &req.BsWorkOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.BsWorkOrderListModel{}
	}

	res = new(bsworkorder.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出工单管理列表
func (c *cBsWorkOrder) Export(ctx context.Context, req *bsworkorder.ExportReq) (res *bsworkorder.ExportRes, err error) {
	err = service.SysBsWorkOrder().Export(ctx, &req.BsWorkOrderListInp)
	return
}

// Edit 更新工单管理
func (c *cBsWorkOrder) Edit(ctx context.Context, req *bsworkorder.EditReq) (res *bsworkorder.EditRes, err error) {
	err = service.SysBsWorkOrder().Edit(ctx, &req.BsWorkOrderEditInp)
	return
}

// View 获取指定工单管理信息
func (c *cBsWorkOrder) View(ctx context.Context, req *bsworkorder.ViewReq) (res *bsworkorder.ViewRes, err error) {
	data, err := service.SysBsWorkOrder().View(ctx, &req.BsWorkOrderViewInp)
	if err != nil {
		return
	}

	res = new(bsworkorder.ViewRes)
	res.BsWorkOrderViewModel = data
	return
}

// Delete 删除工单管理
func (c *cBsWorkOrder) Delete(ctx context.Context, req *bsworkorder.DeleteReq) (res *bsworkorder.DeleteRes, err error) {
	err = service.SysBsWorkOrder().Delete(ctx, &req.BsWorkOrderDeleteInp)
	return
}

// Status 更新工单管理状态
func (c *cBsWorkOrder) Status(ctx context.Context, req *bsworkorder.StatusReq) (res *bsworkorder.StatusRes, err error) {
	err = service.SysBsWorkOrder().Status(ctx, &req.BsWorkOrderStatusInp)
	return
}