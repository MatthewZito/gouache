export type AllNullable<T extends Record<any, any>> = {
  [key in keyof T]: T[key] | null
}
