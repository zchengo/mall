<template>
  <el-header class="header">
    <div class="navbar">
      <el-menu :default-active="navigation[0].index" mode="horizontal" :ellipsis="false"
               v-for="item in navigation" :key="item.index" router>
        <el-menu-item :index="item.index">{{item.name}}</el-menu-item>
      </el-menu>
    </div>
    <div class="avatar">
      <el-dropdown trigger="click">
        <span class="el-dropdown-link">
          <el-avatar :size="36" :src="avatarURL" style="cursor: pointer;"/>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item :icon="Bell">系统通知</el-dropdown-item>
            <el-dropdown-item :icon="Edit" @click="clickFeedback">问题反馈</el-dropdown-item>
            <el-dropdown-item :icon="SwitchButton" @click="logout">退出账户</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <!-- 问题反馈，对话框 -->
  <el-dialog title="请输入您要反馈的内容" v-model="feedbackDialogVisible" top="30vh" width="45%"
             @close="feedbackDialogVisible = false ">
    <el-input v-model="feedback.content" type="textarea" maxlength="100" :autosize="{ minRows: 5}" show-word-limit/>
    <template #footer>
        <span class="dialog-footer">
          <el-button @click="feedbackDialogVisible = false ">取消</el-button>
          <el-button type="primary" @click="confirmSend">发送</el-button>
        </span>
    </template>
  </el-dialog>
  </el-header>
</template>

<script>
import {Edit, ArrowRight, Bell, SwitchButton} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

export default {
  name: "Header",
  setup() {
    return {Edit, Bell, ArrowRight, SwitchButton}
  },
  data() {
    return {
      avatarURL: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
      feedbackDialogVisible: false,
      feedback: {
        content: ''
      }
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
  methods: {
    // 点击反馈问题
    clickFeedback(){
      this.feedbackDialogVisible = true
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
        this.feedbackDialogVisible = false
      })
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
.el-menu{
  height: 30px !important;
  border-bottom: none;
}
.el-menu--horizontal>.el-menu-item {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  border-radius: 8px 8px 0 0;
  margin: 0 5px;
  border-bottom: 2px solid transparent;
  color: var(--el-menu-text-color);
}
.el-menu-item.is-active{
  background-color: #e9f3fd !important;
}
.header {
  width: 100%;
  height: 55px;
  display: flex;
  padding: 0;
  background-color: #FFFFFF;
}

.navbar {
  width: 80%;
  height: 55px;
  display: flex;
  align-items: flex-end;
  margin-left: 8px;
}

.avatar {
  width: 20%;
  height: 55px;
  padding: 0 18px;
  display: inline-flex;
  justify-content: right;
  align-items: center;
}
</style>