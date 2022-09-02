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

export function epochToReadableTime(maybeEpoch: number | null) {
  if (!maybeEpoch) {
    return null
  }

  const d = new Date(maybeEpoch * 1000)

  if (isNaN(d.valueOf())) {
    return null
  }

  return d.toLocaleString('en-US', {
    hour: 'numeric',
    minute: 'numeric',
    hour12: true,
  })
}

export function secondsToMinutes(nSeconds: number | null) {
  if (!nSeconds || nSeconds < 0) {
    return 0
  }

  return (nSeconds / 60).toFixed(1)
}

export function nowPlusNSeconds(nSeconds: number | null) {
  if (!nSeconds || nSeconds < 0) {
    return null
  }

  return new Date(new Date().getTime() + nSeconds * 1000).toLocaleTimeString(
    [],
    { hour: '2-digit', minute: '2-digit' },
  )
}
