<template>
  <div class="p-6 border-b border-gray-200">
    <form v-on:submit="handleSubmit" class="flex gap-2">
      <input 
        v-model="localTodo" 
        type="text" 
        placeholder="Add a new task..." 
        class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
        required
      />
      <button 
        type="submit" 
        class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
      >
        Add
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
});
const emit = defineEmits(['update:modelValue', 'add-todo']);

const localTodo = ref(props.modelValue);

watch(() => props.modelValue, (newVal) => {
  localTodo.value = newVal;
});

const handleSubmit = (event) => {
  event.preventDefault();
  emit('add-todo', localTodo.value);
  emit('update:modelValue', '');
  localTodo.value = '';
};
</script>
