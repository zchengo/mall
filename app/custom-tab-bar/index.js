Component({
  // 页面的初始数据
  data: {
    active: 0,
    tabList: [{ 
      name: '首页',
      icon: '/assets/home-o.png',
      activeIcon: '/assets/home.png',
      navTo: "/pages/home/index",
    },{ 
      name: '订单',
      icon: '/assets/order-o.png',
      activeIcon: '/assets/order.png',
      navTo: "/pages/order/index",
    },{ 
      name: '我的',
      icon: '/assets/mine-o.png',
      activeIcon: '/assets/mine.png',
      navTo: "/pages/mine/index",
    }]
  },
  methods: {
    onChange(event) {
      // event.detail 的值为当前选中项的索引
      this.setData({ active: event.detail });
      wx.switchTab({ url: this.data.tabList[event.detail].navTo })
    },
    init() {
      const page = getCurrentPages().pop();
      this.setData({
      　  active: this.data.tabList.findIndex(item => item.navTo === `/${page.route}`)
      });
    }
  }
})