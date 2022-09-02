<script setup lang="ts">
import { useQuasar } from 'quasar'

import { useErrorHandler } from '@/services'
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

<template lang="pug">
q-layout(view="hHh LpR fFf")
  q-header.bg-primary.text-white(bordered)
    q-toolbar
      q-toolbar-title
        |Resources

      q-btn(
        dense
        flat
        round
        :icon="isDarkMode ? `mdi-moon-waxing-crescent` : `mdi-white-balance-sunny`"
        @click="toggleDarkMode"
      )

  q-page-container.q-pa-lg(style="height: 100vh")
    div.row.justify-center.items-center.full-height
      router-view
</template>
