interface MessageFields {
  internal: string
  friendly: string
}
export abstract class GouacheError extends Error {
  // eslint-disable-next-line @typescript-eslint/parameter-properties
  constructor(public messages: MessageFields) {
    super(messages.internal)

    // preserve the prototype chain in tsc generated js
    Object.setPrototypeOf(this, GouacheError.prototype)
  }

  abstract serialize(): {
    internal: string
    friendly: string
  }
}
