package common

import (
	"imall/global"
	"imall/models"
)

// RestPage 分页查询
// page  设置起始页、每页条数,
// name  查询目标表的名称
// query 查询条件,
// dest  查询结果绑定的结构体,
// bind  绑定表结构对应的结构体
func RestPage(page models.Page, name string, query interface{}, dest interface{}, bind interface{}) int64 {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		global.Db.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	}
	return global.Db.Table(name).Where(query).Find(bind).RowsAffected
}
