<script lang="ts" setup>
import type { ValidationRule } from 'quasar'
import type { PropType } from 'vue'

interface SelectOption {
  value: string
  label: string
  description: string
  icon: string
  color: string
}

type State = typeof props.modelValue

const props = defineProps({
  modelValue: {
    type: Array as PropType<string[]>,
    required: true,
  },
  options: {
    type: Array as PropType<SelectOption[]>,
    required: true,
  },
  rules: {
    type: Array as PropType<ValidationRule[]>,
    default: () => [],
  },
})

const $emit = defineEmits<{
  (e: 'update:modelValue', nextValue: State): void
}>()

const mutableModelValue = computed({
  get() {
    return props.modelValue
  },
  set(v) {
    $emit('update:modelValue', v ?? [])
  },
})

function removeOption(value: string) {
  const foundIdx = props.modelValue.findIndex(tag => tag === value)
  if (foundIdx !== -1) {
    props.modelValue.splice(foundIdx, 1)
  }
}
</script>

<template>
  <q-select
    v-model="mutableModelValue"
    :options="props.options"
    :rules="props.rules"
    option-value="value"
    filled
    clearable
    dense
    emit-value
    map-options
    multiple
    use-chips
  >
    <template #selected-item="{ opt }">
      <q-chip
        :color="opt.color"
        text-color="black"
        removable
        @remove="_ => removeOption(opt.value)"
      >
        <span class="q-mr-sm">
          {{ opt.label }}
        </span>
        <q-icon :name="opt.icon" />
        <q-tooltip>
          {{ opt.description }}
        </q-tooltip>
      </q-chip>
    </template>
  </q-select>
</template>
