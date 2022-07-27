import { createApp } from 'vue'

import App from './App.vue'
import './styles'

import router from '@/router'
import { debugPlugin } from '@/plugins'

createApp(App).use(debugPlugin).use(router).mount('#app')
