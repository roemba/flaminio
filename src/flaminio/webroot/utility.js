/* eslint-disable no-magic-numbers */

/**
 * From: https://stackoverflow.com/questions/2241447/make-foregroundcolor-black-or-white-depending-on-background
 * @return {number}
 */
export function perceivedBrightness(r, g, b) {
	return Math.round(Math.sqrt(
		r * r * .299 +
		g * g * .587 +
		b * b * .114));
}

export function hexToRgb(hex) {
	const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
	return result ? {
		r: parseInt(result[1], 16),
		g: parseInt(result[2], 16),
		b: parseInt(result[3], 16)
	} : null;
}

export function blackOrWhite(backgroundHex) {
	const rgb = hexToRgb(backgroundHex);
	return (perceivedBrightness(rgb.r, rgb.g, rgb.b) > 130 ? "black" : "white");
}

/* eslint-enable no-magic-numbers */
