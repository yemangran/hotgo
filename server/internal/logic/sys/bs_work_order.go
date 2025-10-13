// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysBsWorkOrder struct{}

func NewSysBsWorkOrder() *sSysBsWorkOrder {
	return &sSysBsWorkOrder{}
}

func init() {
	service.RegisterSysBsWorkOrder(NewSysBsWorkOrder())
}

// Model 工单管理ORM模型
func (s *sSysBsWorkOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BsWorkOrder.Ctx(ctx), option...)
}

// List 获取工单管理列表
func (s *sSysBsWorkOrder) List(ctx context.Context, in *sysin.BsWorkOrderListInp) (list []*sysin.BsWorkOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.BsWorkOrderListModel{})

	// 查询创建时间
	if len(in.CreateTime) == 2 {
		mod = mod.WhereBetween(dao.BsWorkOrder.Columns().CreateTime, in.CreateTime[0], in.CreateTime[1])
	}

	// 查询客户名称
	if in.CustomerName != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().CustomerName, in.CustomerName)
	}

	// 查询客户地址
	if in.CustomerAddress != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().CustomerAddress, in.CustomerAddress)
	}

	// 查询客户联系方式
	if in.CustomerContact != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().CustomerContact, in.CustomerContact)
	}

	// 查询客户对接人
	if in.CustomerPerson != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().CustomerPerson, in.CustomerPerson)
	}

	// 查询发起类型
	if in.SendType > 0 {
		mod = mod.Where(dao.BsWorkOrder.Columns().SendType, in.SendType)
	}

	// 查询产品类型
	if in.ProductType > 0 {
		mod = mod.Where(dao.BsWorkOrder.Columns().ProductType, in.ProductType)
	}

	// 查询工单状态
	if in.Status > 0 {
		mod = mod.Where(dao.BsWorkOrder.Columns().Status, in.Status)
	}

	// 查询开票类型
	if in.InvoiceType > 0 {
		mod = mod.Where(dao.BsWorkOrder.Columns().InvoiceType, in.InvoiceType)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.BsWorkOrder.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取工单管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出工单管理
func (s *sSysBsWorkOrder) Export(ctx context.Context, in *sysin.BsWorkOrderListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.BsWorkOrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出工单管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.BsWorkOrderExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增工单管理
func (s *sSysBsWorkOrder) Edit(ctx context.Context, in *sysin.BsWorkOrderEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var member = contexts.GetUser(ctx)
		fmt.Printf("当前访问用户信息：%+v\n", member)
		in.UserId = int(member.Id)

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.BsWorkOrderUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改工单管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.BsWorkOrderInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增工单管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除工单管理
func (s *sSysBsWorkOrder) Delete(ctx context.Context, in *sysin.BsWorkOrderDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除工单管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取工单管理指定信息
func (s *sSysBsWorkOrder) View(ctx context.Context, in *sysin.BsWorkOrderViewInp) (res *sysin.BsWorkOrderViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取工单管理信息失败，请稍后重试！")
		return
	}
	//查询用户信息
	member := entity.AdminMember{}
	if err = dao.AdminMember.Ctx(ctx).Where("id", res.DispatchId).Scan(&member); err != nil {
		err = gerror.Wrap(err, "获取处理人信息失败，请稍后重试！")
		return
	}
	res.DispatchName = member.Username
	return
}

// Status 更新工单管理状态
func (s *sSysBsWorkOrder) Status(ctx context.Context, in *sysin.BsWorkOrderStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.BsWorkOrder.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新工单管理状态失败，请稍后重试！")
		return
	}
	return
}
