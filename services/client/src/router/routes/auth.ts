export default [
  {
    name: 'Login',
    path: '/login',
    component: async () => import('@/views/Login.vue'),
  },
  {
    name: 'Register',
    path: '/register',
    component: async () => import('@/views/Register.vue'),
  },
]
