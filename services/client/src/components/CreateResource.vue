<script lang="ts" setup>
import { showNotification } from '@/plugins'
import { InsufficientDataError, resourceApi, useErrorHandler } from '@/services'
import type { AllNullable, MutableResource } from '@/types'
import { required, listRequired } from '@/utils'
import { availableTags } from '@/mock'
import { useResourceStore } from '@/state'

const E_CANT_CREATE =
  'Something went wrong while creating this resource. Please try again or contact support.'

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const resourceStore = useResourceStore()

const isLoading = ref(false)
const formModel = reactive<MutableResource>({
  title: '',
  tags: [],
})

const shouldDisable = computed(
  () =>
    !Object.values(formModel).every(
      value =>
        (value != null && typeof value === 'string' && value !== '') ||
        (value?.length && value.length > 0),
    ),
)

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

function removeOption(value: string) {
  const foundIdx = formModel.tags.findIndex(tag => tag === value)
  if (foundIdx !== -1) {
    formModel.tags.splice(foundIdx, 1)
  }
}
</script>

<template>
  <q-card style="width: 400px">
    <q-card-section>
      <div class="text-h6">Create Resource</div>
    </q-card-section>

    <q-card-section>
      <!-- @todo reuse -->
      <q-form class="q-pa-md">
        <q-input
          label="Title"
          v-model="formModel.title"
          filled
          dense
          class="q-mb-md"
          :rules="[required('A title is required.')]"
        />
        <q-select
          v-model="formModel.tags"
          :options="availableTags"
          label="Tags"
          option-value="value"
          class="mr-2 extra-dense"
          filled
          clearable
          dense
          emit-value
          map-options
          multiple
          use-chips
          :rules="[listRequired('At least one tag is required.')]"
        >
          <template #selected-item="{ opt }">
            <q-chip
              :color="opt.color"
              removable
              @remove="_ => removeOption(opt.value)"
            >
              <span class="q-mr-sm">
                {{ opt.label }}
              </span>
              <q-icon :name="opt.icon" />
              <q-tooltip>
                {{ opt.description }}
              </q-tooltip>
            </q-chip>
          </template>
        </q-select>
      </q-form>
    </q-card-section>

    <q-card-actions class="justify-between">
      <q-btn label="Close" flat color="grey-6" @click="$emit('close')" />
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
