<script lang="ts" setup>
import TableSkeleton from '@/components/TableSkeleton.vue'
import { useErrorHandler } from '@/services'

import Resource from './Resource.vue'

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
