<script setup lang="ts">
import { useQuasar } from 'quasar'

import NavigationDrawer from '@/components/Layout/NavigationDrawer.vue'
import UserProfile from '@/components/Layout/UserProfile.vue'
import { useErrorHandler } from '@/services'
import { useSessionStore } from '@/state'

const $q = useQuasar()
const sessionStore = useSessionStore()

const leftDrawerOpen = ref(false)

const isDarkMode = computed(() => $q.dark.isActive)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

function toggleDarkMode() {
  $q.dark.toggle()
}

onErrorCaptured((ex: any) => {
  useErrorHandler(ex, { notify: true })

  return false
})
</script>

<template lang="pug">
q-layout(view="hHh LpR fFf")
  q-header.bg-primary.text-white(bordered)
    q-toolbar
      q-btn(
        dense
        flat
        round
        icon="mdi-menu"
        @click="toggleLeftDrawer"
      )

      q-toolbar-title
        | Resources

      q-btn.q-mr-md(
        dense
        flat
        round
        :icon="isDarkMode ? `mdi-moon-waxing-crescent` : `mdi-white-balance-sunny`"
        @click="toggleDarkMode"
      )

      q-btn(dense flat round icon="mdi-account")
        UserProfile(
          :username="sessionStore.username"
          :exp="sessionStore.exp"
        )

  NavigationDrawer(v-model="leftDrawerOpen")

  q-page-container.q-pa-lg
    router-view
</template>
