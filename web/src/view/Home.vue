<template>
  <el-container class="container" direction="horizontal">
    <el-aside class="aside" width="105px">
      <Menu />
    </el-aside>
    <el-container direction="vertical">
      <Header />
      <el-main class="main">
        <transition name="fade" mode="out-in">
          <router-view></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Menu from "@/components/Menu";
import Header from "@/components/Header";

export default {
  name: "Home",
  components: { Header, Menu },
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
  width: 108px;
  height: 55px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  justify-items: center;
  cursor: pointer;
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
  padding: 12px;
}

.fade-enter-from,.fade-leave-to {
  opacity: 0;
}

.fade-enter-active,.fade-leave-active {
  transition: opacity 1s ease;
}
</style>