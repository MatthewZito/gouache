<script lang="ts" setup>
import styled from '@magister_zito/vue3-styled-components'
import { cycleRange } from '@/utils'
// import { v4 as generateId } from 'uuid'

const ModalOverlay = styled.div.attrs({
  role: 'dialog',
  ariaModal: true,
})`
  position: fixed;
  z-index: 9999;
  top: 0;
  left: 0;
  display: flex;
  width: 100vw;
  height: 100vh;
  flex-direction: column;
  backdrop-filter: blur(4px);

  background-color: rgba(0, 0, 0, 0.8);

  @media (min-width: 640px) {
    padding: 1.5rem;
  }

  @media (min-width: 768px) {
    padding: 10vh;
  }

  @media (min-width: 1024px) {
    padding: 12vh;
  }
`
const Container = styled.div`
  display: flex;
  width: 450px;
  min-height: 0;
  flex-direction: column;
  margin: auto;
  background-color: white;
  border-radius: 0.5rem;
`

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const overlayRef = ref<(HTMLDivElement & { $el: HTMLDivElement }) | null>(null)
const modalRef = ref<HTMLDivElement | null>(null)
// const modalId = ref(generateId())

function handleClick(e: MouseEvent) {
  if (e.target === overlayRef.value?.$el) {
    $emit('close')
  }
}

const keyListenersMap = new Map([
  [
    'Escape',
    () => {
      $emit('close')
    },
  ],
  ['Tab', handleTabKey],
])

function handleTabKey(e: KeyboardEvent) {
  // @todo hoist this up
  const focusableModalElements = Array.from(
    modalRef.value?.querySelectorAll?.<HTMLElement>(
      'li, button, textarea, input[type="text"], input[type="radio"], input[type="checkbox"], input[type="search"], select',
    ) || [],
  )

  if (!focusableModalElements.length) {
    return
  }
  const currentActive = document.activeElement
  const currentIdx = focusableModalElements.findIndex(
    el => el === currentActive,
  )

  const next = cycleRange(
    e.shiftKey ? -1 : 1,
    currentIdx,
    focusableModalElements.length,
  )

  const nextEl = focusableModalElements[next]

  nextEl.focus()

  e.preventDefault()
}

watch(modalRef, () => {
  const focusableModalElements =
    modalRef.value?.querySelectorAll?.<HTMLElement>(
      'a[href], button, textarea, input[type="text"], input[type="radio"], input[type="checkbox"], input[type="search"], select',
    )

  if (!focusableModalElements?.length) {
    return
  }

  const firstElement = focusableModalElements[0]

  firstElement.focus()
})

function listener(e: KeyboardEvent) {
  const fn = keyListenersMap.get(e.key)

  fn?.(e)
}

onMounted(() => {
  document.addEventListener('keydown', listener)
})

onBeforeUnmount(() => {
  document.removeEventListener('keydown', listener)
})
</script>

<template>
  <ModalOverlay @click="handleClick" ref="overlayRef">
    <Container ref="modalRef">
      <slot name="content" />
    </Container>
  </ModalOverlay>
</template>
