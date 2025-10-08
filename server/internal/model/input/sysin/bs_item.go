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
)

// BsItemUpdateFields 修改仓库物品字段过滤
type BsItemUpdateFields struct {
	Name     string  `json:"name"     dc:"物品名称"`
	Category int     `json:"category" dc:"物品类别"`
	Spec     string  `json:"spec"     dc:"规格型号"`
	Price    float64 `json:"price"    dc:"单价"`
}

// BsItemInsertFields 新增仓库物品字段过滤
type BsItemInsertFields struct {
	Name     string  `json:"name"     dc:"物品名称"`
	Category int     `json:"category" dc:"物品类别"`
	Spec     string  `json:"spec"     dc:"规格型号"`
	Price    float64 `json:"price"    dc:"单价"`
}

// BsItemEditInp 修改/新增仓库物品
type BsItemEditInp struct {
	entity.BsItem
}

func (in *BsItemEditInp) Filter(ctx context.Context) (err error) {
	// 验证物品名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("物品名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证物品类别
	if err := g.Validator().Rules("required").Data(in.Category).Messages("物品类别不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:0,1").Data(in.Category).Messages("物品类别值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证单价
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("单价最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type BsItemEditModel struct{}

// BsItemDeleteInp 删除仓库物品
type BsItemDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsItemDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BsItemDeleteModel struct{}

// BsItemViewInp 获取指定仓库物品信息
type BsItemViewInp struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsItemViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BsItemViewModel struct {
	entity.BsItem
}

// BsItemListInp 获取仓库物品列表
type BsItemListInp struct {
	form.PageReq
	Name     string `json:"name"     dc:"物品名称"`
	Category int    `json:"category" dc:"物品类别"`
	Spec     string `json:"spec"     dc:"规格型号"`
}

func (in *BsItemListInp) Filter(ctx context.Context) (err error) {
	return
}

type BsItemListModel struct {
	Id       int64   `json:"id"       dc:"主键"`
	Name     string  `json:"name"     dc:"物品名称"`
	Category int     `json:"category" dc:"物品类别"`
	Spec     string  `json:"spec"     dc:"规格型号"`
	Price    float64 `json:"price"    dc:"单价"`
}

// BsItemExportModel 导出仓库物品
type BsItemExportModel struct {
	Name     string  `json:"name"     dc:"物品名称"`
	Category int     `json:"category" dc:"物品类别"`
	Spec     string  `json:"spec"     dc:"规格型号"`
	Price    float64 `json:"price"    dc:"单价"`
}