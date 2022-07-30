import { BaseError } from '.'

export class UserActionException extends BaseError {
  constructor({
    friendly,
    internal = 'View model exception',
  }: {
    friendly: string
    internal?: string
  }) {
    super({ internal, friendly })

    Object.setPrototypeOf(this, UserActionException.prototype)
  }

  serialize() {
    return this.messages
  }
}
