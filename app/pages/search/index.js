// pages/search/index.js
import http from '../../utils/http'
Page({

  // 页面的初始数据
  data: {
    keyWord: '',
    goodsList: []
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.searchGoods()
  },

  // 绑定输入框的值
  onChange(e) {
    this.setData({ keyWord: e.detail});
  },

  // 搜索商品
  async searchGoods() {
    let res = await http.GET('/goods/search?keyWord=' + this.data.keyWord)
    this.setData({ goodsList: res.data.data })
  },

  // 点击商品卡片，跳转到商品详情页
  checkGoodsDetails(event) {
    wx.navigateTo({ url: '/pages/goods/index?id=' + event.currentTarget.id })
  }
})