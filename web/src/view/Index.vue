<template>
  <el-row :gutter="12">
    <el-col :span="24">
      <el-row :gutter="12">
        <el-col :span="4">
          <card title="待付款" :number="pendPay"></card>
        </el-col>
        <el-col :span="4">
          <card title="待发货" :number="payed"></card>
        </el-col>
        <el-col :span="4">
          <card title="配送中" :number="inDelivery"></card>
        </el-col>
        <el-col :span="4">
          <card title="已取消" :number="canceled"></card>
        </el-col>
        <el-col :span="4">
          <card title="已完成" :number="finished"></card>
        </el-col>
        <el-col :span="4">
          <card title="支付金额" :number="payAmount"></card>
        </el-col>
      </el-row>
    </el-col>
    <el-col :span="14">
      <div style="padding-top: 10px">
        <el-card shadow="never">
          <div id="orderData" style="width: 100%; height: 450px"></div>
        </el-card>
      </div>
    </el-col>
    <el-col :span="10">
      <div style="padding-top: 10px">
        <el-row>
          <el-col :span="24">
            <el-card shadow="never">
              <div style="margin-bottom: 12px; margin-left: 4px">常用功能</div>
              <el-row :gutter="12">
                <el-col :span="6">
                  <links title="添加商品" bgColor="#79bbff" to="goods"
                    ><Plus
                  /></links>
                </el-col>
                <el-col :span="6">
                  <links title="订单管理" bgColor="#eebe77" to="order"
                    ><Document
                  /></links>
                </el-col>
                <el-col :span="6">
                  <links title="活动管理" bgColor="#95d475" to="market"
                    ><ShoppingCart
                  /></links>
                </el-col>
                <el-col :span="6">
                  <links title="全部功能" bgColor="#f89898"><Menu /></links>
                </el-col>
              </el-row>
            </el-card>
          </el-col>
          <el-col :span="24">
            <div style="padding-top: 10px">
              <el-card shadow="never">
                <div id="shopData" style="width: 100%; height: 258px"></div>
              </el-card>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import card from "@/components/card";
import links from "@/components/links";
import { Plus, Document, ShoppingCart, Menu } from "@element-plus/icons-vue";

import * as echarts from "echarts";
export default {
  name: "Index",
  components: { card, links, Plus, Document, ShoppingCart, Menu },
  data() {
    return {
      pendPay: 0,
      payed: 0,
      inDelivery: 0,
      canceled: 0,
      finished: 0,
      payAmount: 0.0,
    };
  },
  mounted() {
    this.getTodayData();
    // 本周数据总览
    let orderData = echarts.init(document.getElementById("orderData"), null);
    let shopData = echarts.init(document.getElementById("shopData"), null, {
      height: 275,
    });
    let orderDataOption;
    let shopDataOption;
    this.$axios
      .get("/order/data", {
        params: {
          sid: localStorage.getItem("sid"),
        },
      })
      .then((response) => {
        console.log(response);
        orderDataOption = {
          title: {
            text: "订单数据（当日各时段）",
            textStyle: {
              fontSize: "16px",
              fontWeight: "normal",
            },
          },
          xAxis: {
            type: "category",
            data: [
              "06",
              "07",
              "08",
              "09",
              "10",
              "11",
              "12",
              "13",
              "14",
              "15",
              "16",
              "17",
              "18",
              "19",
              "20",
              "21",
              "22",
              "23",
              "24",
            ],
          },
          yAxis: {
            type: "value",
          },
          grid: {
            left: "3%",
            right: "4%",
            bottom: "3%",
            containLabel: true,
          },
          series: [
            { 
              name: "订单数",
              data: response.data.data.orders,
              type: "bar",
            },
          ],
          tooltip: {
            trigger: "axis",
            axisPointer: {
              type: "shadow",
            },
          },
        };
        orderDataOption && orderData.setOption(orderDataOption);
      });

    this.$axios
      .get("/shop/data", {
        params: {
          sid: localStorage.getItem("sid"),
        },
      })
      .then((response) => {
        console.log(response);
        shopDataOption = {
          title: {
            text: "店铺数据",
            textStyle: {
              fontSize: "16px",
              fontWeight: "normal",
            },
          },
          tooltip: {
            trigger: "axis",
            axisPointer: {
              type: "cross",
              label: {
                backgroundColor: "#6a7985",
              },
            },
          },
          grid: {
            left: "3%",
            right: "4%",
            bottom: "10%",
            containLabel: true,
          },
          xAxis: [
            { 
              type: "category",
              boundaryGap: false,
              data: ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
            },
          ],
          yAxis: [
            {
              type: "value",
            },
          ],
          series: [
            {
              name: "支付金额",
              type: "line",
              stack: "Total",
              areaStyle: {},
              emphasis: {
                focus: "series",
              },
              smooth: true,
              data: response.data.data.amount,
            },
          ],
        };
        shopDataOption && shopData.setOption(shopDataOption);
      });
  },
  methods: {
    // 获取订单数据
    getTodayData() {
      this.$axios
        .get("/today/data", {
          params: {
            sid: localStorage.getItem("sid"),
          },
        })
        .then((response) => {
          this.pendPay = response.data.data.pendPay;
          this.payed = response.data.data.payed;
          this.inDelivery = response.data.data.inDelivery;
          this.canceled = response.data.data.canceled;
          this.finished = response.data.data.finished;
          this.payAmount = response.data.data.payAmount;
        });
    },
  },
};
</script>

<style scoped>
</style>