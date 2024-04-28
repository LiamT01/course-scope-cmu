export const listOfferingsPageSize = 40;

// vercel.json doesn't support rewrites with SvelteKit
export const apiBaseUrl = import.meta.env.DEV ? 'http://localhost:8100' : 'https://course-scope-cmu.alwaysdata.net';
