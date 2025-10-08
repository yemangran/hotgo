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
	"hotgo/internal/library/hgorm/handler"
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

type sSysBsItem struct{}

func NewSysBsItem() *sSysBsItem {
	return &sSysBsItem{}
}

func init() {
	service.RegisterSysBsItem(NewSysBsItem())
}

// Model 仓库物品ORM模型
func (s *sSysBsItem) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BsItem.Ctx(ctx), option...)
}

// List 获取仓库物品列表
func (s *sSysBsItem) List(ctx context.Context, in *sysin.BsItemListInp) (list []*sysin.BsItemListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.BsItemListModel{})

	// 查询物品名称
	if in.Name != "" {
		mod = mod.WhereLike(dao.BsItem.Columns().Name, in.Name)
	}

	// 查询物品类别
	if in.Category > 0 {
		mod = mod.Where(dao.BsItem.Columns().Category, in.Category)
	}

	// 查询规格型号
	if in.Spec != "" {
		mod = mod.WhereLike(dao.BsItem.Columns().Spec, in.Spec)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.BsItem.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取仓库物品列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出仓库物品
func (s *sSysBsItem) Export(ctx context.Context, in *sysin.BsItemListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.BsItemExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出仓库物品-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.BsItemExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增仓库物品
func (s *sSysBsItem) Edit(ctx context.Context, in *sysin.BsItemEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.BsItemUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改仓库物品失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.BsItemInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增仓库物品失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除仓库物品
func (s *sSysBsItem) Delete(ctx context.Context, in *sysin.BsItemDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除仓库物品失败，请稍后重试！")
		return
	}
	return
}

// View 获取仓库物品指定信息
func (s *sSysBsItem) View(ctx context.Context, in *sysin.BsItemViewInp) (res *sysin.BsItemViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取仓库物品信息，请稍后重试！")
		return
	}
	return
}