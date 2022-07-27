<script lang="ts" setup>
import styled from '@magister_zito/vue3-styled-components'
import { v4 as generateId } from 'uuid'

const StyledInput = styled.input`
  padding: 0.5rem;
  border-radius: 0.25rem;
  border: 1px solid green;
  background-color: lightgreen;

  &:focus-visible {
    border: 1px solid darkgreen;
    outline: none;
  }
`

const StyledLabel = styled.label`
  font-size: 12px;
  padding: 0.25rem;
  padding-bottom: 0.1rem;
`

const props = defineProps({
  modelValue: {
    type: [String, Number],
    required: true,
  },
  label: {
    type: String,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'update:modelValue', nextValue: string | number): void
}>()

const $attrs = useAttrs()

const id = ref(generateId())

const mutableModelValue = computed({
  get() {
    return props.modelValue
  },
  set(nextValue: string | number) {
    $emit('update:modelValue', nextValue)
  },
})
</script>

<template>
  <StyledLabel :for="id">
    {{ props.label }}
  </StyledLabel>
  <StyledInput :id="id" v-model="mutableModelValue" v-bind="$attrs" />
</template>
