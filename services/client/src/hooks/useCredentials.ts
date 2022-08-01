interface Constraint {
  (v: Credentials): boolean
}

interface Credentials {
  username: string
  password: string
}

export function useCredentials(withConstraints?: Constraint[]) {
  const shouldDisable = computed(() => {
    const areCredsIncomplete = !formModel.username || !formModel.password

    if (!areCredsIncomplete && withConstraints?.length) {
      return withConstraints.some(constraint => {
        return !constraint(formModel)
      })
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
