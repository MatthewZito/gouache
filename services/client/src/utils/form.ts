/*
Form validation rules. Compatible with vee-validate & Quasar.
*/

/**
 * Create a required value validation rule.
 * @param message A message to display when invalid.
 */
export function required(message: string) {
  return function validate(val: string | null) {
    return (val != null && val !== '') || message
  }
}
/**
 * Create a required list value validation rule.
 * @param message A message to display when invalid.
 */
export function listRequired(message: string) {
  return function validate(val: string[] | null) {
    return (val != null && val.length !== 0) || message
  }
}

/**
 * Create a minimum length validation rule.
 * @param message A message to display when invalid.
 * @param min The minimum value constraint.
 */
export function minLength(message: string, min: number) {
  return function validate(val: string | null) {
    return (val != null && val.length >= min) || message
  }
}

/**
 * Create a maximum length validation rule.
 * @param message A message to display when invalid.
 * @param max The maximum value constraint.
 */
export function maxLength(message: string, max: number) {
  return function validate(val: string | null) {
    return (val != null && val.length <= max) || message
  }
}

/**
 * Passthrough rule decorator - only enforces the given rule if a value was actually provided.
 * Use this for inputs where the value is optional but a rule is enforced if and when a value is provided.
 *
 * @param rule The return value of any `rule` creator function, as above.
 */
export function optionalWithRule(rule: (val: string | null) => string | true) {
  return function validate(val: string | null) {
    return val !== '' && val !== null ? rule(val) : true
  }
}
