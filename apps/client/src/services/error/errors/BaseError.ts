export abstract class BaseError extends Error {
  constructor(
    public messages: {
      internal: string
      friendly: string
    },
  ) {
    super(messages.internal)

    // preserve the prototype chain in tsc generated js
    Object.setPrototypeOf(this, BaseError.prototype)
  }

  abstract serialize(): {
    internal: string
    friendly: string
  }
}
