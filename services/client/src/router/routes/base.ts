export default [
  {
    name: 'Dashboard',
    path: '',
    component: async () => import('@/views/Dashboard/DashboardWrapper.vue'),
    meta: {
      authRequired: true,
    },
  },
]
