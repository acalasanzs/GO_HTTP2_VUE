<template>
  <div class="divide-y divide-gray-200 max-h-96 overflow-y-auto">
    <div v-if="todos.length === 0" class="p-6 text-center text-gray-500">
      No tasks yet. Add one above!
    </div>
    <TodoItem 
      v-for="(todo, index) in todos" 
      :key="index" 
      :todo="todo"
      :index="index"
      @toggle-todo="toggleTodo"
      @remove-todo="removeTodo"
    />
  </div>
</template>

<script setup>
import TodoItem from './TodoItem.vue';

const props = defineProps({
  todos: {
    type: Array,
    required: true,
  },
});
const emit = defineEmits(['toggle-todo', 'remove-todo']);

const toggleTodo = (index) => {
  emit('toggle-todo', index);
};

const removeTodo = (index) => {
  emit('remove-todo', index);
};
</script>
