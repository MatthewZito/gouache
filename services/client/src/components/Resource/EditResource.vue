<script lang="ts" setup>
import type { PropType } from 'vue'

import type { UUID } from '@/types/scalar'

import { showNotification } from '@/plugins'
import {
  ErroneousResponseError,
  resourceApi,
  useErrorHandler,
} from '@/services'
import { availableTags } from '@/mock'
import { useResourceStore } from '@/state'
import { useMutateResource } from '@/hooks'

import GSelect from '@/components/ui/GSelect.vue'

const props = defineProps({
  resourceId: {
    type: String as PropType<UUID>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const { ok, data } = await resourceApi.getResource(props.resourceId)

if (!ok) {
  throw new ErroneousResponseError(
    'Something went wrong while loading the details for this resource.',
  )
}

const resourceStore = useResourceStore()

const { tagsRules, titleRules, formModel, shouldDisable } =
  useMutateResource(data)
const isLoading = ref(false)

async function handleSave() {
  isLoading.value = true

  try {
    if (!data) {
      throw Error('The resource failed to load. This is likely a bug.')
    }

    const { ok } = await resourceApi.updateResource(props.resourceId, formModel)

    if (!ok) {
      throw Error('failed to update resource')
    }

    showNotification('success', 'Successfully updated resource.')

    resourceStore.patchResources({ ...data, ...formModel })
    $emit('close')
  } catch (ex) {
    useErrorHandler(ex, {
      notify: true,
      fallback: 'Something went wrong while updating this resource.',
    })
  } finally {
    isLoading.value = false
  }
}

// @todo
const handleError = console.error
</script>

<template lang="pug">
q-card(
  style="width: 400px"
)
  q-form.q-pa-md(
    @submit.prevent
    @validation-error="handleError"
    greedy
  )
    q-card-section
      div.text-h6
        | Edit Resource

    q-card-section
      q-input.q-mb-md(
        label="Title"
        v-model="formModel.title"
        filled
        dense
        :rules="titleRules"
      )
      GSelect(
        v-model="formModel.tags"
        :options="availableTags"
        label="Tags"
        :rules="tagsRules"
      )

    q-card-actions.justify-between
      q-btn(
        label="Close"
        flat
        color="grey-6"
        @click="$emit('close')"
      )
      div
        q-btn(
          label="Save"
          type="submit"
          unelevated
          color="primary"
          :disable="shouldDisable"
          @click="handleSave"
        )
        q-tooltip(v-if="shouldDisable")
          | The form must be complete prior to submitting.

    q-inner-loading(:showing="isLoading")
      q-spinner.q-mb-sm(
        size="50px"
        color="secondary"
      )
      div.text-secondary.text-bold
        | Updating the resource...
</template>
