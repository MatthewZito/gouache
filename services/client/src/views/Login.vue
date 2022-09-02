<script lang="ts" setup>
import { Loading } from 'quasar'

import GLogo from '@/components/ui/GLogo.vue'
import GPasswordInput from '@/components/ui/GPasswordInput.vue'
import { useCredentials } from '@/hooks'
import { authApi, ErroneousResponseError, useErrorHandler } from '@/services'
import { useSessionStore } from '@/state'
import { required } from '@/utils'

const $router = useRouter()
const sessionStore = useSessionStore()
const { formModel, shouldDisable } = useCredentials()

async function handleSubmitLogin() {
  Loading.show()
  try {
    const { ok, data } = await authApi.login(formModel)

    if (!ok) {
      throw new ErroneousResponseError('Invalid credentials.')
    }

    sessionStore.setUserState(data)
    $router.push({ name: 'Dashboard' })
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback: 'Unable to login.',
    })
  } finally {
    Loading.hide()
  }
}
</script>

<template lang="pug">
q-card(style="width: 400px")
  q-form.q-pa-sm(@submit.prevent)
    q-card-section
      GLogo

    q-card-section
      div.text-h6
      | Login with an existing account

    q-card-section
      q-input.q-mb-md(
        label="Username"
        v-model="formModel.username"
        filled
        dense
        autocomplete="username"
        :rules="[required('A username is required.')]"
      )

      GPasswordInput(v-model="formModel.password")

    q-card-actions.q-pa-none.q-pt-sm.justify-between
      q-btn(
        label="Register"
        unelevated
        color="primary"
        outline
        @click="$router.push({ name: 'Register' })"
      )

      div
        q-btn(
          label="Login"
          type="submit"
          unelevated
          color="primary"
          :disable="shouldDisable"
          @click="handleSubmitLogin"
        )

        q-tooltip(v-if="shouldDisable")
          | You must provide a username and password to login.
</template>
