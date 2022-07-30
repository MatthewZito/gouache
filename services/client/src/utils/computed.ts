/**
 * Cycle to the next or previous index in a range of indices.
 * @param direction Direction in which to cycle e.g. 1 for next index, -1 for previous index
 * @param current Starting index from which to resolve next/previous index
 * @param total Total number of indices
 */
export function cycleRange(direction: -1 | 1, current: number, total: number) {
	return (current + direction + total) % total;
}
