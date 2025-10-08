// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BsItem is the golang structure of table hg_bs_item for DAO operations like Where/Data.
type BsItem struct {
	g.Meta   `orm:"table:hg_bs_item, do:true"`
	Id       interface{} // 主键
	Name     interface{} // 物品名称
	Category interface{} // 物品类别
	Spec     interface{} // 规格型号
	Price    interface{} // 单价
}
