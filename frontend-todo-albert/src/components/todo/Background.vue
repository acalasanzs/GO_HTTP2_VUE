<template>
  <div class="absolute inset-0 z-0 overflow-hidden">
    <div class="absolute inset-0 opacity-20">
      <div 
        v-for="(item, i) in randomValues" 
        :key="i" 
        class="absolute rounded-lg bg-green-500"
        :style="{
          width: `${item.width}px`,
          height: `${item.height}px`,
          top: `${item.top}%`,
          left: `${item.left}%`,
          transform: `rotate(${item.rotate}deg)`,
          opacity: item.opacity
        }"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const generateRandomValue = () => ({
  width: Math.random() * 200 + 50,
  height: Math.random() * 200 + 50,
  top: Math.random() * 100,
  left: Math.random() * 100,
  rotate: Math.random() * 360,
  opacity: Math.random() * 0.5 + 0.1,
});

const randomValues = ref(Array.from({ length: 10 }, () => generateRandomValue()));
</script>
