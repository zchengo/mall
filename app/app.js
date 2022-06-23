import http from './utils/http'

App({
  onLaunch() {
    // 自动登录
    wx.login({
      async success (res) {
        if (res.code) {
          let response = await http.POST('/login',{ 
            code: res.code 
          })
          wx.setStorageSync('openId', response.data.data.openId);
        }
      }
    })
  },
  globalData: {
    userInfo: null
  }
})
