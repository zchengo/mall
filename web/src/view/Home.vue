<template>
  <el-container class="container" direction="horizontal">
    <el-aside class="aside" width="110px">
      <div class="logo">
        <img src="../assets/logo.png" style=""/>
        <div class="logo_name">iM</div>
      </div>
      <el-menu :default-active="defaultActive"
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
    </el-aside>
    <el-container direction="vertical">
      <el-header class="header">
      <div class="navbar">
        <el-icon size="20px" style="padding: 5px;cursor: pointer;" @click="refresh"><RefreshRight /></el-icon>
      </div>
      <div class="avatar">
        <div class="operate">
          <el-icon size="20px" @click="onFullscreen" style="cursor: pointer;padding: 10px;"><FullScreen /></el-icon>
          <el-icon size="20px" style="cursor: pointer;padding: 10px;"><Bell /></el-icon>
          <el-icon size="20px" @click="feedbackPopoverVisible = true" style="cursor: pointer;padding: 10px;"><Service /></el-icon>
        </div>
        <el-dropdown trigger="click">
          <span class="el-dropdown-link">
            <el-avatar :size="30" :src="avatarURL" style="cursor: pointer;border: 1px solid #d9ecff;"/>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <div style="text-align: center;">
                <el-avatar :size="45" :src="avatarURL" style="cursor: pointer;border: 1px solid #d9ecff;"/>
                <div>SID: {{sid}}</div>
              </div>
              <el-dropdown-item :icon="SwitchButton" @click="logout" divided>退出账户</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>
    <el-menu :default-active="defaultNav" 
             class="el-menu-demo"
             v-if="navigation.length > 0"
             mode="horizontal" router>
      <el-menu-item v-for="item in navigation" :key="item.index" :index="item.index">{{item.name}}</el-menu-item>
    </el-menu> 
      <el-main class="main">
        <transition :duration="{ enter: 800, leave: 100 }" name="el-fade-in-linear" mode="out-in">
          <router-view></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import {RefreshRight, Edit, ArrowRight, SwitchButton, Service, FullScreen, Bell} from "@element-plus/icons-vue";
import {
  HomeFilled,
  GoodsFilled,
  List,
  Ticket
} from "@element-plus/icons-vue";

export default {
  name: "Home",
  components: {RefreshRight, Service, FullScreen, Bell, HomeFilled,
    GoodsFilled,
    List,
    Ticket },
  setup() {
    return {Edit, Bell, ArrowRight, SwitchButton}
  },
  data() {
    return {
      sid: '',
      avatarURL: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
      feedback: {
        content: ''
      },
      fullscreenStatus: false,
      feedbackPopoverVisible: false,
      defaultNav: '/index',

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
  computed: {
    navigation: {
      get() {
        return this.$store.state.navigation
      },
      set(val) {
        this.$store.state.navigation = val
      }
    },
    defaultActive: {
      get() {
        return this.$store.state.defaultActive
      },
      set(val) {
        this.$store.state.defaultActive = val
      }
    }
  },
  updated() {
    this.defaultNav = this.$store.state.navigation[0].index
  },
  mounted() {
    this.sid = localStorage.getItem('sid')
    this.defaultActive = this.$store.state.defaultActive
  },
  methods: {

    // 刷新
    refresh(){
      window.location.reload()
    },

    // 
    getNavigation(){
      return this.$store.getter
    },

    // 全屏或退出全屏 
    onFullscreen() {
      if(this.fullscreenStatus){
        document.webkitExitFullscreen()
        this.fullscreenStatus = false
      } else {
        document.documentElement.webkitRequestFullscreen()
        this.fullscreenStatus = true
      }
    },

    // 退出登录
    logout() {
      localStorage.clear();
      sessionStorage.clear();
      this.$router.push('/');
    },

    changeNav(menu) {
      this.$store.commit("addNav", menu)
    }
  }
};
</script>

<style scoped>
.container {
  height: 100vh;
  background-color: #e8ecf1;
}

.aside {
  background-color: #1b273f;
}

.logo {
  width: 110px;
  height: 55px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  justify-items: center;
  cursor: pointer;
}
.logo img{
  width: 24px;
  height: 24px;
  border-radius: 5px;
  box-shadow: 0px 0px 3px #ebeff4;
}
.logo_name{
  color: #ebeff4;
  font-size: 25px;
  font-weight: bolder;
  padding-left: 5px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.name {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  color: #409eff;
  font-size: 18px;
  padding-top: 2px;
  margin-left: 3px;
  font-weight: 600;
  letter-spacing: 2px;
  font-family: PingFang SC, sans-serif;
}

.image {
  width: 45px;
  height: 45px;
}

.main {
  padding: 10px;
  overflow-y: hidden;
}

.header {
  width: 100%;
  height: 55px;
  display: flex;
  padding: 0;
  border-bottom: 1px solid #f0f2f4;
  background-color: #FFFFFF;
}

.navbar {
  width: 80%;
  height: 55px;
  display: flex;
  align-items: center;
  margin-left: 8px;
}
.operate{
  height: 55px;
  margin: 0 20px;
  display: inline-flex;
  justify-items: center;
  align-items: center;
}
.avatar {
  width: 20%;
  height: 55px;
  padding: 0 18px;
  display: inline-flex;
  justify-content: right;
  align-items: center;
}
.el-menu{
  height: 35px !important;
  padding: 5px !important;
  border-bottom: none !important;
  background-color: none !important;
  border: none !important;
}
.el-menu-item{
  max-height: 50px !important;
  border: none !important;
  border-radius: 10px;
}
.el-menu--horizontal>.el-menu{
  height: 35px !important;
  padding: 0 5px;
  border-bottom: none !important;
  background-color: none !important;
  border: none !important;
}
.el-menu--horizontal>.el-menu-item {
  border: none;
  border-radius: 20px;
  margin: 0 3px;
}
.el-menu--horizontal>.el-menu-item.is-active{
  background-color: #e9f3fd !important;
}
</style>