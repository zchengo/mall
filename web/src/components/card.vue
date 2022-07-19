<template>
  <el-card shadow="hover" body-style="cursor: pointer;">
    <div class="title">{{ title }}</div>
    <div v-if="number > 0" class="number">{{ number }}</div>
    <div v-else class="number">-</div>
  </el-card>
</template>

<script>
export default {
  props: ["title", "number"],
};
</script>

<style scoped>
.title {
  color: #666;
  font-size: 15px;
  line-height: 1;
}

.number {
  font-size: 25px;
  color: #111;
  letter-spacing: 0;
  line-height: 1;
  margin-top: 10px;
  font-family: PingFang SC, Microsoft YaHei;
}
</style>