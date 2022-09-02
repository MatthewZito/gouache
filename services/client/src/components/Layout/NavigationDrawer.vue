<script lang="ts" setup>
interface NavigationRecord {
  icon: string
  title: string
  routeName: string
}

const props = defineProps<{
  modelValue: boolean
}>()

const $emit = defineEmits<{
  (e: 'update:modelValue', nextValue: boolean): void
}>()

const $router = useRouter()

const mutableModelValue = computed({
  get() {
    return props.modelValue
  },
  set(nextValue: boolean) {
    $emit('update:modelValue', nextValue)
  },
})

const navigationRecords: NavigationRecord[] = [
  {
    icon: 'mdi-home',
    title: 'Dashboard',
    routeName: 'Dashboard',
  },
  {
    icon: 'mdi-database',
    title: 'Resources',
    routeName: 'Resource',
  },
  {
    icon: 'mdi-file-chart',
    title: 'Reporting',
    routeName: 'Reporting',
  },
]
</script>

<template lang="pug">
q-drawer(
  show-if-above
  v-model="mutableModelValue"
  side="left"
  bordered
)
  q-list.menu-list(
    padding
  )
    q-item(
      v-for="({ icon, title, routeName }, idx) in navigationRecords"
      clickable
      v-ripple
      :key="idx"
      @click="$router.push({ name: routeName })"
    )
      q-item-section(
        avatar
      )
        q-icon(
          :name="icon"
        )

      q-item-section
        | {{ title }}
</template>
