<script lang="ts" setup>
import type { Report } from '@/types'

import { useErrorHandler, reportingApi } from '@/services'
import { toReadableDate } from '@/utils'

import { headers } from './templates'
import ReportDetails from '@/components/Reports/ReportDetails.vue'

const isLoading = ref(false)
const showViewReportDialog = ref(false)
const selectedReport = ref<Report | null>(null)
const reports = ref<Report[]>([])

function handleRowClick(e: Event, row: Report, idx: number) {
  selectedReport.value = row
  showViewReportDialog.value = true
}

async function fetchReports() {
  try {
    isLoading.value = true

    const { ok, data } = await reportingApi.getReports()
    if (!ok) {
      throw Error('@todo')
    }

    reports.value = data.items
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback:
        'Something went wrong and we were unable to fetch your reports.',
    })
  } finally {
    isLoading.value = false
  }
}

await fetchReports()

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })

  return false
})
</script>

<template lang="pug">
q-table.sticky-table(
  dense
  flat
  hide-bottom
  virtual-scroll
  hide-pagination
  :virtual-scroll-item-size="20"
  :virtual-scroll-sticky-size-start="20"
  :rows-per-page-options="[0]"
  :columns="headers"
  :rows="reports"
  @row-click="handleRowClick"
  :loading="isLoading"
)
  template(#body-cell-name="slotProps: { row: Report }")
    q-td(key="name")
      | {{ slotProps.row.name }}

  template(#body-cell-caller="slotProps: { row: Report }")
    q-td(key="caller")
      | {{ slotProps.row.caller }}

  template(#body-cell-data="slotProps: { row: Report }")
    q-td(key="data")
      | {{ slotProps.row.data }}

  template(#body-cell-ts="slotProps: { row: Report }")
    q-td(key="ts")
      | {{ toReadableDate(slotProps.row.ts) }}

  template(#loading)
    q-inner-loading(showing)
      q-spinner.q-mb-sm(
        size="50px"
        color="secondary"
      )
      div.text-secondary.text-bold
        | Loading...

Teleport(to="#portal")
  q-dialog(
    v-model="showViewReportDialog"
    @close="showViewReportDialog = false")

    suspense
      template(#default)
        ReportDetails(
          v-if="selectedReport"
          :report-id="selectedReport.id"
          @close="showViewReportDialog = false"
        )

      template(#fallback)
        q-card.q-px-md.q-pt-md
          q-skeleton(
            type="text"
            style="width: 200px; height: 60px"
          )
          q-skeleton(
            type="text"
            style="width: 200px; height: 60px"
          )
          q-skeleton(type="QBtn")


</template>

<style lang="scss">
.tags:not(:first-child) {
  margin-left: 0.25rem;
}

.sticky-table {
  height: 100%;
  overflow-y: auto;

  thead tr th {
    position: sticky;
    z-index: 1;
    background: white;
  }

  thead tr:first-child th {
    top: 0;
  }
}
</style>
