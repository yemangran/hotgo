// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BsService is the golang structure of table hg_bs_service for DAO operations like Where/Data.
type BsService struct {
	g.Meta   `orm:"table:hg_bs_service, do:true"`
	Id       interface{} // 主键
	Pid      interface{} // 父键
	Level    interface{} // 关系树级别
	Tree     interface{} // 关系树
	Name     interface{} // 服务名称
	Price    interface{} // 价格
	OrderNum interface{} // 序号
}
