<template>
  <div class="container">
    <div class="content"><slot/></div>
  </div>
</template>

<script>
export default {
  name: "Container"
}
</script>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  border-radius: 5px;
  background-color: #FFFFFF;
}
.content{
  padding: 15px;
}
</style>