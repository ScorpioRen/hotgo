// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table hg_sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta    `orm:"table:hg_sys_dict_data, dto:true"`
	Id        interface{} // 字典编码
	Label     interface{} // 字典标签
	Value     interface{} // 字典键值
	Type      interface{} // 字典类型
	ListClass interface{} // 表格回显样式
	IsDefault interface{} // 是否默认
	Sort      interface{} // 字典排序
	Remark    interface{} // 备注
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}