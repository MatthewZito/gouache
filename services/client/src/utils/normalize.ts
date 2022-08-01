export function normalizeNullish(maybeNullish?: any) {
  return maybeNullish == null ? 'N/A' : maybeNullish
}
