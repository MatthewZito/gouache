<script lang="ts" setup>
import { showNotification } from '@/plugins'
import { InsufficientDataError, logger, resourceApi } from '@/services'
import type { AllNullable, Resource } from '@/types'
import { required } from '@/utils'

const $emit = defineEmits<{
  (e: 'close', refresh: boolean): void
}>()

const isLoading = ref(false)
const formModel = reactive<AllNullable<Omit<Resource, 'id'>>>({
  key: null,
  value: null,
  expires: null,
})

const shouldDisable = computed(
  () => !Object.values(formModel).every(value => value != null && value !== ''),
)

function validateFormModel(
  model: AllNullable<Omit<Resource, 'id'>>,
): asserts model is Resource {
  if (!model.expires) {
    throw new InsufficientDataError('expires', 'Something went wrong.')
  }

  if (!model.key) {
    throw new InsufficientDataError('key', 'Something went wrong.')
  }

  if (!model.value) {
    throw new InsufficientDataError('value', 'Something went wrong.')
  }
}

async function handleSave() {
  isLoading.value = true

  try {
    validateFormModel(formModel)
    const { ok } = await resourceApi.createResource(formModel)

    if (!ok) {
      throw Error('failed to create resource')
    }

    showNotification('success', 'Successfully created new resource.')

    $emit('close', true)
  } catch (ex) {
    logger.error('failed to create resource', ex)
    showNotification('error', 'Failed to create resource.')
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <q-card style="width: 400px">
    <q-card-section>
      <div class="text-h6">Create Resource</div>
    </q-card-section>

    <q-card-section>
      <q-form class="q-pa-md">
        <q-input
          label="Key"
          v-model="formModel.key"
          filled
          dense
          class="q-mb-md"
          :rules="[required('A key is required.')]"
        />
        <q-input
          label="Value"
          v-model="formModel.value"
          filled
          dense
          class="q-mb-md"
          :rules="[required('A value is required.')]"
        />
        <q-input
          label="Expires"
          v-model="formModel.expires"
          filled
          dense
          class="q-mb-md"
          :rules="[required('An expiration is required.')]"
        />
      </q-form>
    </q-card-section>

    <q-card-actions class="justify-between">
      <q-btn label="Close" flat color="grey-6" @click="$emit('close', false)" />
      <div>
        <q-btn
          label="Save"
          unelevated
          color="primary"
          :disable="shouldDisable"
          @click="handleSave"
        />
        <q-tooltip v-if="shouldDisable">
          The form must be complete prior to submitting.
        </q-tooltip>
      </div>
    </q-card-actions>

    <q-inner-loading :showing="isLoading">
      <q-spinner class="q-mb-sm" size="50px" color="secondary" />
      <div class="text-secondary text-bold">Creating the resource...</div>
    </q-inner-loading>
  </q-card>
</template>
