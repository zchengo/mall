Page({

  // 页面的初始数据
  data: {
    email: '',
    content: ''
  },

  // 提交反馈
  async submitForm(options) {
    let res = await request.POST('/feedback/submit',{
      email: this.data.email,
      content: this.data.content
    })
  }
})