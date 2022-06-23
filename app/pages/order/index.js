// pages/order/index.js
import http from '../../utils/http'
import Dialog from '@vant/weapp/dialog/dialog';

Page({

  /**
   * 页面的初始数据
   */
  data: {
    tabList: ['未完成','已完成','',''],
    active: 0,
    orderList: []
  },

  tabsChange(event) {
    this.getOrderList(event.detail.name)
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getTabBar().init();
    this.getOrderList(0)
  },

  // 获取订单列表
  async getOrderList(orderType){
    let res = await http.GET('/order/list',{
      type: orderType+1,
      openId: wx.getStorageSync('openId')
    });
    this.setData({orderList: res.data.data})
  },

  // 取消订单
  async cancelOrder(event){
    let res = await http.PUT('/order/update',{
      id: event.currentTarget.id,
      status: 2
    });
  },

  // 确认取消订单
  confirmCancelOrder(event){
    let _this = this
    Dialog.confirm({
      title: '确认取消订单吗'
    }).then(() => {
      _this.cancelOrder(event)
      _this.getOrderList()
    });
  }
})