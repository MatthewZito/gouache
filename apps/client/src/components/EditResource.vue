<script lang="ts" setup>
import type { PropType } from 'vue'
import type { Resource } from '@/types'
import { required } from '@/utils'
import { showNotification } from '@/plugins'
import { logger, resourceApi } from '@/services'

const props = defineProps({
  resource: {
    type: Object as PropType<Resource>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'close', refresh: boolean): void
}>()

const isLoading = ref(false)
const formModel = reactive<Omit<Resource, 'id' | 'key'>>({
  value: props.resource.value,
  expires: props.resource.expires,
})

const shouldDisable = computed(
  () => !Object.values(formModel).every(value => value != null && value !== ''),
)

async function handleSave() {
  isLoading.value = true

  try {
    const { ok } = await resourceApi.updateResource(
      props.resource.key,
      formModel,
    )

    if (!ok) {
      throw Error('failed to update resource')
    }

    showNotification('success', 'Successfully updated new resource.')

    $emit('close', true)
  } catch (ex) {
    logger.error('failed to update resource', ex)
    showNotification('error', 'Failed to update resource.')
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <q-card style="width: 400px">
    <q-card-section>
      <div class="text-h6">Edit Resource</div>
    </q-card-section>

    <q-card-section>
      <q-form class="q-pa-md">
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
      <div class="text-secondary text-bold">Updating the resource...</div>
    </q-inner-loading>
  </q-card>
</template>
