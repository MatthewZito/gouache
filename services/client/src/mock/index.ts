export interface TagOption {
  label: string
  value: string
  description: string
  icon: string
  color: string
}

export const availableTags: TagOption[] = [
  {
    label: 'Art',
    value: 'art',
    description: 'About art',
    icon: 'mdi-palette',
    color: 'accent',
  },
  {
    label: 'Music',
    value: 'music',
    description: 'About music',
    icon: 'mdi-music',
    color: 'accent',
  },
]
