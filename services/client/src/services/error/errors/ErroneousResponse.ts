import { GouacheError } from '.'

export class ErroneousResponseError extends GouacheError {
  constructor(public friendly: string) {
    super({ internal: 'Http request failure', friendly })

    Object.setPrototypeOf(this, ErroneousResponseError.prototype)
  }

  serialize() {
    return this.messages
  }
}
