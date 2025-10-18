package sys

import (
	"context"
	"hotgo/api/admin/bsreportdetail"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
)

var (
	BsReportDetail = cBsReportDetail{}
)

type cBsReportDetail struct{}

func (c *cBsReportDetail) List(ctx context.Context, req *bsreportdetail.ListReq) (res *bsreportdetail.ListRes, err error) {
	list, totalCount, err := service.SysBsReportDetail().List(ctx, &req.BsReportDetailListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*sysin.BsReportDetailListModel{}
	}
	res = new(bsreportdetail.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *cBsReportDetail) Delete(ctx context.Context, req *bsreportdetail.DeleteReq) (res *bsreportdetail.DeleteRes, err error) {
	err = service.SysBsReportDetail().Delete(ctx, &req.BsReportDetailDeleteInp)
	return
}
