<template>
  <div class="bg-green-600 p-6">
    <h1 class="text-2xl font-bold text-white">{{ title }}</h1>
  </div>
</template>

<script setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
});
</script>
