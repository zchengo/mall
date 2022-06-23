<template>
  <div class="container">
    <div class="head">
      <div class="head_data"><div class="number">{{pendPay}}</div><div class="title">待付款</div></div>
      <div class="head_data"><div class="number">{{canceled}}</div><div class="title">已取消</div></div>
      <div class="head_data"><div class="number">{{pendPayed}}</div><div class="title">已付款</div></div>
      <div class="head_data"><div class="number">{{inDelivery}}</div><div class="title">配送中</div></div>
      <div class="head_data"><div class="number">{{finished}}</div><div class="title">已完成</div></div>
    </div>
    <div class="main">
      <BoxCard width="100%" title="本周数据总览" :subTitle="nowTime">
        <div id="week" style="width: 100%;height: 450px;"></div>
      </BoxCard>
    </div>
  </div>
</template>

<script>
import BoxCard from "@/components/BoxCard";
import * as echarts from 'echarts';

export default {
  name: "Index",
  components: {BoxCard},
  data() {
    return {
      pendPay: '--',
      canceled: '--',
      pendPayed: '--',
      inDelivery: '--',
      finished: '--',
      nowTime: '更新于 ' + this.getNowTime()
    }
  },
  mounted() {
    this.getOrderData()
    // 本周数据总览
    let weekChart = echarts.init(document.getElementById('week'), null);
    let option;
    this.$axios.get('/week/data').then((response) => {
      option = {
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985'
            }
          }
        },
        legend: {
          data: ['订单量', '交易金额']
        },
        toolbox: {
          feature: {
            saveAsImage: {}
          }
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true
        },
        xAxis: [
          {
            type: 'category',
            boundaryGap: false,
            data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
          }
        ],
        yAxis: [
          {
            type: 'value'
          }
        ],
        series: [
          {
            name: '订单量',
            type: 'line',
            stack: 'Total',
            areaStyle: {},
            emphasis: {
              focus: 'series'
            },
            data: response.data.data.orders
          },
          {
            name: '交易金额',
            type: 'line',
            stack: 'Total',
            areaStyle: {},
            emphasis: {
              focus: 'series'
            },
            data: response.data.data.amount
          }
        ]
      };
      option && weekChart.setOption(option);
    })
  },
  methods: {
    // 获取订单数据
    getOrderData() {
      this.$axios.get('/today/data').then((response) => {
        this.pendPay = response.data.data.pendPay
        this.canceled = response.data.data.canceled
        this.pendPayed = response.data.data.pendPayed
        this.inDelivery = response.data.data.inDelivery
        this.finished = response.data.data.finished
      })
    },

    // 获取当前时间
    getNowTime() {
      let date = new Date();
      let year = date.getFullYear();
      let month = date.getMonth() + 1;
      let day = date.getDate();
      let hour = date.getHours();
      let minute = date.getMinutes();
      let second = date.getSeconds();
      return year + '/' + this.addZero(month) + '/' + this.addZero(day) + ' ' +
          this.addZero(hour) + ':' + this.addZero(minute) + ':' + this.addZero(second);
    },

    addZero(s) {
      return s < 10 ? ('0' + s) : s;
    }
  }
}
</script>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  border-radius: 5px;
}

.head {
  width: 100%;
  height: 100px;
  background-color: #3978F4;
  border-radius: 12px;
  display: inline-flex;
}
.head_data{
  width: 20%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 12px;
}
.head_data:hover{
  cursor: pointer;
}
.head_data .number{
  font-size: 30px;
  color: #f5f7fc;
  font-weight: 450;
  font-family: Helvetica Neue,Helvetica,Tahoma,Arial,serif;
}
.head_data .title{
  color: #dde4f1;
  padding: 5px 0;
}

.main {
  width: 100%;
  display: inline-flex;
  justify-content: space-around;
}
</style>