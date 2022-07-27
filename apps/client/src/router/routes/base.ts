export default [
  {
    name: 'Dashboard',
    path: '',
    component: async () => import('@/views/Dashboard.vue'),
    meta: {
      authRequired: true,
    },
  },
]
