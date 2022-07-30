<script lang="ts" setup>
import { listRequired } from '@/utils'
import { availableTags } from '@/mock'
import type { PropType } from 'vue'

const props = defineProps({
  modelValue: {
    type: Array as PropType<string[]>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'update:modelValue', nextValue: typeof props.modelValue): void
}>()

const mutableModelValue = computed({
  get() {
    return props.modelValue
  },
  set(v: string[]) {
    console.log({ v })
    $emit('update:modelValue', v)
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
    :options="availableTags"
    :rules="[listRequired('At least one tag is required.')]"
    label="Tags"
    option-value="value"
    class="mr-2 extra-dense"
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
