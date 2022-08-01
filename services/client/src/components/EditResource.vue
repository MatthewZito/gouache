<script lang="ts" setup>
import type { PropType } from 'vue'
import type { MutableResource } from '@/types'
import { listRequired, required } from '@/utils'
import { showNotification } from '@/plugins'
import {
  ErroneousResponseError,
  resourceApi,
  useErrorHandler,
} from '@/services'
import { availableTags } from '@/mock'
import { useResourceStore } from '@/state'
import { UUID } from '@/types/scalar'
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

const isLoading = ref(false)
const formModel = reactive<MutableResource>({
  title: data.title,
  tags: data.tags,
})

const shouldDisable = computed(
  () =>
    !Object.values(formModel).every(
      value =>
        (value != null && typeof value === 'string' && value !== '') ||
        (value?.length && value.length > 0),
    ),
)

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
</script>

<template>
  <q-card style="width: 400px">
    <q-card-section>
      <div class="text-h6">Edit Resource</div>
    </q-card-section>

    <q-card-section>
      <q-form class="q-pa-md">
        <q-input
          label="Title"
          v-model="formModel.title"
          filled
          dense
          class="q-mb-md"
          :rules="[required('A title is required.')]"
        />
        <GSelect
          v-model="formModel.tags"
          :options="availableTags"
          label="Tags"
          :rules="[listRequired('At least one tag is required.')]"
        />
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
      <div class="text-secondary text-bold">Updating the resource...</div>
    </q-inner-loading>
  </q-card>
</template>
