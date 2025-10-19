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
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
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

// Process 处理工单
func (s *sSysBsWorkOrder) Process(ctx context.Context, in *sysin.BsWorkOrderProcessInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 1. 验证工单ID
		if in.Id <= 0 {
			return gerror.New("工单ID不能为空")
		}

		// 2. 验证明细数据
		if len(in.ServiceList) == 0 {
			return gerror.New("工单明细不能为空")
		}

		// 3. 更新工单主表信息
		if _, err = tx.Ctx(ctx).Model(dao.BsWorkOrder.Table()).
			Fields(sysin.BsWorkOrderUpdateFields{}).
			WherePri(in.Id).
			Data(in.BsWorkOrder).
			Update(); err != nil {
			return gerror.Wrap(err, "更新工单信息失败，请稍后重试！")
		}

		// 4. 计算总金额，并分类明细数据
		var (
			totalMoney    float64
			insertDetails []g.Map                 // 待新增的明细
			updateDetails []entity.BsReportDetail // 待更新的明细
		)

		for _, detail := range in.ServiceList {
			// 统计总金额
			totalMoney += detail.Price

			// 确保 workOrderId 正确
			detail.WorkOrderId = int(in.Id)

			// 根据是否有 ID 分类
			if detail.Id > 0 {
				// 添加到更新列表
				updateDetails = append(updateDetails, detail)
			} else {
				// 添加到新增列表
				insertDetails = append(insertDetails, g.Map{
					dao.BsReportDetail.Columns().ServiceId:   detail.ServiceId,
					dao.BsReportDetail.Columns().WorkOrderId: detail.WorkOrderId,
					dao.BsReportDetail.Columns().HandleType:  detail.HandleType,
					dao.BsReportDetail.Columns().Price:       detail.Price,
					dao.BsReportDetail.Columns().Remark:      detail.Remark,
				})
			}
		}

		// 5. 批量新增明细（如果有）
		if len(insertDetails) > 0 {
			if _, err = tx.Ctx(ctx).Model(dao.BsReportDetail.Table()).
				Data(insertDetails).
				Insert(); err != nil {
				return gerror.Wrap(err, "批量新增工单明细失败")
			}
		}

		// 6. 更新明细（逐条更新）
		for _, detail := range updateDetails {
			if _, err = tx.Ctx(ctx).Model(dao.BsReportDetail.Table()).
				WherePri(detail.Id).
				Data(g.Map{
					dao.BsReportDetail.Columns().ServiceId:   detail.ServiceId,
					dao.BsReportDetail.Columns().WorkOrderId: detail.WorkOrderId,
					dao.BsReportDetail.Columns().HandleType:  detail.HandleType,
					dao.BsReportDetail.Columns().Price:       detail.Price,
					dao.BsReportDetail.Columns().Remark:      detail.Remark,
				}).
				Update(); err != nil {
				return gerror.Wrapf(err, "更新工单明细失败，明细ID:%d", detail.Id)
			}
		}

		// 7. 更新工单总金额和状态
		//工单状态 1待分派 2处理中 3已完结
		if _, err = tx.Ctx(ctx).Model(dao.BsWorkOrder.Table()).
			WherePri(in.Id).
			Data(g.Map{
				dao.BsWorkOrder.Columns().TotalMoney: totalMoney,
				dao.BsWorkOrder.Columns().Status:     3, // 假设 2 表示已处理，根据实际业务调整
			}).
			Update(); err != nil {
			return gerror.Wrap(err, "更新工单总金额失败")
		}

		g.Log().Infof(ctx, "工单处理成功，工单ID:%d, 总金额:%.2f, 新增明细:%d条, 更新明细:%d条",
			in.Id, totalMoney, len(insertDetails), len(updateDetails))

		return nil
	})
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

	// 查询设备型号
	if in.DeviceModel != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().DeviceModel, in.DeviceModel)
	}

	// 查询部件名称
	if in.ComponentName != "" {
		mod = mod.WhereLike(dao.BsWorkOrder.Columns().ComponentName, in.ComponentName)
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

// GenerateReport 生成工单报告
func (s *sSysBsWorkOrder) GenerateReport(ctx context.Context, id int64) (html string, err error) {
	// 1. 获取工单基本信息
	var workOrder *entity.BsWorkOrder
	if err = s.Model(ctx).WherePri(id).Scan(&workOrder); err != nil {
		err = gerror.Wrap(err, "获取工单信息失败")
		return
	}

	if workOrder == nil {
		err = gerror.New("工单不存在")
		return
	}

	// 2. 获取处理人信息
	var member entity.AdminMember
	dispatchName := "未分配"
	if workOrder.DispatchId > 0 {
		if err = dao.AdminMember.Ctx(ctx).Where("id", workOrder.DispatchId).Scan(&member); err == nil {
			dispatchName = member.Username
		}
	}

	// 3. 获取工单明细
	var details []*entity.BsReportDetail
	if err = dao.BsReportDetail.Ctx(ctx).
		Where(dao.BsReportDetail.Columns().WorkOrderId, id).
		Scan(&details); err != nil {
		err = gerror.Wrap(err, "获取工单明细失败")
		return
	}

	// 4. 构建模板数据
	type DetailData struct {
		No             int // 序号
		ServiceName    string
		HandleTypeName string
		Price          float64
	}

	var detailList []DetailData
	for idx, detail := range details {
		// 获取服务项目名称
		var service entity.BsService
		serviceName := "未知服务"
		if err = dao.BsService.Ctx(ctx).Where("id", detail.ServiceId).Scan(&service); err == nil {
			serviceName = service.Name
		}

		// 转换处理类型
		handleTypeName := "未知"
		switch detail.HandleType {
		case "1":
			handleTypeName = "维修"
		case "2":
			handleTypeName = "更换"
		case "3":
			handleTypeName = "其他"
		}

		detailList = append(detailList, DetailData{
			No:             idx + 1, // 直接计算序号
			ServiceName:    serviceName,
			HandleTypeName: handleTypeName,
			Price:          detail.Price,
		})
	}

	//TODO: 添加字典项对应的展示

	// 5. 准备模板数据
	templateData := g.Map{
		"Id":                 workOrder.Id,
		"CustomerName":       workOrder.CustomerName,
		"CustomerContact":    workOrder.CustomerContact,
		"CustomerAddress":    workOrder.CustomerAddress,
		"CustomerPerson":     workOrder.CustomerPerson,
		"ProblemDescription": workOrder.ProblemDescription,
		"DispatchName":       dispatchName,
		"CreateTime":         workOrder.CreateTime.Format("Y-m-d H:i:s"),
		"DefectDescription":  workOrder.DefectDescription,
		"Details":            detailList,
		"TotalMoney":         workOrder.TotalMoney,
		"DeviceModel":        workOrder.DeviceModel,
		"ComponentName":      workOrder.ComponentName,
		"GenerateTime":       gtime.Now().Format("Y-m-d H:i:s"),
	}

	// 6. 读取并渲染模板
	tmplPath := gfile.Join(gfile.Pwd(), "resource/template/report/work_order.html")
	if !gfile.Exists(tmplPath) {
		err = gerror.Newf("模板文件不存在: %s", tmplPath)
		return
	}

	// 读取模板内容
	tmplContent := gfile.GetContents(tmplPath)
	if tmplContent == "" {
		err = gerror.New("模板文件为空")
		return
	}

	// 解析并执行模板
	html, err = gview.New().ParseContent(ctx, tmplContent, templateData)
	if err != nil {
		err = gerror.Wrap(err, "渲染模板失败")
		return
	}

	return
}
