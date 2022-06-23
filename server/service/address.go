package service

import (
	"imall/common"
	"imall/global"
	"imall/models/app"
)

type AppAddressService struct {
}

// 保存收货地址
func (a *AppAddressService) Save(param app.AddressSaveParam) int64 {
	address := app.Address{
		OpenId:          param.OpenId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
	}
	if param.Id > 0 {
		// 更新收货地址
		address.Id = param.Id
		address.Updated = common.NowTime()
		if address.IsDefault == 1 {
			var addressId uint
			row := global.Db.Table("address").Select("id").
				Where("is_default = ? and open_id = ?", 1, address.OpenId).Take(&addressId).RowsAffected
			if row > 0 {
				global.Db.Table("address").Where("id = ?", addressId).Update("is_default", 2)
				return global.Db.Updates(&address).RowsAffected
			} else {
				return global.Db.Updates(&address).RowsAffected
			}
		}
		return global.Db.Updates(&address).RowsAffected
	} else {
		// 添加收货地址
		address.Created = common.NowTime()
		if address.IsDefault == 1 {
			var addressId uint64
			row := global.Db.Table("address").Select("id").
				Where("is_default = ? and open_id = ?", 1, address.OpenId).Take(&addressId).RowsAffected
			if row > 0 {
				global.Db.Table("address").Where("id = ?", addressId).Update("is_default", 2)
				return global.Db.Create(&address).RowsAffected
			} else {
				return global.Db.Create(&address).RowsAffected
			}
		}
		return global.Db.Create(&address).RowsAffected
	}
}

// 删除收货地址
func (a *AppAddressService) Delete(param app.AddressDeleteParam) int64 {
	return global.Db.Delete(&app.Address{}, param.Id).RowsAffected
}

// 获取收货地址列表
func (a *AppAddressService) GetList(param app.AddressListParam) []app.AddressList {
	addressList := make([]app.AddressList, 0)
	global.Db.Table("address").Where("open_id = ?", param.OpenId).Order("is_default").Find(&addressList)
	return addressList
}

// 获取收货地址信息
func (a *AppAddressService) GetInfo(param app.AddressInfoParam) app.AddressInfo {
	var addressInfo app.AddressInfo
	global.Db.Table("address").First(&addressInfo, param.Id)
	return addressInfo
}
