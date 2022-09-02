import { RESOURCE_TITLE_MAX_CHARS } from '@/meta'
import type { MutableResource, Resource } from '@/types'
import { generateValidator, listRequired, maxLength, required } from '@/utils'

/**
 * Create pre-packaged business logic and state for managing user-input Resource data.
 *
 * @param data The initial form model data.
 */
export function useMutateResource(data: Pick<Resource, 'tags' | 'title'>) {
  /**
   * Rules for validating the Resource tags.
   */
  const tagsRules = [listRequired('At least one tag is required.')]

  /**
   * Rules for validating the Resource title.
   */
  const titleRules = [
    required('A title is required.'),
    maxLength(
      `The supplied title must be less than ${
        RESOURCE_TITLE_MAX_CHARS + 1
      } characters in length.`,
      RESOURCE_TITLE_MAX_CHARS,
    ),
  ]

  /**
   * A validator function for use with the form model title.
   *
   * @internal
   */
  const validateTitle = generateValidator(titleRules)

  /**
   * The form model.
   */
  const formModel = reactive<MutableResource>({
    title: data.title,
    tags: data.tags,
  })

  /**
   * A computed flag indicating whether or not the submit/save button should be disabled, qua the input validations.
   */
  const shouldDisable = computed(
    () => !validateTitle(formModel.title) || formModel.tags.length <= 0,
  )

  return {
    tagsRules,
    titleRules,
    formModel,
    shouldDisable,
  }
}
