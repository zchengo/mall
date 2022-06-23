<template>
  <div class="box" :style="boxStyle">
    <div class="title">
      <div>{{ title }}</div>
      <div style="font-size: 15px;color: #3d4c5d;padding: 0 12px;font-weight: 430;">{{ subTitle }}</div>
    </div>
    <div class="content"><slot/></div>
  </div>
</template>

<script>
export default {
  name: "BoxCard",
  props: ['width', 'height', 'title', 'subTitle'],
  setup(props) {
    return {
      boxStyle: {
        width: props.width,
        maxHeight: props.height
      }
    }
  }
}
</script>

<style scoped>
.box {
  background-color: white;
  border-radius: 5px;
  padding: 12px;
  margin: 12px 0;
}

.title {
  width: 100%;
  height: 22px;
  display: inline-flex;
  justify-content: space-between;
  align-items: center;
  font-size: 16px;
  font-weight: 450;
  padding: 0 8px;
  border-left: 5px solid #1890ff;
}

.content {
  width: 100%;
  margin: 10px 0;
  height: 92%;
}
</style>