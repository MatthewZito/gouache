<script lang="ts" setup>
import { useCredentials } from '@/hooks'
import { authApi, ErroneousResponseError, useErrorHandler } from '@/services'
import { maxLength, minLength, required } from '@/utils'
import { Loading } from 'quasar'
import { useSessionStore } from '@/state'
import GPasswordInput from '@/components/ui/GPasswordInput.vue'
import GLogo from '../components/ui/GLogo.vue'
import { PASSWORD_MAX_CHARS, PASSWORD_MIN_CHARS } from '@/meta'

const $router = useRouter()
const sessionStore = useSessionStore()
const { formModel, shouldDisable } = useCredentials([
  ({ password }) => password.length >= PASSWORD_MIN_CHARS,
  ({ password }) => password.length <= PASSWORD_MAX_CHARS,
])

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

<template>
  <div class="row justify-center items-center full-height">
    <q-card style="width: 400px">
      <q-form
        class="q-pa-md"
        @submit.prevent
        @validation-error="handleError"
        greedy
      >
        <q-card-section>
          <GLogo />
        </q-card-section>
        <q-card-section>
          <div class="text-h6">Register a new account</div>
        </q-card-section>
        <q-card-section>
          <q-input
            label="Username"
            v-model="formModel.username"
            filled
            dense
            class="q-mb-md"
            autocomplete="username"
            :rules="[required('A username is required.')]"
          />
          <GPasswordInput
            v-model="formModel.password"
            :rules="[
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
            ]"
          />
        </q-card-section>

        <q-card-actions class="justify-between">
          <q-btn
            label="Login"
            unelevated
            color="primary"
            outline
            @click="$router.push({ name: 'Login' })"
          />
          <div>
            <q-btn
              label="Register"
              type="submit"
              unelevated
              color="primary"
              :disable="shouldDisable"
              @click="handleSubmitRegister"
            />
            <q-tooltip v-if="shouldDisable">
              You must provide a valid username and password to register.
            </q-tooltip>
          </div>
        </q-card-actions>
      </q-form>
    </q-card>
  </div>
</template>
