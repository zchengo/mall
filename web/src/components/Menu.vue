<template>
  <el-menu default-active="/index"
           background-color="#1b273f"
           :unique-opened="true" text-color="#DCDFE6" router>
    <el-menu-item index="/index" @click="changeNav(this.index)">
      <el-icon><home-filled /></el-icon><span>首页</span>
    </el-menu-item>
    <el-menu-item index="/goods" @click="changeNav(this.goods)">
      <el-icon><goods-filled /></el-icon><span>商品</span>
    </el-menu-item>
    <el-menu-item index="/order" @click="changeNav(this.order)">
      <el-icon><list /></el-icon><span>订单</span>
    </el-menu-item>
    <el-menu-item index="/market" @click="changeNav(this.market)">
      <el-icon><ticket /></el-icon><span>营销</span>
    </el-menu-item>
  </el-menu>
</template>

<script>
import {
  HomeFilled,
  GoodsFilled,
  List,
  Ticket
} from "@element-plus/icons-vue";

export default {
  name: "Menu",
  components: {
    HomeFilled,
    GoodsFilled,
    List,
    Ticket
  },
  data() {
    return {
      index: [{
        name: '我的主页',
        index: '/index',
      }],
      goods: [{
        name: '商品管理',
        index: '/goods'
      }, {
        name: '类目管理',
        index: '/category'
      }],
      order: [{
        name: '订单管理',
        index: '/order',
      }],
      market: [{
        name: '活动管理',
        index: '/market',
      }]
    }
  },
  methods: {
    changeNav(menu) {
      this.$store.commit("addNav", menu)
    }
  }
}
</script>

<style scoped>
.el-menu {
  width: 165px !important;
  height: 40px !important;
  border-right: 0 !important;
}
.menu-item{
  background-color: #3d4c5d;
}
.el-menu-item{
  min-width: 105px !important;
  max-height: 50px !important;
}
.el-menu-item.is-active{
  /*background-color: #409EFF !important;*/
}
</style>