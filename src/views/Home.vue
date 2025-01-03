<script setup lang="ts">
import { computed, onMounted, onUnmounted, Ref, ref } from "vue";
import { getStatuses, Status } from "../api";
import Host from "../components/Host.vue";
import { onlineTimeout } from "../api/utils.ts";

const statuses: Ref<Record<string, Status>> = ref({});
const offline = ref(0);
const showOptions = ref(false);
const selectedOption = ref('');
const sortedKey = ref('name');
const reverse = ref(false);

const toggleOptions = () => {
  showOptions.value = !showOptions.value;
};

const selectOption = (option: string) => {
  selectedOption.value = option;
  sortedKey.value = option;
};

const onlineNum = computed(() => {
  const nowTimestamp = Date.now() / 1000;
  let online = 0;
  for (const status of Object.values(statuses.value)) {
    if (nowTimestamp - status.meta.observed_at < onlineTimeout) {
      online++;
    }
  }
  offline.value = Object.values(statuses.value).length - online;
  return online;
});

function sortByName(statusMap: Record<string, Status>) {
  return Object.fromEntries(
      Object.entries(statusMap).sort((a, b) => a[1].meta.name.localeCompare(b[1].meta.name))
  );
}

function sortByUpTime(statusMap: Record<string, Status>) {
  return Object.fromEntries(
      Object.entries(statusMap).sort((a, b) => a[1].meta.uptime - b[1].meta.uptime)
  );
}

function updateSort(statusMap: Record<string, Status>) {
  let sortedMap = statusMap;
  if (sortedKey.value === 'name') {
    sortedMap = sortByName(statusMap);
  } else if (sortedKey.value === 'uptime') {
    sortedMap = sortByUpTime(statusMap);
  } else if (sortedKey.value === 'memory') {
    sortedMap = Object.fromEntries(
        Object.entries(statusMap).sort((a, b) => a[1].hardware.mem.total - b[1].hardware.mem.total)
    );
  } else if (sortedKey.value === 'network') {
    sortedMap = Object.fromEntries(
        Object.entries(statusMap).sort((a, b) => a[1].hardware.net.down + a[1].hardware.net.up - b[1].hardware.net.down - b[1].hardware.net.up)
    );
  }


  if (reverse.value) {
    sortedMap = Object.fromEntries(Object.entries(sortedMap).reverse());
  }
  return sortedMap;
}

const timer = setInterval(async () => {
  statuses.value = updateSort(await getStatuses());
}, 1000);

onMounted(async () => {
  statuses.value = updateSort(await getStatuses());
});

onUnmounted(() => {
  clearInterval(timer);
});
</script>

<template>
  <div class="overview">
    <h2>Overview: {{ onlineNum }} Online {{ offline }} Offline</h2>
  </div>
  <div class="tabs" style="display: flex">
    <button class="button-a" @click="toggleOptions">Sort by</button>
    <transition name="slide-fade">
      <div v-if="showOptions" class="options">
        <button :class="{ selected: selectedOption === 'name' }" @click="() => selectOption('name')">Name</button>
        <button :class="{ selected: selectedOption === 'uptime' }" @click="() => selectOption('uptime')">Uptime</button>
        <button :class="{ selected: selectedOption === 'memory' }" @click="() => selectOption('memory')">Memory</button>
        <button :class="{ selected: selectedOption === 'network' }" @click="() => selectOption('network')">Network</button>
<!--        <button :class="{ selected: selectedOption === 'cpu' }" @click="() => selectOption('cpu')">CPU Percent</button>-->
        <button :class="{ selected: reverse }" @click="() => reverse= !reverse" id="reverse-button">Reverse</button>
      </div>
    </transition>
  </div>
  <div class="grid-container">
    <Host class="grid-item" v-for="(status, id) in statuses" :key="id" :status="status" />
  </div>
</template>

<style scoped>

.tabs{
  margin-bottom: 1rem;
}

.button-a{
  margin-left: 10px;
}

button {
  margin: 10px 10px 0 0;
  padding: 0.5rem 0.5rem;
  border-radius: 50px;
  border: none;
  background: #36a7ec;

}
#reverse-button {
  background: #05c860;
}
#reverse-button.selected {
  background: #ff6347;
}

button.selected {
  background: #ff6347; /* Selected button background color */
}

.overview {
  display: flex;
  justify-content: center;
  padding: 10px;
}

.grid-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
  padding: 10px;
}

.options {
  display: flex;
  gap: 10px;
}

.slide-fade-enter-active, .slide-fade-leave-active {
  transition: all 0.5s ease;
}

.slide-fade-enter-from, .slide-fade-leave-to {
  transform: translateX(-100%);
  opacity: 0;
}
</style>