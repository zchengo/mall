import { areaList } from '@vant/area-data';
import http from '../../../utils/http'

Page({

  // 页面的初始数据
  data: {
    areaList,
    id: '',
    name: '',
    mobile: '',
    province: '',
    city: '',
    district: '',
    detailedAddress: '',
    isDefault: 2,
    area: '',
    show: false,
    checked: false,
  },

  // 展示地区选择弹出框
  showPopup(){
    this.setData({ show: true })
  },

  // 确实选择
  confirmSelect(event){
    this.setData({ show: false });
    this.setData({ province: event.detail.values[0].name });
    this.setData({ city: event.detail.values[1].name });
    this.setData({ district: event.detail.values[2].name });
    this.setData({ area: event.detail.values[0].name + " " + event.detail.values[1].name + " " + event.detail.values[2].name});
  },

  // 取消选择
  cancelSelect({detail}){
    this.setData({ show: false })
  },

  // 改变默认地址
  changeSwitch({detail}){
    this.setData({ checked: detail })
    if (detail) {
      this.setData({ isDefault: 1 })
    }
  },

  // 生命周期函数--监听页面加载
  async onLoad(options){
    if (options.id > 0){
      let res = await http.GET('/address/info',{ id: options.id });
      if(res.data.code === 200){
        var response = res.data.data;
        this.setData({
          id: response.id,
          name: response.name,
          mobile: response.mobile,
          province: response.province,
          city: response.city,
          district: response.district,
          detailedAddress: response.detailedAddress,
          area: response.province + ' ' + response.city + ' ' + response.district,
          isDefault: response.isDefault
        })
        if (this.data.isDefault === 1){
          this.setData({ checked: true })
        }
      }
    }
  },

  // 保存地址
  async submitForm(options) {
    let res = await http.POST('/address/save',{
      id: this.data.id,
      name: this.data.name,
      mobile: this.data.mobile,
      postalCode: this.data.postalCode,
      province: this.data.province,
      city: this.data.city,
      district: this.data.district,
      detailedAddress: this.data.detailedAddress,
      isDefault: this.data.isDefault,
      openId: wx.getStorageSync('openId')
    })
    if(res.data.code === 200){ 
      wx.navigateBack({ url: '/pages/address/list/index' })
    }
  }
})