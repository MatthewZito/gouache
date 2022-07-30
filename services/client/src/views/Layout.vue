<script setup lang="ts">
import { useErrorHandler } from '@/services'
import { useQuasar } from 'quasar'
const $q = useQuasar()

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

<template>
  <q-layout view="hHh LpR fFf">
    <q-header bordered class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="mdi-menu" @click="toggleLeftDrawer" />

        <q-toolbar-title> Resources </q-toolbar-title>

        <q-btn
          dense
          flat
          round
          :icon="
            isDarkMode ? `mdi-moon-waxing-crescent` : `mdi-white-balance-sunny`
          "
          @click="toggleDarkMode"
        />
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container class="q-pa-lg">
      <router-view />
    </q-page-container>
  </q-layout>
</template>
