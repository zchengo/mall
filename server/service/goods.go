package service

import (
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"strings"
)

type WebGoodsService struct {
}

type AppGoodsService struct {
}

// 创建商品
func (g *WebGoodsService) Create(param web.GoodsCreateParam) int64 {
	goods := web.Goods{
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Name:       param.Name,
		Price:      param.Price,
		Quantity:   param.Quantity,
		ImageUrl:   param.ImageUrl,
		Remark:     param.Remark,
		Status:     1,
		Created:    common.NowTime(),
	}
	return global.Db.Create(&goods).RowsAffected
}

// 删除商品
func (g *WebGoodsService) Delete(param web.GoodsDeleteParam) int64 {
	return global.Db.Delete(&web.Goods{}, param.Id).RowsAffected
}

// 更新商品
func (g *WebGoodsService) Update(param web.GoodsUpdateParam) int64 {
	goods := web.Goods{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Name:       param.Name,
		Price:      param.Price,
		Quantity:   param.Quantity,
		ImageUrl:   param.ImageUrl,
		Remark:     param.Remark,
		Updated:    common.NowTime(),
	}
	return global.Db.Model(&goods).Updates(goods).RowsAffected
}

// 更新商品状态
func (g *WebGoodsService) UpdateStatus(param web.GoodsStatusUpdateParam) int64 {
	goods := web.Goods{
		Id:     param.Id,
		Status: param.Status,
	}
	return global.Db.Model(&goods).Update("status", goods.Status).RowsAffected
}

// 获取商品列表
func (g *WebGoodsService) GetList(param web.GoodsListParam) ([]web.GoodsList, int64) {
	query := &web.Goods{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		Status:     param.Status,
	}
	goodsList := make([]web.GoodsList, 0)
	rows := common.RestPage(param.Page, "goods", query, &goodsList, &[]web.Goods{})
	return goodsList, rows
}

// 获取商品列表
func (g *AppGoodsService) GetList(param app.GoodsQueryParam) []app.GoodsList {
	// 查询分类列表
	categoryList := make([]app.Category, 0)
	cids := make([]uint64, 0)
    global.Db.Table("category").Where("parent_id = ?", param.CategoryId).Find(&categoryList)
	for _, c := range categoryList {
		cids = append(cids, c.Id)
	}

	// 根据分类编号查询商品
	goodsList := make([]app.GoodsList, 0)
	global.Db.Table("goods").Where("category_id in ? and status = 1", cids).Find(&goodsList)
	return goodsList
}

// 获取商品详情
func (g *AppGoodsService) GetInfo(param app.GoodsInfoParam) app.GoodsInfo {
	var goods app.Goods
	global.Db.Table("goods").First(&goods, param.Id)
	goodsInfo := app.GoodsInfo{
		Id:       goods.Id,
		Title:    goods.Title,
		Name:     goods.Name,
		Price:    goods.Price,
		ImageUrl: goods.ImageUrl,
		Remark:   goods.Remark,
	}
	return goodsInfo
}

// 搜索商品
func (g *AppGoodsService) Search(param app.GoodsSearchParam) []app.GoodsList {
	goodsList := make([]app.GoodsList, 0)
	keyWord := strings.Join([]string{"%", "%"}, param.KeyWord)
	global.Db.Table("goods").Where("title like ? and status = 1", keyWord).Find(&goodsList)
	return goodsList
}
