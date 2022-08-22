import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'

import router from '@/router'
import { useSessionStore } from '@/state'
import { debugPlugin, registerNotifyPlugin } from '@/plugins'
import { Dialog, Loading, Notify, Quasar } from 'quasar'
import quasarIconSet from 'quasar/icon-set/svg-mdi-v6'
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
  .then(() => {
    return router.replace({ name: 'Dashboard' })
  })
  .finally(() => {
    vm.use(router)
    vm.mount('#app')
  })
