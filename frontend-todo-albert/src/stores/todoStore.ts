import { defineStore } from 'pinia'

export const useTodoStore = defineStore('todo', {
  state: () => ({
    todos: [
      { text: 'Learn Vue 3', completed: true, id: 1 },
      { text: 'Build a todo app', completed: false, id: 2 },
      { text: 'Master Tailwind CSS', completed: false, id: 3 },
    ],
    id: 3,
  }),
  getters: {
    completedCount(state) {
      return state.todos.filter((todo) => todo.completed).length
    },
  },
  actions: {
    addTodo(todoText) {
      if (todoText.trim()) {
        this.todos.push({ text: todoText.trim(), completed: false, id: ++this.id })
      }
    },
    toggleTodo(index) {
      const res = this.todos.find((x) => x.id === index)
      if (!res) throw new Error('Something :)')
      res.completed = !res.completed
    },
    removeTodo(index) {
      this.todos.splice(this.todos.findIndex(x=>x.id===index), 1)
    },
    clearCompleted() {
      this.todos = this.todos.filter((todo) => !todo.completed)
    },
  },
})
