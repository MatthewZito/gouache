<script setup lang="ts">
import { useErrorHandler } from '@/services'
import { useQuasar } from 'quasar'
const $q = useQuasar()

const isDarkMode = computed(() => $q.dark.isActive)

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

    <q-page-container class="q-pa-lg" style="height: 100vh">
      <div class="row justify-center items-center full-height">
        <router-view />
      </div>
    </q-page-container>
  </q-layout>
</template>
