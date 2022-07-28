export default [
  {
    name: 'Dashboard',
    path: '',
    component: async () => import('@/views/DashboardWrapper.vue'),
    meta: {
      authRequired: true,
    },
  },
]
