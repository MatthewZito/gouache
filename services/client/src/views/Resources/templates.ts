export const headers = [
  {
    name: 'title',
    label: 'Title',
    field: 'title',
    align: 'left' as const,
  },
  {
    name: 'createdAt',
    label: 'Created At',
    field: 'createdAt',
    align: 'right' as const,
  },
  {
    name: 'updatedAt',
    label: 'Updated At',
    field: 'updatedAt',
    align: 'right' as const,
  },
  {
    name: 'tags',
    label: 'Tags',
    field: 'tags',
    align: 'left' as const,
  },

  {
    field: 'spacer',
    label: '',
    name: 'spacer',
    style: 'width: 99%',
  },
  {
    field: 'action',
    label: '',
    name: 'action',
  },
]
