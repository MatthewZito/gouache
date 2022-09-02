<script lang="ts" setup>
import { Loading } from 'quasar'

import GPasswordInput from '@/components/ui/GPasswordInput.vue'
import { useCredentials } from '@/hooks'
import { PASSWORD_MAX_CHARS, PASSWORD_MIN_CHARS } from '@/meta'
import { authApi, ErroneousResponseError, useErrorHandler } from '@/services'
import { useSessionStore } from '@/state'
import { maxLength, minLength, required } from '@/utils'

import GLogo from '../components/ui/GLogo.vue'

const $router = useRouter()
const sessionStore = useSessionStore()
const { formModel, shouldDisable } = useCredentials([
  ({ password }) => password.length >= PASSWORD_MIN_CHARS,
  ({ password }) => password.length <= PASSWORD_MAX_CHARS,
])

const passwordRules = [
  required('A password is required.'),
  minLength(
    `The supplied password must be greater than ${
      PASSWORD_MIN_CHARS - 1
    } characters in length.`,
    PASSWORD_MIN_CHARS,
  ),
  maxLength(
    `The supplied password must be less than ${
      PASSWORD_MAX_CHARS + 1
    } characters in length.`,
    PASSWORD_MAX_CHARS,
  ),
]

async function handleSubmitRegister() {
  Loading.show()
  try {
    const { ok, data } = await authApi.register(formModel)

    if (!ok || !data) {
      throw new ErroneousResponseError(
        'Something went wrong while registering this username.',
      )
    }

    sessionStore.setUserState(data)
    $router.push({ name: 'Dashboard' })
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback: 'Something went wrong while registering.',
    })
  } finally {
    Loading.hide()
  }
}

// @todo
const handleError = console.error
</script>

<template lang="pug">
div.row.justify-center.items-center.full-height
  q-card(style="width: 400px")
    q-form.q-pa-md(
      @submit.prevent
      @validation-error="handleError"
      greedy
    )
      q-card-section
        GLogo
      q-card-section
        div.text-h6
          | Register a new account

      q-card-section
        q-input.q-mb-md(
          label="Username"
          v-model="formModel.username"
          filled
          dense
          autocomplete="username"
          :rules="[required('A username is required.')]"
        )
        GPasswordInput(
          v-model="formModel.password"
          :rules="passwordRules"
        )

        q-card-actions.justify-between
          q-btn(
            label="Login"
            unelevated
            color="primary"
            outline
            @click="$router.push({ name: 'Login' })"
          )
          div
            q-btn(
              label="Register"
              type="submit"
              unelevated
              color="primary"
              :disable="shouldDisable"
              @click="handleSubmitRegister"
            )
            q-tooltip(v-if="shouldDisable")
              | You must provide a valid username and password to register.

</template>
