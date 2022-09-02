import { GouacheError } from '.'

export class InsufficientDataError extends GouacheError {
  constructor({ field, friendly }: { field: string; friendly: string }) {
    super({ internal: `Missing data for field: ${field}`, friendly })

    Object.setPrototypeOf(this, InsufficientDataError.prototype)
  }

  serialize() {
    return this.messages
  }
}
