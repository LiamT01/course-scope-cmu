export default function round(num: number, places: number) {
	const multiplier = Math.pow(10, places);
	return Math.round(num * multiplier) / multiplier;
}

// Force round, padding with 0s, returning string
export function roundFixed(num: number, places: number) {
	const multiplier = Math.pow(10, places);
	return (Math.round(num * multiplier) / multiplier).toFixed(places);
}
