export type UUID = string

export type UNIXTimestamp = string

export type GouacheValidationRule<T> = (
  value: T,
) => Promise<boolean | string> | boolean | string
