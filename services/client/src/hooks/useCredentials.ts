interface Constraint {
  (v: Credentials): boolean
}

interface Credentials {
  username: string
  password: string
}

/**
 * Create pre-packaged business logic and state for managing user-input credentials.
 *
 * @param withConstraints Credentials validation constraints to be applied during input validation.
 * Each constraint must be a predicate function that accepts as input the credentials and returns a boolean
 * indicating whether they are valid.
 */
export function useCredentials(withConstraints?: Constraint[]) {
  const shouldDisable = computed(() => {
    const areCredsIncomplete = !formModel.username || !formModel.password

    if (!areCredsIncomplete && withConstraints?.length) {
      return withConstraints.some(constraint => !constraint(formModel))
    }

    return areCredsIncomplete
  })

  const formModel = reactive<Credentials>({
    username: '',
    password: '',
  })

  return {
    formModel,
    shouldDisable,
  }
}
