"use strict";

const aspect = 11 / 8.5;
const raspect = 1 / aspect;
const initialWidth = 1080 * raspect;
const initialHeight = initialWidth * aspect;

function scale() {
	const r = 1 / initialWidth;
	const scale = r * Math.min(window.innerWidth, window.innerHeight * raspect);
	// const off = (window.innerWidth - document.body.getBoundingClientRect().width) / 2;

	document.body.style.setProperty('transform', ``)
	document.body.style.setProperty('--scale', `${scale}`);
	document.body.style.setProperty('transform-origin', 'left top');
	document.body.style.setProperty('transform', `scale(var(--scale))`);
	document.body.style.setProperty('width', `${initialWidth}px`);
	document.body.style.setProperty('height', `${initialHeight}px`);
	document.body.style.setProperty('overflow', `hidden`);
}

// window.addEventListener('resize', () => {
// 	scale();
// });
//
// setTimeout(() => {
// 	scale();
// })
//
// setTimeout(() => {
// 	scale();
// })
