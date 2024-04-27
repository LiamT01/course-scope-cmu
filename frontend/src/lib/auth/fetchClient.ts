import {handleTokenExpiryClient} from "$lib/auth/handleClient";
import type {Fetch} from "$lib/types";
import {isTokenValid} from "$lib/auth/stores";
import displayErrors from "$lib/util/displayErrors";
import {tokenStore, expiryStore} from "$lib/auth/stores";
import {get} from "svelte/store";

interface Init extends RequestInit {
    // Override headers
    headers?: { [key: string]: string };
}

interface InitWithToken extends Init {
    token: string | null;
    expiry: string | null;
}

const fetchBase = async (
    fetch: Fetch,
    input: RequestInfo | URL,
    init?: InitWithToken,
): Promise<Response> => {
    if (init?.token && isTokenValid(init?.expiry)) {
        init = init ?? {};
        init.headers = init.headers ?? {};

        init.headers["Authorization"] = `Bearer ${init.token}`;
    }

    try {
        const response = await fetch(input, init);

        if (!response.ok) {
            const data = await response.json();
            displayErrors(data.detail ?? data.message);
            return new Response(JSON.stringify(data), {status: response.status});
        }

        return response;
    } catch (e) {
        console.error(e);
        throw e;
    }
}

// (1) If token is expired, logs user out and prompts user to log in again.
// (2) Includes valid token in headers.
// (3) Displays any errors for api call.
export const fetchWithinPage = async (
    input: RequestInfo | URL,
    init?: Init,
): Promise<Response> => {
    const initWithToken: InitWithToken = {
        token: get(tokenStore),
        expiry: get(expiryStore),
        ...init,
    }
    await handleTokenExpiryClient(fetch, initWithToken?.expiry, true);
    return await fetchBase(fetch, input, initWithToken);
}

// (1) If token is expired, logs user out and prompts user to log in again.
// (2) Includes valid token in headers.
// (3) Displays any errors for api call.
export const fetchWithinLoad = async (
    fetch: Fetch,
    input: RequestInfo | URL,
    init?: InitWithToken,
): Promise<Response> => {
    await handleTokenExpiryClient(fetch, init?.expiry, false);
    return await fetchBase(fetch, input, init);
}
