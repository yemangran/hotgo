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

	"hotgo/utility/tree"
)

// BsServiceUpdateFields 修改服务项目字段过滤
type BsServiceUpdateFields struct {
	Pid      int64   `json:"pid"      dc:"父键"`
	Level    int64   `json:"level"    dc:"关系树级别"`
	Tree     string  `json:"tree"     dc:"关系树"`
	Name     string  `json:"name"     dc:"服务名称"`
	Price    float64 `json:"price"    dc:"价格"`
	OrderNum int     `json:"orderNum" dc:"序号"`
}

// BsServiceInsertFields 新增服务项目字段过滤
type BsServiceInsertFields struct {
	Pid      int64   `json:"pid"      dc:"父键"`
	Level    int64   `json:"level"    dc:"关系树级别"`
	Tree     string  `json:"tree"     dc:"关系树"`
	Name     string  `json:"name"     dc:"服务名称"`
	Price    float64 `json:"price"    dc:"价格"`
	OrderNum int     `json:"orderNum" dc:"序号"`
}

// BsServiceEditInp 修改/新增服务项目
type BsServiceEditInp struct {
	entity.BsService
}

func (in *BsServiceEditInp) Filter(ctx context.Context) (err error) {
	// 验证服务名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("服务名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证价格
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("价格最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type BsServiceEditModel struct{}

// BsServiceDeleteInp 删除服务项目
type BsServiceDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsServiceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type BsServiceDeleteModel struct{}

// BsServiceViewInp 获取指定服务项目信息
type BsServiceViewInp struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *BsServiceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type BsServiceViewModel struct {
	entity.BsService
}

// BsServiceListInp 获取服务项目列表
type BsServiceListInp struct {
	form.PageReq
	Id   int64  `json:"id"  dc:"主键"`
	Pid  int64  `json:"pid" dc:"父键"`
	Name string `json:"name" dc:"服务名称"`
}

func (in *BsServiceListInp) Filter(ctx context.Context) (err error) {
	return
}

type BsServiceListModel struct {
	Id       int64   `json:"id"       dc:"主键"`
	Pid      int64   `json:"pid"      dc:"父键"`
	Name     string  `json:"name"     dc:"服务名称"`
	Price    float64 `json:"price"    dc:"价格"`
	OrderNum int     `json:"orderNum" dc:"序号"`
}

// BsServiceTreeOption 关系树选项
type BsServiceTreeOption struct {
	Id       int64       `json:"id"   dc:"主键"`
	Pid      int64       `json:"pid"  dc:"父键"`
	Name     string      `json:"name" dc:"服务名称"`
	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *BsServiceTreeOption) ID() int64 {
	return t.Id
}

// PID 获取父级节点ID
func (t *BsServiceTreeOption) PID() int64 {
	return t.Pid
}

// SetChildren 设置子节点数据
func (t *BsServiceTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}
