<script setup lang="ts">
import router from "../router/index.ts"
</script>

<template>
  <!--    三部分 左 点击侧栏, 中 logo, 右 语言-->
  <div class="navbar">
    <div class="navbar-left">
    </div>
    <div class="navbar-center">
    </div>
    <div class="navbar-right">
      <div v-for="route in router.getRoutes()">
        <router-link :to="route.path" :key="route.name">
          {{ route.name }}
        </router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.navbar {
  z-index: 3;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  height: 80px;
  background-color: #d0e9ff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.navbar-left img {
  height: 40px;
}

.navbar-left,
.navbar-right {
  display: flex;
  align-items: center;
}

.navbar-center img {
  height: 80px;
}

</style>