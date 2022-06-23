<template>
  <div class="container">
    <div class="label">{{label}}</div>
    <div class="content"><slot/></div>
  </div>
</template>

<script>
export default {
  name: "Descriptions",
  props: ['label']
}
</script>

<style scoped>
.container{
  width: 90%;
  margin: 5px 5%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  padding: 5px;
}
.label{
  width: 15%;
  padding-right: 5px;
  display: flex;
  font-weight: 450;
  justify-content: flex-end;
}
.content{
  width: 85%;
  display: flex;
  align-items: center;
  padding: 0 10px;
}
</style>