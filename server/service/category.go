package service

import (
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
)

type WebCategoryService struct {
}

type AppCategoryService struct {
}

// 创建商品类目
func (c *WebCategoryService) Create(param web.CategoryCreateParam) int64 {
	category := web.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	return global.Db.Create(&category).RowsAffected
}

// 删除商品类目
func (c *WebCategoryService) Delete(param web.CategoryDeleteParam) int64 {
	rows := global.Db.Where("parent_id = ?", param.Id).Delete(&web.Category{}).RowsAffected
	if rows < 0 {
		return 0
	}
	return global.Db.Delete(&web.Category{}, param.Id).RowsAffected
}

// 更新商品类目
func (c *WebCategoryService) Update(param web.CategoryUpdateParam) int64 {
	category := web.Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(&category).RowsAffected
}

// 获取商品类目列表
func (g *WebCategoryService) GetList(param web.CategoryQueryParam) []web.CategoryList {
	query := &web.Category{
		Name:     param.Name,
		ParentId: param.ParentId,
	}
	categoryList := make([]web.CategoryList, 0)
	global.Db.Table("category").Where(query).Find(&categoryList)
	return categoryList
}

// 获取商品类目选项
func (c *WebCategoryService) GetOption() (option []web.CategoryOption) {
	selectList := make([]web.CategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

// 获取商品类目选项
func (c *AppCategoryService) GetOption() (option []app.CategoryOption) {
	categorys := make([]app.Category, 0)
	categoryOptions := make([]app.CategoryOption, 0)
	global.Db.Table("category").Where("level = 1").Find(&categorys)
	for _, item := range categorys {
		option := app.CategoryOption{
			Id:   item.Id,
			Text: item.Name,
		}
		categoryOptions = append(categoryOptions, option)
	}
	return categoryOptions
}

// 获取树形结构的选项
func getTreeOptions(id uint64, cateList []web.CategoryList) (option []web.CategoryOption) {
	optionList := make([]web.CategoryOption, 0)
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2) {
			option := web.CategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 2 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}
