// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/bsservice"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	BsService = cBsService{}
)

type cBsService struct{}

// List 查看服务项目列表
func (c *cBsService) List(ctx context.Context, req *bsservice.ListReq) (res *bsservice.ListRes, err error) {
	list, totalCount, err := service.SysBsService().List(ctx, &req.BsServiceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.BsServiceListModel{}
	}

	res = new(bsservice.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新服务项目
func (c *cBsService) Edit(ctx context.Context, req *bsservice.EditReq) (res *bsservice.EditRes, err error) {
	err = service.SysBsService().Edit(ctx, &req.BsServiceEditInp)
	return
}

// View 获取指定服务项目信息
func (c *cBsService) View(ctx context.Context, req *bsservice.ViewReq) (res *bsservice.ViewRes, err error) {
	data, err := service.SysBsService().View(ctx, &req.BsServiceViewInp)
	if err != nil {
		return
	}

	res = new(bsservice.ViewRes)
	res.BsServiceViewModel = data
	return
}

// Delete 删除服务项目
func (c *cBsService) Delete(ctx context.Context, req *bsservice.DeleteReq) (res *bsservice.DeleteRes, err error) {
	err = service.SysBsService().Delete(ctx, &req.BsServiceDeleteInp)
	return
}

// TreeOption 获取服务项目关系树选项
func (c *cBsService) TreeOption(ctx context.Context, req *bsservice.TreeOptionReq) (res *bsservice.TreeOptionRes, err error) {
	data, err := service.SysBsService().TreeOption(ctx)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		res = (*bsservice.TreeOptionRes)(&data)
	} else {
		temp := make(bsservice.TreeOptionRes, 0)
		res = &temp
	}
	return
}