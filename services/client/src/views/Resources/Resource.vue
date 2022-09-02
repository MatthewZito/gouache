<script lang="ts" setup>
import CreateResource from '@/components/Resource/CreateResource.vue'
import EditResource from '@/components/Resource/EditResource.vue'
import { showNotification } from '@/plugins'
import { useErrorHandler, resourceApi } from '@/services'
import { useResourceStore } from '@/state'
import type { Resource } from '@/types'
import { toReadableDate } from '@/utils'

import { headers } from './templates'

const resourceStore = useResourceStore()

const isLoading = ref(false)
const showEditResourceDialog = ref(false)
const showCreateResourceDialog = ref(false)
const selectedResource = ref<Resource | null>(null)

const paginationConfig = computed(() => ({
  rowsPerPage: 0,
  rowsNumber: resourceStore.resources.length,
}))

function handleRowClick(e: Event, row: Resource, idx: number) {
  selectedResource.value = row
  showEditResourceDialog.value = true
}

async function handleDeleteResource(e: Event, row: Resource) {
  isLoading.value = true
  try {
    const { ok } = await resourceApi.deleteResource(row.id)

    if (!ok) {
      throw Error('failed to delete resource')
    }

    showNotification('success', 'Successfully deleted resource.')

    resourceStore.removeResource(row.id)
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback: 'Something went wrong while deleting this resource.',
    })
  } finally {
    isLoading.value = false
  }
}

async function fetchResources() {
  isLoading.value = true

  try {
    const { ok, data } = await resourceApi.getResources()

    if (!ok) {
      throw Error('@todo')
    }

    resourceStore.setResources(data ?? [])
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback:
        'Something went wrong and we were unable to fetch your resources.',
    })
  } finally {
    isLoading.value = false
  }
}

await fetchResources()

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
  :pagination="paginationConfig"
  :rows-per-page-options="[0]"
  :columns="headers"
  :rows="resourceStore.resources"
  @row-click="handleRowClick"
  :loading="isLoading"
)
  template(#top-right)
    q-btn(
      label="Create Resource"
      color="primary"
      @click="showCreateResourceDialog = true"
    )

  template(#body-cell-createdAt="slotProps: { row: Resource }")
    q-td(key="createdAt")
      | {{ toReadableDate(slotProps.row.createdAt) }}

  template(#body-cell-updatedAt="slotProps: { row: Resource }")
    q-td(key="updatedAt")
      | {{ toReadableDate(slotProps.row.updatedAt) }}

  template(#body-cell-tags="slotProps: { row: Resource }")
    q-td(
      key="tags"
      :props="slotProps"
    )
      span.tags(v-for="tag in slotProps.row.tags")
        q-badge
          | {{ tag }}

  template(#body-cell-action="slotProps: { row: Resource }")
    q-td(
      key="action"
      :props="slotProps"
    )
      q-btn(
        icon="mdi-delete"
        size="sm"
        flat
        round
        color="negative"
        @click.prevent.stop="e => handleDeleteResource(e, slotProps.row)"
      )

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
    v-model="showEditResourceDialog"
    @close="showEditResourceDialog = false")

    suspense
      template(#default)
        EditResource(
          v-if="selectedResource"
          :resource-id="selectedResource.id"
          @close="showEditResourceDialog = false"
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

Teleport(to="#portal")
  q-dialog(
    v-model="showCreateResourceDialog"
    @close="showCreateResourceDialog = false"
  )
    CreateResource(
      @close="showCreateResourceDialog = false"
    )
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
