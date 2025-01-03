<template>
  <div class="online-status">
    <div class="dot" :style="{ backgroundColor: props.color }"></div>
    <div class="pulse" v-if="isOnline" :style="{ backgroundColor: props.spreadColor }"></div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
  color: string;
  spreadColor: string;
  isOnline: boolean;
}>();
</script>

<style scoped>
.online-status {
  position: relative;
  width: 15px;
  height: 15px;
}

.dot {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  z-index: 1; /* Ensure the base color layer is on top */
}

.pulse {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  animation: pulse 1.5s infinite;
  z-index: 0; /* Ensure the radar layer is below the base color layer */
}

@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  100% {
    transform: scale(2);
    opacity: 0;
  }
}
</style>