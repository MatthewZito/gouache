/**
 * Format a date or date string as MM/DD/YYYY.
 * Returns `null` if the provided raw date / string is falsy or otherwise would result in an Invalid Date result.
 */
export function toReadableDate(date?: Date | string) {
  if (!date) {
    return null
  }

  const d = new Date(date)

  return isNaN(d as unknown as number)
    ? null
    : d.toLocaleDateString('en-US', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
      })
}
