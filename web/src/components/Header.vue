<template>
  <div>
    <el-header class="header">
      <div class="navbar">
        <el-icon size="20px" style="padding: 5px;cursor: pointer;"><Fold /></el-icon>
      </div>
      <div class="avatar">
        <div class="operate">
          <el-icon size="20px" @click="onFullscreen" style="cursor: pointer;padding: 10px;"><FullScreen /></el-icon>
          <el-icon size="20px" style="cursor: pointer;padding: 10px;"><Bell /></el-icon>
          <el-popover title="向客服反馈问题" placement="bottom" :width="280" :visible="feedbackPopoverVisible" trigger="click">
            <template #reference>
              <el-icon size="20px" @click="feedbackPopoverVisible = true" style="cursor: pointer;padding: 10px;"><Service /></el-icon>
            </template>
            <el-input v-model="feedback.content" type="textarea" maxlength="100" :autosize="{ minRows: 3}" show-word-limit/>
            <div style="padding-top: 10px; float: right;">
              <el-button size="small" @click="cancelFeedback">取消</el-button>
              <el-button size="small" type="primary" @click="confirmSend">提交</el-button>
            </div>
          </el-popover>
        </div>
        <el-dropdown trigger="click">
          <span class="el-dropdown-link">
            <el-avatar :size="30" :src="avatarURL" style="cursor: pointer;border: 1px solid #d9ecff;"/>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <div style="text-align: center;">
                <el-avatar :size="45" :src="avatarURL" style="cursor: pointer;border: 1px solid #d9ecff;"/>
                <div>ID: 203050</div>
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
  </div>
</template>

<script>
import {Fold, Edit, ArrowRight, SwitchButton, Service, FullScreen, Bell} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

export default {
  name: "Header",
  components: {Fold, Service, FullScreen, Bell},
  setup() {
    return {Edit, Bell, ArrowRight, SwitchButton}
  },
  data() {
    return {
      avatarURL: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
      feedback: {
        content: ''
      },
      fullscreenStatus: false,
      feedbackPopoverVisible: false,
      defaultNav: '/index'
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
    }
  },
  updated() {
    this.defaultNav = this.$store.state.navigation[0].index
  },
  methods: {

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

    // 确认发送反馈内容
    confirmSend() {
      this.$axios.post('/feedback/send', {
        content: this.feedback.content
      }).then((response) => {
        if (response.data.code === 200){
          ElMessage({message: '反馈成功，我们将尽快回复！', type: 'success'})
        }
        console.log(response)
      })
      this.feedback.content = ''
      this.feedbackPopoverVisible = false
    },

    // 取消反馈
    cancelFeedback(){
      this.feedback.content = ''
      this.feedbackPopoverVisible = false
    },

    // 退出登录
    logout() {
      localStorage.clear();
      sessionStorage.clear();
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
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
  height: 40px !important;
  border-top: 2px;
  padding: 5px;
  border-bottom: none !important;
}
.el-menu--horizontal>.el-menu-item {
  border: none;
  border-radius: 20px;
  margin: 0 3px;
}
.el-menu-item.is-active{
  background-color: #e9f3fd !important;
  /* border-bottom: 2px solid transparent; */
}
</style>