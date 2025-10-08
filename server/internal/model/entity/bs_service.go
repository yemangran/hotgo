// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BsService is the golang structure for table bs_service.
type BsService struct {
	Id       uint64  `json:"id"       orm:"id"        description:"主键"`
	Pid      int64   `json:"pid"      orm:"pid"       description:"父键"`
	Level    int64   `json:"level"    orm:"level"     description:"关系树级别"`
	Tree     string  `json:"tree"     orm:"tree"      description:"关系树"`
	Name     string  `json:"name"     orm:"name"      description:"服务名称"`
	Price    float64 `json:"price"    orm:"price"     description:"价格"`
	OrderNum int     `json:"orderNum" orm:"order_num" description:"序号"`
}
