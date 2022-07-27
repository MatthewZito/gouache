<script lang="ts" setup>
import styled from '@magister_zito/vue3-styled-components'
import Table from '@/components/ui/Table.vue'
import Modal from '@/components/ui/Modal.vue'
import EditResource from '@/components/Resource/EditResource.vue'

const Container = styled.div`
  width: 100%;
  display: flex;
  justify-content: center;
`

interface Resource {
  key: string
  value: string
  expires: number
}

const headers = [
  { label: 'Key', field: 'key', align: 'left' as const },
  { label: 'Value', field: 'value', align: 'left' as const },
  { label: 'Expires', field: 'expires', align: 'right' as const },
]

// @todo paginate or virtual scroll
const resources = ref<Resource[]>([
  { key: 'key', value: 'test', expires: 10000 },
  { key: 'key', value: 'test', expires: 10000 },
  { key: 'key', value: 'test', expires: 10000 },
])

const showModal = ref(false)
const selectedRow = ref<Resource | null>(null)

function handleRowClick(row: Resource, idx: number) {
  selectedRow.value = row
  showModal.value = true
}

function handleCloseModal() {
  selectedRow.value = null
  showModal.value = false
}
</script>

<template>
  <Container>
    <Table
      :headers="headers"
      :columns="resources"
      @row-click="handleRowClick"
    />
  </Container>

  <Teleport to="#portal">
    <Modal v-if="showModal && selectedRow" @close="showModal = false">
      <template #content>
        <EditResource :selected-row="selectedRow" @close="handleCloseModal" />
      </template>
    </Modal>
  </Teleport>
</template>
