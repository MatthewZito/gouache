/*
Form validation rules. Compatible with vee-validate & Quasar.
*/

import { GouacheValidationRule } from '@/types/scalar'

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
export function optionalWithRule<T = string | null>(
  rule: GouacheValidationRule<T>,
) {
  return function validate(val: T) {
    return val !== null ? rule(val) : true
  }
}

/**
 * Generate a validator function given a list of rules. The returned validator function `validate`
 * evaluates a given input against each of the rules in the order in which they were passed to the generator,
 * and returns a boolean indicating whether any of the rules were violated.
 *
 * Each rule is invoked lazily - if a rule is violated, the validator immediately returns.
 * @param rules A list of rules to validate against.
 * @returns A validator function.
 */
export function generateValidator<T = string | null>(
  rules: GouacheValidationRule<T>[],
) {
  function validate(val: T) {
    return rules.reduce<boolean>((acc, rule) => {
      if (!acc) {
        return false
      }

      return rule(val) === true
    }, true)
  }

  return validate
}
