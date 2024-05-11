import { browser } from '$app/environment';
import { addErrorToast } from '$lib/toast/stores';
import displayErrors from '$lib/util/displayErrors';
import { isTokenExpired } from '$lib/auth/tokenValidity';
import { invalidateAll } from '$app/navigation';

export async function handleTokenExpiryClient(
	fetch: (input: RequestInfo | URL, init?: RequestInit | undefined) => Promise<Response>,
	expiry: string | null | undefined
): Promise<boolean> {
	if (!browser) {
		return false;
	}

	if (isTokenExpired(expiry)) {
		addErrorToast('Your session has expired. Please log in again.');

		try {
			const response = await fetch('/account/logout', {
				method: 'POST'
			});

			if (!response.ok) {
				const data = await response.json();
				displayErrors(data.detail ?? data.message);
			}

			await invalidateAll();
		} catch (e) {
			console.error(e);
		}

		return true;
	}

	return false;
}
