import http from '../../../utils/http'
import Dialog from '@vant/weapp/dialog/dialog'

Page({

  // 页面的初始数据
  data: {
    addressList: [],
    showEmpty: false,
  },

  // 跳转到添加地址
  toAddForm: function(event) {
    wx.navigateTo({ url: '/pages/address/form/index' })
  },

  // 跳转到编辑地址
  toEditForm: function(event) {
    wx.navigateTo({ url: '/pages/address/form/index?id=' + event.currentTarget.id })
  },

  // 生命周期函数--监听页面显示
  async onShow() {
    let res = await http.GET('/address/list',{ openId: wx.getStorageSync('openId') })
    this.setData({ addressList: res.data.data })
    if (this.data.addressList.length === 0){
      this.setData({ showEmpty: true })
    }
  },

  // 点击删除按钮，出现提示
  onClickDelete: function(event) {
    let _this = this;
    Dialog.confirm({
      message: '确定删除吗？',
    }).then(() => {
      _this.delete(event)
    })
  },
  
  // 删除收货地址
  async delete(event){
    await http.DELETE('/address/delete',{
      id: parseInt(event.currentTarget.id)
    })
    this.onShow()
  }
})