<script lang="ts" setup>
import Resource from './Resource.vue'
import TableSkeleton from '@/components/TableSkeleton.vue'
import { useErrorHandler } from '@/services'

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })
  // do not propagate error to parent
  return false
})
</script>

<template lang="pug">
suspense
  template(#default)
    Resource
  template(#fallback)
    TableSkeleton
</template>
