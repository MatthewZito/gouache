export default [
  {
    name: 'Dashboard',
    path: '/dashboard',
    component: async () => import('@/views/Dashboard/Dashboard.vue'),
    meta: {
      authRequired: true,
    },
  },
  {
    name: 'Resource',
    path: '/resources',
    component: async () => import('@/views/Resources/ResourceWrapper.vue'),
    meta: {
      authRequired: true,
    },
  },
  {
    name: 'Reporting',
    path: '/reporting',
    component: async () => import('@/views/Reporting/ReportingWrapper.vue'),
    meta: {
      authRequired: true,
    },
  },
]
