<script lang="ts" setup>
import EditResource from '@/components/EditResource.vue'
import type { Resource } from '@/types'
import { useErrorHandler } from '@/services'
import CreateResource from '@/components/CreateResource.vue'
import { resourceApi } from '@/services'
import { toReadableDate } from '@/utils'
import { headers } from './templates'
import { useResourceStore } from '@/state'

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

async function fetchResources() {
  isLoading.value = true

  try {
    const { ok, data } = await resourceApi.getResources()
    if (!ok) {
      throw Error()
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
    :rows="resourceStore.resources"
    @row-click="handleRowClick"
  >
    <template #top-right>
      <q-btn
        label="Create Resource"
        color="primary"
        @click="showCreateResourceDialog = true"
      />
    </template>

    <template #body-cell-createdAt="slotProps: { row: Resource }">
      <q-td key="createdAt">
        {{ toReadableDate(slotProps.row.createdAt) }}
      </q-td>
    </template>

    <template #body-cell-updatedAt="slotProps: { row: Resource }">
      <q-td key="updatedAt">
        {{ toReadableDate(slotProps.row.updatedAt) }}
      </q-td>
    </template>

    <template #body-cell-tags="slotProps: { row: Resource }">
      <q-td key="tags" :props="slotProps">
        <span v-for="tag in slotProps.row.tags" class="tags">
          <q-badge>
            {{ tag }}
          </q-badge>
        </span>
      </q-td>
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
      <suspense>
        <template #default>
          <EditResource
            v-if="selectedResource"
            :resource-id="selectedResource.id"
            @close="showEditResourceDialog = false"
          />
        </template>

        <template #fallback>
          <q-card class="q-px-md q-pt-md">
            <q-skeleton type="text" style="width: 200px; height: 60px" />
            <q-skeleton type="text" style="width: 200px; height: 60px" />
            <q-skeleton type="QBtn" />
          </q-card>
        </template>
      </suspense>
    </q-dialog>
  </Teleport>

  <Teleport to="#portal">
    <q-dialog
      v-model="showCreateResourceDialog"
      @close="showCreateResourceDialog = false"
    >
      <CreateResource @close="showCreateResourceDialog = false" />
    </q-dialog>
  </Teleport>
</template>

<style lang="scss" scoped>
.tags:not(:first-child) {
  margin-left: 0.25rem;
}

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
