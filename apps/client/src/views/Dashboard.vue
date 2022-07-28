<script lang="ts" setup>
import EditResource from '@/components/EditResource.vue'
import type { Resource } from '@/types'
import { logger, useErrorHandler } from '@/services'
import CreateResource from '../components/CreateResource.vue'
import { resourceApi } from '@/services'

const headers = [
  {
    name: 'Key',
    label: 'Key',
    field: 'key',
    align: 'left' as const,
  },
  {
    name: 'Value',
    label: 'Value',
    field: 'value',
    align: 'left' as const,
  },
  {
    name: 'Expires',
    label: 'Expires',
    field: 'expires',
    align: 'right' as const,
  },
  {
    field: 'spacer',
    label: '',
    name: 'spacer',
    style: 'width: 99%',
  },
]

// @todo paginate or virtual scroll
const resources = ref<Resource[]>([])

//  { id: generateId(), key: 'key', value: 'test', expires: 10000 },
//   { id: generateId(), key: 'key', value: 'test', expires: 20000 },
//   { id: generateId(), key: 'key', value: 'test', expires: 30000 },
const isLoading = ref(false)
const showEditResourceDialog = ref(false)
const showCreateResourceDialog = ref(false)
const selectedResource = ref<Resource | null>(null)

const paginationConfig = computed(() => ({
  rowsPerPage: 0,
  rowsNumber: resources.value.length,
}))

function handleRowClick(e: Event, row: Resource, idx: number) {
  selectedResource.value = row
  showEditResourceDialog.value = true
}

async function fetchResources() {
  isLoading.value = true

  try {
    const { ok, data, error } = await resourceApi.getResources()
    if (!ok) {
      throw Error(error)
    }

    resources.value = data
  } catch (ex) {
    logger.error('failed to fetch resources', ex)
  } finally {
    isLoading.value = false
  }
}

await fetchResources()

async function handleCreateResourceClosed(refresh: boolean) {
  showCreateResourceDialog.value = false

  if (refresh) {
    // @todo patch
    await fetchResources()
  }
}

async function handleEditResourceClosed(refresh: boolean) {
  showEditResourceDialog.value = false

  if (refresh) {
    // @todo patch
    await fetchResources()
  }
}

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })

  return false
})
</script>

<template>
  <q-table
    dense
    flat
    hide-bottom
    virtual-scroll
    hide-pagination
    class="sticky-table"
    :virtual-scroll-item-size="20"
    :virtual-scroll-sticky-size-start="20"
    :pagination="paginationConfig"
    :rows-per-page-options="[0]"
    :columns="headers"
    :rows="resources"
    @row-click="handleRowClick"
  >
    <template #top-right>
      <q-btn
        label="Create Resource"
        color="primary"
        @click="showCreateResourceDialog = true"
      />
    </template>

    <template #loading>
      <q-inner-loading :showing="isLoading">
        <q-spinner class="q-mb-sm" size="50px" color="secondary" />
        <div class="text-secondary text-bold">Loading...</div>
      </q-inner-loading>
    </template>
  </q-table>

  <Teleport to="#portal">
    <q-dialog
      v-model="showEditResourceDialog"
      @close="showEditResourceDialog = false"
    >
      <EditResource
        v-if="selectedResource"
        :resource="selectedResource"
        @close="handleEditResourceClosed"
      />
    </q-dialog>
  </Teleport>

  <Teleport to="#portal">
    <q-dialog
      v-model="showCreateResourceDialog"
      @close="showCreateResourceDialog = false"
    >
      <CreateResource @close="handleCreateResourceClosed" />
    </q-dialog>
  </Teleport>
</template>

<style lang="scss" scoped>
.sticky-table {
  overflow-y: auto;
  height: 100%;

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
