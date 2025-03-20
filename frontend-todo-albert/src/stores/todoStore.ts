import { defineStore } from 'pinia'

export const useTodoStore = defineStore('todo', {
  state: () => ({
    todos: [
      { text: 'Learn Vue 3', completed: true },
      { text: 'Build a todo app', completed: false },
      { text: 'Master Tailwind CSS', completed: false },
    ],
  }),
  getters: {
    completedCount(state) {
      return state.todos.filter((todo) => todo.completed).length
    },
  },
  actions: {
    addTodo(todoText) {
      if (todoText.trim()) {
        this.todos.push({ text: todoText.trim(), completed: false })
      }
    },
    toggleTodo(index) {
      this.todos[index].completed = !this.todos[index].completed
    },
    removeTodo(index) {
      this.todos.splice(index, 1)
    },
    clearCompleted() {
      this.todos = this.todos.filter((todo) => !todo.completed)
    },
  },
})
