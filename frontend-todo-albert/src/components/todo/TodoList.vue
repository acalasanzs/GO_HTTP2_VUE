<template>
  <div class="divide-y divide-gray-200 max-h-96 overflow-y-auto">
    <div v-if="todos.length === 0" class="p-6 text-center text-gray-500">
      No tasks yet. Add one above!
    </div>
    <div class="todo-container" ref="container">
      <TransitionGroup
        name="todo"
        tag="div"
        :css="false"
        @leave="onLeave"
        @before-enter="onBeforeEnter"
        @enter="onEnter"
      >
        <TodoItem
          v-for="todo in todos"
          :key="todo.id"
          :index="todo.id"
          :todo="todo"
          @toggle-todo="toggleTodo"
          @remove-todo="removeTodo"
        />
      </TransitionGroup>
    </div>
  </div>
</template>

<script setup lang="ts">
import gsap from 'gsap'
import { ref, reactive, onUnmounted, TransitionGroup } from 'vue'
import TodoItem from './TodoItem.vue'

const container = ref(null)
const animatingElements = reactive(new Map())

const props = defineProps({
  todos: {
    type: Array,
    required: true,
  },
})

function onBeforeEnter(el) {
  el.style.opacity = 0
  el.style.transform = 'translateY(-20px)'
}

function onEnter(el, done) {
  gsap.to(el, {
    duration: 0.4,
    y: 0,
    opacity: 1,
    ease: 'power1.inOut',
    onComplete: done,
  })
}

function onLeave(el, done) {
  // Store reference to the element being animated
  const elementId = el.getAttribute('data-id')

  if (animatingElements.has(elementId)) {
    // If this element is already being animated, update its animation
    const existingAnim = animatingElements.get(elementId)
    existingAnim.kill() // Kill the existing animation
  }

  // Position the element absolutely for the exit animation
  el.style.position = 'absolute'
  el.style.width = '100%'
  el.style.zIndex = '10'

  // Find sibling elements that need to be repositioned
  const siblings = Array.from(el.parentNode.children).filter((x) => {
    return (
      x !== el &&
      x.getBoundingClientRect().bottom > el.getBoundingClientRect().top &&
      !x.style.position.includes('absolute')
    )
  }) as HTMLDivElement[]

  const animations = []

  // Animate siblings if needed
  if (siblings.length) {
    const firstSibling = siblings[0]
    const elHeight = el.getBoundingClientRect().height

    // Set initial state for the animation
    firstSibling.style.transform = `translateY(${elHeight}px)`
    firstSibling.style.marginBottom = `${elHeight}px`

    // Create animation for the sibling
    const siblingAnim = gsap.to(firstSibling, {
      duration: 0.4,
      ease: 'power1.inOut',
      marginBottom: '0px',
      y: 0,
      delay: 0.05,
    })

    animations.push(siblingAnim)
  } else if (container.value) {
    // If no siblings, add padding to the container
    const elHeight = el.getBoundingClientRect().height
    container.value.style.paddingBottom = `${elHeight}px`

    // Create animation for the container padding
    const containerAnim = gsap.to(container.value, {
      duration: 0.4,
      ease: 'power1.inOut',
      paddingBottom: '0px',
    })

    animations.push(containerAnim)
  }

  // Create animation for the element itself
  const elementAnim = gsap.to(el, {
    opacity: 0,
    duration: 0.7,
    ease: 'power1.inOut',
    x: '-100%',
    onComplete: () => {
      // Remove from tracking when animation completes
      animatingElements.delete(elementId)
      done()
    },
  })

  animations.push(elementAnim)

  // Store the animation for potential future updates
  animatingElements.set(elementId, {
    element: el,
    kill: () => animations.forEach((a) => a.kill()),
  })
}

const emit = defineEmits(['toggle-todo', 'remove-todo'])

const toggleTodo = (index, item) => {
  // Set a data-id attribute to identify this element
  if (item && !item.hasAttribute('data-id')) {
    item.setAttribute('data-id', `todo-${index}`)
  }
  emit('toggle-todo', index)
}

const removeTodo = (index, item) => {
  // Set a data-id attribute to identify this element
  if (item && !item.hasAttribute('data-id')) {
    item.setAttribute('data-id', `todo-${index}`)
  }
  emit('remove-todo', index)
}

// Cleanup any ongoing animations when component is unmounted
onUnmounted(() => {
  animatingElements.forEach(({ kill }) => kill())
  animatingElements.clear()
})
</script>

<style scoped>
.todo-container {
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  min-height: 20px; /* Ensure container has height even when empty */
}
</style>
