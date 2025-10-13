// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BsWorkOrder is the golang structure for table bs_work_order.
type BsWorkOrder struct {
	Id                 uint64      `json:"id"                 orm:"id"                  description:"主键"`
	UserId             int         `json:"userId"             orm:"user_id"             description:"用户Id"`
	CreateTime         *gtime.Time `json:"createTime"         orm:"create_time"         description:"创建时间"`
	CustomerName       string      `json:"customerName"       orm:"customer_name"       description:"客户名称"`
	CustomerAddress    string      `json:"customerAddress"    orm:"customer_address"    description:"客户地址"`
	CustomerContact    string      `json:"customerContact"    orm:"customer_contact"    description:"客户联系方式"`
	CustomerPerson     string      `json:"customerPerson"     orm:"customer_person"     description:"客户对接人"`
	ProblemDescription string      `json:"problemDescription" orm:"problem_description" description:"问题描述"`
	DefectDescription  string      `json:"defectDescription"  orm:"defect_description"  description:"检测描述"`
	DispatchId         int         `json:"dispatchId"         orm:"dispatch_id"         description:"分配处理人Id"`
	SendType           int         `json:"sendType"           orm:"send_type"           description:"发起类型"`
	AcceptType         int         `json:"acceptType"         orm:"accept_type"         description:"收取方式"`
	ProductType        int         `json:"productType"        orm:"product_type"        description:"产品类型"`
	Status             int         `json:"status"             orm:"status"              description:"工单状态"`
	SuggestSolution    []int       `json:"suggestSolution"    orm:"suggest_solution"    description:"建议处理方案(多选)"`
	ActualSolution     []int       `json:"actualSolution"     orm:"actual_solution"     description:"实际处理方案(多选)"`
	TotalMoney         float64     `json:"totalMoney"         orm:"total_money"         description:"总金额"`
	InvoiceType        int         `json:"invoiceType"        orm:"invoice_type"        description:"开票类型"`
	Remark             string      `json:"remark"             orm:"remark"              description:"备注"`
}
