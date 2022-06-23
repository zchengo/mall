// 通用请求路径basePath
var basePath = 'http://localhost:8000/app'

// 发送GET请求
function GET(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/json' },
      method: 'GET',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送POST请求
function POST(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/x-www-form-urlencoded' },
      method: 'POST',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送PUT请求
function PUT(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/x-www-form-urlencoded' },
      method: 'PUT',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送DELETE请求
function DELETE(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/json' },
      method: 'DELETE',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

module.exports = {
  GET, POST, PUT, DELETE
}