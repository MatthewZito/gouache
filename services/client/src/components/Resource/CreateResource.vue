<script lang="ts" setup>
import GSelect from '@/components/ui/GSelect.vue'
import { useMutateResource } from '@/hooks'
import { availableTags } from '@/mock'
import { showNotification } from '@/plugins'
import { InsufficientDataError, resourceApi, useErrorHandler } from '@/services'
import { useResourceStore } from '@/state'
import type { AllNullable, MutableResource } from '@/types'

const E_CANT_CREATE =
  'Something went wrong while creating this resource. Please try again or contact support.'

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const resourceStore = useResourceStore()

const { tagsRules, titleRules, formModel, shouldDisable } = useMutateResource({
  title: '',
  tags: [],
})
const isLoading = ref(false)

function validateFormModel(
  model: AllNullable<MutableResource>,
): asserts model is MutableResource {
  if (!model.title) {
    throw new InsufficientDataError({
      field: 'title',
      friendly: E_CANT_CREATE,
    })
  }
}

async function handleSave() {
  isLoading.value = true

  try {
    validateFormModel(formModel)
    const { ok, data: id } = await resourceApi.createResource(formModel)

    if (!ok) {
      throw Error('failed to create resource')
    }

    showNotification('success', 'Successfully created new resource.')

    const now = new Date().toISOString()
    resourceStore.prependResource({
      ...formModel,
      id,
      createdAt: now,
      updatedAt: now,
    })

    $emit('close')
  } catch (ex) {
    useErrorHandler(ex, { notify: true, fallback: E_CANT_CREATE })
  } finally {
    isLoading.value = false
  }
}
</script>

<template lang="pug">
q-card(
  style="width: 400px"
)
  q-card-section
    div.text-h6
      | Create Resource

  q-card-section
    q-form.q-pa-sm
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

  q-card-actions.q-pa-none.q-pt-sm.justify-between
    q-btn(
      label="Close"
      flat
      color="grey-6"
      @click="$emit('close')"
    )
    div
      q-btn(
        label="Save"
        unelevated
        color="primary"
        :disable="shouldDisable"
        @click="handleSave"
      )
      q-tooltip(
        v-if="shouldDisable"
      )
        | The form must be complete prior to submitting.

  q-inner-loading(:showing="isLoading")
    q-spinner.q-mb-sm(
      size="50px"
      color="secondary"
      )
    div.text-secondary.text-bold
      | Creating the resource...
</template>
