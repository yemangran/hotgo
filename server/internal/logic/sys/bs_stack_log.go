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
	"hotgo/internal/library/hgorm"
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

type sSysBsStackLog struct{}

func NewSysBsStackLog() *sSysBsStackLog {
	return &sSysBsStackLog{}
}

func init() {
	service.RegisterSysBsStackLog(NewSysBsStackLog())
}

// Model 库存流水ORM模型
func (s *sSysBsStackLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BsStackLog.Ctx(ctx), option...)
}

// List 获取库存流水列表
func (s *sSysBsStackLog) List(ctx context.Context, in *sysin.BsStackLogListInp) (list []*sysin.BsStackLogListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.FieldsPrefix(dao.BsStackLog.Table(), sysin.BsStackLogListModel{})
	mod = mod.Fields(hgorm.JoinFields(ctx, sysin.BsStackLogListModel{}, &dao.BsItem, "bsItem"))
	mod = mod.Fields(hgorm.JoinFields(ctx, sysin.BsStackLogListModel{}, &dao.BsWorkOrder, "bsWorkOrder"))

	// 关联表字段
	mod = mod.LeftJoinOnFields(dao.BsItem.Table(), dao.BsStackLog.Columns().ItemId, "=", dao.BsItem.Columns().Id)
	mod = mod.LeftJoinOnFields(dao.BsWorkOrder.Table(), dao.BsStackLog.Columns().WorkOrderId, "=", dao.BsWorkOrder.Columns().Id)

	// 查询物品名称
	if in.BsItemName != "" {
		mod = mod.WherePrefixLike(dao.BsItem.Table(), dao.BsItem.Columns().Name, in.BsItemName)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.BsStackLog.Table() + "." + dao.BsStackLog.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取库存流水列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出库存流水
func (s *sSysBsStackLog) Export(ctx context.Context, in *sysin.BsStackLogListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.BsStackLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出库存流水-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.BsStackLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增库存流水
func (s *sSysBsStackLog) Edit(ctx context.Context, in *sysin.BsStackLogEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.BsStackLogUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改库存流水失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.BsStackLogInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增库存流水失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除库存流水
func (s *sSysBsStackLog) Delete(ctx context.Context, in *sysin.BsStackLogDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除库存流水失败，请稍后重试！")
		return
	}
	return
}

// View 获取库存流水指定信息
func (s *sSysBsStackLog) View(ctx context.Context, in *sysin.BsStackLogViewInp) (res *sysin.BsStackLogViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取库存流水信息，请稍后重试！")
		return
	}
	return
}