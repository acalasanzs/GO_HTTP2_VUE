<template>
  <div ref="item" class="p-4 flex items-center gap-3 hover:bg-green-50 transition-colors">
    <input
      type="checkbox"
      :checked="todo.completed"
      @change="handleToggle"
      class="h-5 w-5 rounded border-gray-300 text-green-600 focus:ring-green-500"
    />
    <span
      :class="{ 'line-through text-gray-400': todo.completed }"
      class="flex-1"
    >
      {{ todo.text }}
    </span>
    <button
      @click="handleRemove"
      class="text-red-500 hover:text-red-700 focus:outline-none"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M3 6h18"></path>
        <path
          d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"
        ></path>
        <line x1="10" y1="11" x2="10" y2="17"></line>
        <line x1="14" y1="11" x2="14" y2="17"></line>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { useTemplateRef } from 'vue';

const item = useTemplateRef("item")
const props = defineProps({
  todo: {
    type: Object,
    required: true,
  },
  index: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits(['toggle-todo', 'remove-todo']);

const handleToggle = () => {
  emit('toggle-todo', props.index, item.value);
};

const handleRemove = () => {
  emit('remove-todo', props.index, item.value);
};
</script>