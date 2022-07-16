<template>
  <div id="app">
    <transition :duration="{ enter: 800, leave: 100 }" name="el-fade-in-linear" mode="out-in">
      <router-view />
    </transition>
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<style scoped>
#app {
  height: 100%;
  width: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  background-color: #f0f3f6;
}
</style>
