// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/utility/tree"
)

type sSysBsService struct{}

func NewSysBsService() *sSysBsService {
	return &sSysBsService{}
}

func init() {
	service.RegisterSysBsService(NewSysBsService())
}

// Model 服务项目ORM模型
func (s *sSysBsService) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BsService.Ctx(ctx), option...)
}

// List 获取服务项目列表
func (s *sSysBsService) List(ctx context.Context, in *sysin.BsServiceListInp) (list []*sysin.BsServiceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.BsServiceListModel{})

	// 查询主键
	if in.Id > 0 {
		mod = mod.Where(dao.BsService.Columns().Id, in.Id)
	}

	// 查询父键
	if in.Pid > 0 {
		mod = mod.Where(dao.BsService.Columns().Pid, in.Pid)
	}

	// 树形列表判断是否需要分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.Order(dao.BsService.Columns().OrderNum)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取服务项目列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增服务项目
func (s *sSysBsService) Edit(ctx context.Context, in *sysin.BsServiceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		newPid, newLevel, newTree, err := hgorm.AutoUpdateTree(ctx, &dao.BsService, int64(in.Id), in.Pid)
		in.Tree = newTree
		if err != nil {
			return err
		}
		in.Pid = newPid
		in.Level = int64(newLevel)
		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.BsServiceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改服务项目失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.BsServiceInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增服务项目失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除服务项目
func (s *sSysBsService) Delete(ctx context.Context, in *sysin.BsServiceDeleteInp) (err error) {
	count, err := dao.BsService.Ctx(ctx).Where(dao.BsService.Columns().Pid, in.Id).Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if count > 0 {
		return gerror.New("请先删除该服务项目下的所有下级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除服务项目失败，请稍后重试！")
		return
	}
	return
}

// View 获取服务项目指定信息
func (s *sSysBsService) View(ctx context.Context, in *sysin.BsServiceViewInp) (res *sysin.BsServiceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取服务项目信息，请稍后重试！")
		return
	}
	return
}

// TreeOption 获取服务项目关系树选项
func (s *sSysBsService) TreeOption(ctx context.Context) (nodes []tree.Node, err error) {
	var models []*sysin.BsServiceTreeOption
	if err = s.Model(ctx).Fields(sysin.BsServiceTreeOption{}).OrderAsc(dao.BsService.Columns().Pid).OrderDesc(dao.BsService.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取服务项目关系树选项失败！")
		return
	}
	nodes = make([]tree.Node, len(models))
	for i, v := range models {
		nodes[i] = v
	}
	return tree.ListToTree(0, nodes)
}
