import { QTableProps } from 'quasar'

export const headers: QTableProps['columns'] = [
  {
    name: 'name',
    label: 'Name',
    field: 'name',
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'caller',
    label: 'Caller',
    field: 'caller',
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'ts',
    label: 'Created At',
    field: 'ts',
    sortable: true,
    align: 'right' as const,
  },

  {
    field: 'spacer',
    label: '',
    name: 'spacer',
    style: 'width: 99%',
  },
]
