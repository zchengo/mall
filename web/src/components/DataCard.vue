<template>
  <div class="main">
    <div class="icon"><slot/></div>
    <div class="data">
      <div style="font-weight: 400;font-size: 15px;">{{title}}</div>
      <div style="font-size: 22px;font-weight: 460;">{{number}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DataCard",
  props: ['title','number'],
}
</script>

<style scoped>
.main{
  width: 25%;
  height: 80px;
  display: inline-flex;
  background-color: white;
  align-items: center;
  margin: 5px;
  border-radius: 5px;
}
.icon{
  width: 35%;
  text-align: center;
}
.data{
  width: 65%;
}
</style>