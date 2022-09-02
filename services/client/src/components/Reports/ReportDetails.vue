<script lang="ts" setup>
import type { PropType } from 'vue'

import { ErroneousResponseError, reportingApi } from '@/services'
import type { UUID } from '@/types/scalar'
import { toReadableDate } from '@/utils'

const props = defineProps({
  reportId: {
    type: String as PropType<UUID>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const { ok: topLevelOk, data: report } = await reportingApi.getReport(
  props.reportId,
)

if (!topLevelOk) {
  throw new ErroneousResponseError(
    'Something went wrong while loading the details for this report.',
  )
}
</script>

<template lang="pug">
q-card.q-pa-sm(
  style="width: 400px"
)
  q-card-section
    div.text-h6
      | Report Details

  q-card-section
    table.report-table
      tr
        th
          | Report Name
        td
          | {{ report.name }}
      tr
        th
          | Originating System
        td
          | {{ report.caller }}
      tr
        th
          | Report Created At
        td
          | {{ toReadableDate(report.ts) }}

  q-card-section
    div.text-subtitle2.text-bold.q-pb-sm
      | Report Data (raw)
    div(style="background-color:lightgray;border-radius:4px; padding:2rem")
      span(style="white-space:pre;color:black;font-weight:700")
        | {{ JSON.parse(report.data) }}

  q-card-actions.q-pa-none.q-pt-sm.justify-between
    q-btn(
      label="Close"
      flat
      color="grey-6"
      @click="$emit('close')"
    )
</template>

<style lang="scss">
/*
  For some reason scoped breaks pug.
  @todo Evaluate fix https://github.com/vitejs/vite/issues/9407
*/
.report-table {
  width: 100%;

  & th {
    text-align: left;
  }
  & tr {
    text-align: right;
  }
}
</style>
