<template>
  <el-card shadow="hover" body-style="cursor: pointer;padding: 15px;" @click="navTo">
    <div class="content">
      <div class="icon" :style="{backgroundColor: bgColor}">
      <el-icon size="30px" color="#FFFFFF"><slot /></el-icon>
      </div>
      <div class="title">{{ title }}</div>
    </div>
  </el-card>
</template>

<script>
export default {
  props: ["bgColor",'title','to'],
  methods: {

    navTo(){
      this.$router.push({
        name: this.to,
        params: {
          dialog: true
        }
      });
      this.$store.commit("setMenu", '/' + this.to)
    }
  },
}
</script>

<style scoped>
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.icon {
  width: 50px;
  height: 50px;
  border-radius: 5px;
  /* background-color: #79bbff; */
  display: flex;
  align-items: center;
  justify-content: center;
}
.title {
  font-size: 13px;
  padding: 3px 0;
}

.el-card {
    border: none !important;
}
</style>