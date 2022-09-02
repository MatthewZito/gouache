import { GouacheError } from '.'

export class ErroneousResponseError extends GouacheError {
  public readonly friendly: string

  constructor(friendly: string) {
    super({ internal: 'Http request failure', friendly })
    this.friendly = friendly

    Object.setPrototypeOf(this, ErroneousResponseError.prototype)
  }

  serialize() {
    return this.messages
  }
}
