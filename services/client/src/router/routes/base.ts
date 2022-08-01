export default [
  {
    name: 'Dashboard',
    path: '/dashboard',
    component: async () => import('@/views/Dashboard/DashboardWrapper.vue'),
    meta: {
      authRequired: true,
    },
  },
]
