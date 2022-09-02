<script lang="ts" setup>
import { useSessionStore } from '@/state'
import { nowPlusNSeconds } from '@/utils'

const props = defineProps({
  username: {
    type: String,
    required: true,
  },
  exp: {
    type: Number,
    required: true,
  },
})

const sessionStore = useSessionStore()
</script>

<template>
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
              {{ props.username }}
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
            Next session renewal at
            <span class="text-weight-bold">
              {{ nowPlusNSeconds(props.exp) }}
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
</template>
