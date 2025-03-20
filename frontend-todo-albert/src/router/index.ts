import { createRouter, createWebHistory } from 'vue-router'
import TodoView from '@src/views/TodoView.vue'

const routes = [
  {
    path: '/',
    name: 'Todo',
    component: TodoView,
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/AboutView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
