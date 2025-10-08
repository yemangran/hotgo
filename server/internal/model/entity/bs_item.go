// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BsItem is the golang structure for table bs_item.
type BsItem struct {
	Id       uint64  `json:"id"       orm:"id"       description:"主键"`
	Name     string  `json:"name"     orm:"name"     description:"物品名称"`
	Category int     `json:"category" orm:"category" description:"物品类别"`
	Spec     string  `json:"spec"     orm:"spec"     description:"规格型号"`
	Price    float64 `json:"price"    orm:"price"    description:"单价"`
}
