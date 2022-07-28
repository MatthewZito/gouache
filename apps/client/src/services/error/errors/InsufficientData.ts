import { BaseError } from '.';

export class InsufficientDataError extends BaseError {
	constructor(protected field: string, public friendly: string) {
		super({ internal: `Missing data for field: ${field}`, friendly });

		Object.setPrototypeOf(this, InsufficientDataError.prototype);
	}

	serialize() {
		return this.messages;
	}
}
