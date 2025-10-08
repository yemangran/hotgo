// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BsStackLogUpdateFields 修改库存流水字段过滤
type BsStackLogUpdateFields struct {
	ItemId      int         `json:"itemId"      dc:"物品Id"`
	Qty         int         `json:"qty"         dc:"数量"`
	WorkOrderId int         `json:"workOrderId" dc:"工单Id"`
	CreateTime  *gtime.Time `json:"createTime"  dc:"创建时间"`
	Remark      string      `json:"remark"      dc:"备注"`
}

// BsStackLogInsertFields 新增库存流水字段过滤
type BsStackLogInsertFields struct {
	ItemId      int         `json:"itemId"      dc:"物品Id"`
	Qty         int         `json:"qty"         dc:"数量"`
	WorkOrderId int         `json:"workOrderId" dc:"工单Id"`
	CreateTime  *gtime.Time `json:"createTime"  dc:"创建时间"`
	Remark      string      `json:"remark"      dc:"备注"`
}

// BsStackLogEditInp 修改/新增库存流水
type BsStackLogEditInp struct {
	entity.BsStackLog
}

func (in *BsStackLogEditInp) Filter(ctx context.Context) (err error) {
	// 验证物品Id
	if err := g.Validator().Rules("required").Data(in.ItemId).Messages("物品Id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证数量
	if err := g.Validator().Rules("required").Data(in.Qty).Messages("数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证工单Id
	if err := g.Validator().Rules("required").Data(in.WorkOrderId).Messages("工单Id不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type BsStackLogEditModel struct{}

// BsStackLogDeleteInp 删除库存流水
type BsStackLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsStackLogDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BsStackLogDeleteModel struct{}

// BsStackLogViewInp 获取指定库存流水信息
type BsStackLogViewInp struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsStackLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BsStackLogViewModel struct {
	entity.BsStackLog
}

// BsStackLogListInp 获取库存流水列表
type BsStackLogListInp struct {
	form.PageReq
	BsItemName string `json:"bsItemName" dc:"物品名称"`
}

func (in *BsStackLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type BsStackLogListModel struct {
	Id             int64       `json:"id"             dc:"主键"`
	Qty            int         `json:"qty"            dc:"数量"`
	CreateTime     *gtime.Time `json:"createTime"     dc:"创建时间"`
	BsItemName     string      `json:"bsItemName"     dc:"物品名称"`
	BsItemCategory int         `json:"bsItemCategory" dc:"物品类别"`
	BsItemSpec     string      `json:"bsItemSpec"     dc:"规格型号"`
	BsWorkOrderId  int64       `json:"bsWorkOrderId"  dc:"工单"`
}

// BsStackLogExportModel 导出库存流水
type BsStackLogExportModel struct {
	Qty            int         `json:"qty"            dc:"数量"`
	CreateTime     *gtime.Time `json:"createTime"     dc:"创建时间"`
	BsItemName     string      `json:"bsItemName"     dc:"物品名称"`
	BsItemCategory int         `json:"bsItemCategory" dc:"物品类别"`
	BsItemSpec     string      `json:"bsItemSpec"     dc:"规格型号"`
}