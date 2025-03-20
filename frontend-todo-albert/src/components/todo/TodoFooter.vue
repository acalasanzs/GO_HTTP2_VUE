<template>
  <div class="p-4 bg-gray-50 text-sm text-gray-600 flex justify-between items-center">
    <span>{{ completedCount }} of {{ todos.length }} tasks completed</span>
    <button 
      v-if="completedCount > 0"
      @click="handleClear" 
      class="text-green-600 hover:text-green-800 focus:outline-none text-sm font-medium"
    >
      Clear completed
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  todos: {
    type: Array,
    required: true,
  },
});
const emit = defineEmits(['clear-completed']);

const completedCount = computed(() => {
  return props.todos.filter(todo => todo.completed).length;
});

const handleClear = () => {
  emit('clear-completed');
};
</script>
