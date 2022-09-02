import { createPinia } from 'pinia'
import { Dialog, Loading, Notify, Quasar } from 'quasar'
import quasarIconSet from 'quasar/icon-set/svg-mdi-v6'
import { createApp } from 'vue'

import { debugPlugin, registerNotifyPlugin } from '@/plugins'
import router from '@/router'
import { useSessionStore } from '@/state'

import App from './App.vue'
import '@quasar/extras/mdi-v6/mdi-v6.css'
import 'quasar/src/css/index.sass'
import { quasarStyles } from './styles/quasar'

const vm = createApp(App)
  .use(debugPlugin)
  .use(registerNotifyPlugin)
  .use(createPinia())
  .use(Quasar, {
    config: {
      ...quasarStyles,
      loading: {},
      notify: {},
    },
    iconSet: quasarIconSet,
    plugins: [Dialog, Loading, Notify],
  })

useSessionStore()
  .verifySession()
  .catch(async () => router.replace({ name: 'Login' }))
  .finally(() => {
    vm.use(router)
    vm.mount('#app')
  })
