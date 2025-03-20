<template>
  <div class="min-h-screen bg-gradient-to-br from-green-50 to-green-100 flex items-center justify-center p-4 sm:p-6 md:p-8">
    <!-- Geometric background pattern -->
    <Background />

    <!-- Todo list container -->
    <div class="relative z-10 w-full max-w-md bg-white rounded-xl shadow-xl overflow-hidden">
      <TodoHeader title="Todo List" />
      <TodoForm @add-todo="handleAddTodo" />
      <TodoList 
        :todos="todoStore.todos" 
        @toggle-todo="handleToggleTodo" 
        @remove-todo="handleRemoveTodo" 
      />
      <TodoFooter 
       :todos="todoStore.todos" 
        :completedCount="todoStore.completedCount" 
        :total="todoStore.todos.length" 
        @clear-completed="handleClearCompleted" 
      />
    </div>
  </div>
</template>

<script setup>
import { useTodoStore } from '@src/stores/todoStore';
import Background from '@src/components/todo/Background.vue';
import TodoHeader from '@src/components/todo/TodoHeader.vue';
import TodoForm from '@src/components/todo/TodoForm.vue';
import TodoList from '@src/components/todo/TodoList.vue';
import TodoFooter from '@src/components/todo/TodoFooter.vue';

const todoStore = useTodoStore();

const handleAddTodo = (todoText) => {
  todoStore.addTodo(todoText);
};

const handleToggleTodo = (index) => {
  todoStore.toggleTodo(index);
};

const handleRemoveTodo = (index) => {
  todoStore.removeTodo(index);
};

const handleClearCompleted = () => {
  todoStore.clearCompleted();
};
</script>
