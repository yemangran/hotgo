// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BsWorkOrderUpdateFields 修改工单管理字段过滤
type BsWorkOrderUpdateFields struct {
	UserId             int         `json:"userId"             dc:"用户Id"`
	CreateTime         *gtime.Time `json:"createTime"         dc:"创建时间"`
	CustomerName       string      `json:"customerName"       dc:"客户名称"`
	CustomerAddress    string      `json:"customerAddress"    dc:"客户地址"`
	CustomerContact    string      `json:"customerContact"    dc:"客户联系方式"`
	CustomerPerson     string      `json:"customerPerson"     dc:"客户对接人"`
	ProblemDescription string      `json:"problemDescription" dc:"问题描述"`
	DefectDescription  string      `json:"defectDescription"  dc:"检测描述"`
	DispatchId         int         `json:"dispatchId"         dc:"分配处理人Id"`
	SendType           int         `json:"sendType"           dc:"发起类型"`
	AcceptType         int         `json:"acceptType"         dc:"收取方式"`
	ProductType        int         `json:"productType"        dc:"产品类型"`
	Status             int         `json:"status"             dc:"工单状态"`
	SuggestSolution    string      `json:"suggestSolution"    dc:"建议处理方案(多选)"`
	ActualSolution     string      `json:"actualSolution"     dc:"实际处理方案(多选)"`
	TotalMoney         float64     `json:"totalMoney"         dc:"总金额"`
	InvoiceType        int         `json:"invoiceType"        dc:"开票类型"`
	Remark             string      `json:"remark"             dc:"备注"`
}

// BsWorkOrderInsertFields 新增工单管理字段过滤
type BsWorkOrderInsertFields struct {
	UserId             int         `json:"userId"             dc:"用户Id"`
	CreateTime         *gtime.Time `json:"createTime"         dc:"创建时间"`
	CustomerName       string      `json:"customerName"       dc:"客户名称"`
	CustomerAddress    string      `json:"customerAddress"    dc:"客户地址"`
	CustomerContact    string      `json:"customerContact"    dc:"客户联系方式"`
	CustomerPerson     string      `json:"customerPerson"     dc:"客户对接人"`
	ProblemDescription string      `json:"problemDescription" dc:"问题描述"`
	DefectDescription  string      `json:"defectDescription"  dc:"检测描述"`
	DispatchId         int         `json:"dispatchId"         dc:"分配处理人Id"`
	SendType           int         `json:"sendType"           dc:"发起类型"`
	AcceptType         int         `json:"acceptType"         dc:"收取方式"`
	ProductType        int         `json:"productType"        dc:"产品类型"`
	Status             int         `json:"status"             dc:"工单状态"`
	SuggestSolution    string      `json:"suggestSolution"    dc:"建议处理方案(多选)"`
	ActualSolution     string      `json:"actualSolution"     dc:"实际处理方案(多选)"`
	TotalMoney         float64     `json:"totalMoney"         dc:"总金额"`
	InvoiceType        int         `json:"invoiceType"        dc:"开票类型"`
	Remark             string      `json:"remark"             dc:"备注"`
}

// BsWorkOrderEditInp 修改/新增工单管理
type BsWorkOrderEditInp struct {
	entity.BsWorkOrder
}

func (in *BsWorkOrderEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户Id
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户Id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证客户名称
	if err := g.Validator().Rules("required").Data(in.CustomerName).Messages("客户名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证发起类型
	if err := g.Validator().Rules("required").Data(in.SendType).Messages("发起类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证收取方式
	if err := g.Validator().Rules("required").Data(in.AcceptType).Messages("收取方式不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证产品类型
	if err := g.Validator().Rules("required").Data(in.ProductType).Messages("产品类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证工单状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("工单状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("工单状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证开票类型
	if err := g.Validator().Rules("required").Data(in.InvoiceType).Messages("开票类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type BsWorkOrderEditModel struct{}

// BsWorkOrderDeleteInp 删除工单管理
type BsWorkOrderDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsWorkOrderDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BsWorkOrderDeleteModel struct{}

// BsWorkOrderViewInp 获取指定工单管理信息
type BsWorkOrderViewInp struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsWorkOrderViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BsWorkOrderViewModel struct {
	entity.BsWorkOrder
}

// BsWorkOrderListInp 获取工单管理列表
type BsWorkOrderListInp struct {
	form.PageReq
	CustomerName    string `json:"customerName"    dc:"客户名称"`
	CustomerAddress string `json:"customerAddress" dc:"客户地址"`
	CustomerContact string `json:"customerContact" dc:"客户联系方式"`
	CustomerPerson  string `json:"customerPerson"  dc:"客户对接人"`
	Status          int    `json:"status"          dc:"工单状态"`
}

func (in *BsWorkOrderListInp) Filter(ctx context.Context) (err error) {
	return
}

type BsWorkOrderListModel struct {
	Id              int64       `json:"id"              dc:"主键"`
	CreateTime      *gtime.Time `json:"createTime"      dc:"创建时间"`
	CustomerName    string      `json:"customerName"    dc:"客户名称"`
	CustomerAddress string      `json:"customerAddress" dc:"客户地址"`
	CustomerContact string      `json:"customerContact" dc:"客户联系方式"`
	CustomerPerson  string      `json:"customerPerson"  dc:"客户对接人"`
	SendType        int         `json:"sendType"        dc:"发起类型"`
	AcceptType      int         `json:"acceptType"      dc:"收取方式"`
	ProductType     int         `json:"productType"     dc:"产品类型"`
	Status          int         `json:"status"          dc:"工单状态"`
	TotalMoney      float64     `json:"totalMoney"      dc:"总金额"`
	InvoiceType     int         `json:"invoiceType"     dc:"开票类型"`
}

// BsWorkOrderExportModel 导出工单管理
type BsWorkOrderExportModel struct {
	CreateTime      *gtime.Time `json:"createTime"      dc:"创建时间"`
	CustomerName    string      `json:"customerName"    dc:"客户名称"`
	CustomerAddress string      `json:"customerAddress" dc:"客户地址"`
	CustomerContact string      `json:"customerContact" dc:"客户联系方式"`
	CustomerPerson  string      `json:"customerPerson"  dc:"客户对接人"`
	SendType        int         `json:"sendType"        dc:"发起类型"`
	AcceptType      int         `json:"acceptType"      dc:"收取方式"`
	ProductType     int         `json:"productType"     dc:"产品类型"`
	Status          int         `json:"status"          dc:"工单状态"`
	SuggestSolution string      `json:"suggestSolution" dc:"建议处理方案(多选)"`
	ActualSolution  string      `json:"actualSolution"  dc:"实际处理方案(多选)"`
	TotalMoney      float64     `json:"totalMoney"      dc:"总金额"`
	InvoiceType     int         `json:"invoiceType"     dc:"开票类型"`
}

// BsWorkOrderStatusInp 更新工单管理状态
type BsWorkOrderStatusInp struct {
	Id     int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
	Status int   `json:"status" dc:"状态"`
}

func (in *BsWorkOrderStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("主键不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type BsWorkOrderStatusModel struct{}