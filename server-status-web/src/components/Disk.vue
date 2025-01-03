<script setup lang="ts">
import {format2Size, getBaseColor, getBlankColor} from '../api/utils.ts';

const props = defineProps<{
  mountpoint: string;
  device: string;
  used: number;
  total: number;
  fstype: string;
}>();

const colorUsed = getBaseColor(props.used / props.total * 100);
const colorBlank = getBlankColor(props.used / props.total * 100);
</script>

<template>

  <div class="disk">
    <div class="hover-text">
      <div class="left-text">
        <span class="disk-text">{{ props.mountpoint }}</span>
      </div>
      <div class="right-text">
        <span class="disk-text">{{ format2Size(props.used, props.total) }} [{{ props.fstype }}]</span>
      </div>
    </div>
    <div class="progress-bar" style="display: flex; align-items: center">
      <div class="disk-total" :style="{ backgroundColor: colorBlank }" style="margin-right: 0.5rem;">
        <div class="disk-used"
             :style="{ width: props.used / props.total * 100 + '%', backgroundColor: colorUsed }"></div>
      </div>
      <div style="color: var(--text-color-2)">
        <div class="percentage" style="text-align: right">{{ (props.used / props.total * 100).toFixed(1) }}%</div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.disk {
  margin-top: 0.75rem;
}

.disk-total {
  height: 0.618rem;
  width: 90%;
  border-radius: 1rem;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  overflow: hidden; /* Ensure the used part doesn't overflow */
}

.disk-used {
  height: 100%;
  left: 0;
  top: 0;
  clip-path: inset(0 0 0 0 round 1rem);
}

.hover-text {
  width: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  text-align: center;
}

.left-text {
  display: flex;
  justify-content: flex-start;
}

.right-text {
  display: flex;
  justify-content: flex-end;
}

.disk-text {
  font-size: 16px;
  color: var(--text-color-2);
}

.percentage{
  font-size: 14px;
  color: var(--text-color-2);
}
</style>