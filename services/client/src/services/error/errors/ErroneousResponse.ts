import { BaseError } from '.';

export class ErroneousResponseError extends BaseError {
	constructor(public friendly: string) {
		super({ internal: 'Http request failure', friendly });

		Object.setPrototypeOf(this, ErroneousResponseError.prototype);
	}

	serialize() {
		return this.messages;
	}
}
