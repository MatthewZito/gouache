<script lang="ts" setup>
import styled from '@magister_zito/vue3-styled-components'
import type { PropType } from 'vue'

const StyledTable = styled.table`
  border-collapse: collapse;
  width: 100%;
  max-width: 80%;
  height: 100%;
`

const StyledTableHeader = styled.th`
  padding: 0.75rem;
  text-align: left;
  background-color: #04aa6d;
  border: 1px solid #04bb6e;
  color: white;
`
const StyledTableRow = styled('tr', {
  isHeader: { type: Boolean, default: false },
})`
  transition: all 0.25s;
  ${({ isHeader }) =>
    !isHeader
      ? `
  cursor: pointer;
  &:hover {
    background-color: #ddd;
  }
`
      : ''}
`

const StyledTableData = styled.td`
  padding: 0.75rem;
  border: 1px solid #ddd;
`

type HeaderAlignment = 'left' | 'right' | 'center'

interface Header {
  label: string
  field: string
  align: HeaderAlignment
}

type Column = any

const props = defineProps({
  headers: {
    type: Array as PropType<Header[]>,
    required: true,
  },
  // @todo paginate or virtual scroll
  columns: {
    type: Array as PropType<Column[]>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'rowClick', row: Column, idx: number): void
}>()
</script>

<template>
  <StyledTable>
    <StyledTableRow :is-header="true">
      <StyledTableHeader v-for="{ label } in props.headers">
        {{ label }}
      </StyledTableHeader>
      <StyledTableHeader style="width: 99%" />
    </StyledTableRow>

    <StyledTableRow
      v-for="(column, idx) in props.columns"
      :key="idx"
      :is-header="false"
      @click="$emit('rowClick', column, idx)"
    >
      <StyledTableData
        v-for="{ align, field } in props.headers"
        :style="`text-align: ${align}`"
      >
        {{ column[field as keyof typeof column] }}
      </StyledTableData>
      <StyledTableData style="width: 99%" />
    </StyledTableRow>
  </StyledTable>
</template>
