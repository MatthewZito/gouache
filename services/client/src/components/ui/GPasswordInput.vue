<script lang="ts" setup>
import { required } from '@/utils'
import type { ValidationRule } from 'quasar'
import type { PropType } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    required: true,
  },
  rules: {
    type: Array as PropType<ValidationRule[]>,
    default: () => [required('A password is required.')],
  },
})

const $emit = defineEmits<{
  (e: 'update:modelValue', nextValue: string): void
}>()

const mutableModelValue = computed({
  get() {
    return props.modelValue
  },
  set(v) {
    $emit('update:modelValue', v)
  },
})

const showPassword = ref(false)
</script>

<template lang="pug">
q-input.q-mb-md(
  label="Password"
  v-model="mutableModelValue"
  filled
  :type="showPassword ? 'text' : 'password'"
  autocomplete="current-password"
  dense
  :rules="props.rules"
)
  template(#append)
    q-icon.cursor-pointer(
      :name="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
      @click="showPassword = !showPassword"
    )
</template>
