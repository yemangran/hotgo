package sys

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysBsReportDetail struct{}

// Delete implements service.ISysBsReportDetail.
func (s *sSysBsReportDetail) Delete(ctx context.Context, in *sysin.BsReportDetailDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除工单明细失败")
		return
	}
	return
}

// List implements service.ISysBsReportDetail.
func (s *sSysBsReportDetail) List(ctx context.Context, in *sysin.BsReportDetailListInp) (list []*sysin.BsReportDetailListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.BsReportDetailListModel{})
	// 查询明细的id
	if in.Id > 0 {
		mod = mod.Where(dao.BsReportDetail.Columns().Id, in.Id)
	}
	// 排序
	mod = mod.OrderDesc(dao.BsReportDetail.Columns().Id)
	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取工单明细列表失败")
		return
	}
	return
}

// Model implements service.ISysBsReportDetail.
func (s *sSysBsReportDetail) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.BsReportDetail.Ctx(ctx), option...)
}

func NewSysBsReportDetail() *sSysBsReportDetail {
	return &sSysBsReportDetail{}
}

func init() {
	service.RegisterSysBsReportDetail(NewSysBsReportDetail())
}
