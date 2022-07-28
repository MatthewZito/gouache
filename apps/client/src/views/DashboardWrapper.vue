<script lang="ts" setup>
import Dashboard from './Dashboard.vue'
import TableSkeleton from '@/components/TableSkeleton.vue'
import { useErrorHandler } from '@/services'

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })
  // do not propagate error to parent
  return false
})
</script>

<template>
  <suspense>
    <template #default>
      <Dashboard />
    </template>
    <template #fallback>
      <TableSkeleton />
    </template>
  </suspense>
</template>
