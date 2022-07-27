<script lang="ts" setup>
import styled from '@magister_zito/vue3-styled-components'
import { PropType } from 'vue'
import Field from '../ui/Field.vue'

const StyledForm = styled.form`
  display: flex;
  flex-direction: column;
`

const ModalContentContainer = styled.div`
  padding: 1rem;
  min-height: 8rem;
  display: flex;
  justify-content: center;
`

const StyledFooter = styled.footer`
  display: flex;
  justify-content: space-between;
  padding: 0.75rem;
  border-top: 1px solid lightgray;
`

const HeaderText = styled.h1`
  font-size: 1rem;
  color: white;
`

const StyledHeader = styled.header`
  position: relative;
  display: flex;
  flex: none;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  border-radius: 0.5rem 0.5rem 0 0;
  background-color: green;
`

interface Resource {
  key: string
  value: string
  expires: number
}

const props = defineProps({
  selectedRow: {
    type: Object as PropType<Resource>,
    required: true,
  },
})

const $emit = defineEmits<{
  (e: 'close'): void
}>()

const formModel = reactive({
  key: props.selectedRow.key,
  value: props.selectedRow.value,
  expires: props.selectedRow.expires,
})

const isSubmitDisabled = computed(() => {
  const formValues = Object.values(formModel)
  const nRequiredValues = formValues.length

  return (
    formValues.filter(value => value != null && value !== '').length !==
    nRequiredValues
  )
})

function handleSaveEdit() {
  console.log({ formModel })
}
</script>

<template>
  <StyledHeader>
    <HeaderText> Edit Resource </HeaderText>
    <button @click="$emit('close')">Close</button>
  </StyledHeader>
  <ModalContentContainer>
    <StyledForm>
      <Field label="Key" v-model="formModel.key" maxlength="10" />
      <Field label="Value" v-model="formModel.value" maxlength="10" />
      <Field label="Expires" v-model="formModel.expires" maxlength="10" />
    </StyledForm>
  </ModalContentContainer>

  <StyledFooter>
    <button @click="$emit('close')">Close</button>
    <button @click="handleSaveEdit" :disabled="isSubmitDisabled">Save</button>
  </StyledFooter>
</template>
