// pages/order/confirm/index.js
import http from '../../../utils/http'

Page({

  // 页面的初始数据
  data: {
    goodsItem: [],
    totalPrice: 0.00,
    totalGoodsCount: 0,
    show: false,
    showView: true
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getCartInfo()
  },

  // 获取购物车信息
  async getCartInfo(){
    let res = await http.GET('/cart/info', {openId: wx.getStorageSync('openId')})
    this.setData({
      goodsItem: res.data.data.cartItem,
      totalPrice: res.data.data.totalPrice
    })
    let totalGoodsCount = 0
    for (let i = 0; i < this.data.goodsItem.length; i++) {
      totalGoodsCount = totalGoodsCount + this.data.goodsItem[i].count
      if (this.data.goodsItem[i].id == this.data.goodsId) {
        this.setData({goodsCount: this.data.goodsItem[i].count})
        console.log(this.data.goodsItem[i].count);
      }
    }
    this.setData({totalGoodsCount: totalGoodsCount})
  },

  // 提交订单
  submitOrder(){
    this.setData({show: true})
  },

  // 确认支付
  confirmPay(){
    this.setData({showView: false})
    http.POST('/order/submit', {openId: wx.getStorageSync('openId')})
  },

  // 返回首页
  backToHome(){
    wx.switchTab({url: '/pages/home/index'})
  }
})