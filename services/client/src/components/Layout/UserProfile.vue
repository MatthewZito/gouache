<script lang="ts" setup>
import { useSessionStore } from '@/state'
import { nowPlusNSeconds } from '@/utils'

const props = defineProps<{
  username: string | null
  exp: number | null
}>()

const sessionStore = useSessionStore()
</script>

<template lang="pug">
q-menu.q-pa-xs
  q-list
    q-item
      q-item-section(avatar)
        q-icon(
          name="mdi-account"
          size="md"
          color="secondary"
        )
      q-item-section.text-subtitle1
        q-item-label
          | Logged in as
          span.text-weight-bold
            | {{ props.username }}

    q-item
      q-item-section(avatar)
        q-icon(
          name="mdi-clock"
          size="md"
          color="secondary"
        )

      q-item-section.text-subtitle1
        q-item-label
          | Next session renewal at
          span.text-weight-bold
            | {{ nowPlusNSeconds(props.exp) }}

    q-item(
      clickable
      @click="sessionStore.logout"
    )
      q-item-section(avatar)
        q-icon(
          name="mdi-logout"
          size="md"
          color="secondary"
        )

      q-item-section.text-subtitle1
        q-item-label
          | Logout
</template>
