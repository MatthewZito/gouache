<script setup lang="ts">
import { useErrorHandler } from '@/services'
import { useSessionStore } from '@/state'
import { epochToReadableTime, normalizeNullish } from '@/utils'
import { useQuasar } from 'quasar'

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
          class="q-mr-md"
        />

        <q-btn dense flat round icon="mdi-account">
          <q-menu class="q-pa-xs">
            <q-list>
              <q-item>
                <q-item-section avatar>
                  <q-icon name="mdi-account" size="md" color="secondary" />
                </q-item-section>

                <q-item-section class="text-subtitle1">
                  <q-item-label>
                    Logged in as
                    <span class="text-weight-bold">
                      {{ sessionStore.username }}
                    </span>
                  </q-item-label>
                </q-item-section>
              </q-item>

              <q-item>
                <q-item-section avatar>
                  <q-icon name="mdi-clock" size="md" color="secondary" />
                </q-item-section>

                <q-item-section class="text-subtitle1">
                  <q-item-label>
                    Next session renewal
                    <span class="text-weight-bold">
                      {{
                        normalizeNullish(epochToReadableTime(sessionStore.exp))
                      }}
                    </span>
                  </q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable @click="sessionStore.logout">
                <q-item-section avatar>
                  <q-icon name="mdi-logout" size="md" color="secondary" />
                </q-item-section>
                <q-item-section class="text-subtitle1">
                  <q-item-label>Logout</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </q-btn>
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
