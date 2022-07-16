<template>
  <div class="content">
    <img src="../assets/404.png"/>
    <div class="prompt">哎呀，页面找不到啦！</div>
    <el-button type="primary" @click="refresh">刷新试试</el-button>
  </div>
</template>

<script>
export default {
  name: "Error",
  methods: {
    refresh() {
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
.content{
  width: 100%;
  height: 100vh;
  display: block;
  text-align: center;
}
.content img{
  width: 500px; 
  height: 350px;
  margin: 80px 30px 0;
}
.prompt{
  font-size: 15px;
  color: #606266;
  margin: 10px 0 30px;
}
</style>