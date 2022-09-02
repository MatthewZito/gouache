<script lang="ts" setup>
import TableSkeleton from '@/components/TableSkeleton.vue'
import { useErrorHandler } from '@/services'

import Reporting from './Reporting.vue'

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })
  // do not propagate error to parent
  return false
})
</script>

<template lang="pug">
suspense
  template(#default)
    Reporting
  template(#fallback)
    TableSkeleton
</template>
