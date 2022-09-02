export type UUID = string

export type UNIXTimestamp = string

export type GouacheValidationRule<T> = (
  value: T,
) => boolean | string | Promise<boolean | string>
