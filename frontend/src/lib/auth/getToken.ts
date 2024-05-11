import type { Cookies } from '@sveltejs/kit';
import { isTokenValid } from '$lib/auth/tokenValidity';

// Returns valid tokens, or null if invalid
export const getTokenFromCookies = async ({
	cookies
}: {
	cookies: Cookies;
}): Promise<{
	token: string | null;
	expiry: string | null;
}> => {
	const token = cookies.get('token') ?? null;
	const expiry = cookies.get('expiry') ?? null;

	if (token && expiry && isTokenValid(expiry)) {
		return {
			token,
			expiry
		};
	} else {
		return {
			token: null,
			expiry: null
		};
	}
};
