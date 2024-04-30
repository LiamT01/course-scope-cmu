import {isTokenValid} from "$lib/auth/tokenValidity";
import {addErrorToast} from "$lib/toast/stores";
import {fetchWithinLoad} from "$lib/auth/fetchWrappers";
import type {User} from "$lib/types";
import {browser} from '$app/environment';
import {apiBaseUrl} from "$lib/constants";

// export const ssr = false;

export const load = async ({data, fetch}): Promise<{
    token: string | null,
    expiry: string | null,
    user: User | null,
}> => {
    if (!data.token || !data.expiry || !isTokenValid(data.expiry)) {
        return {
            token: null,
            expiry: null,
            user: null,
        };
    }

    const userResponse = await fetchWithinLoad(
        fetch,
        `${apiBaseUrl}/users/me`,
        {
            token: data.token,
            expiry: data.expiry,
        }
    )

    if (userResponse.ok) {
        const userData: User = await userResponse.json();
        if (!userData.activated && browser) {
            addErrorToast(
                "Your account is not activated. You can request a new activation link in the dashboard."
            );
        }
        return {
            token: data.token,
            expiry: data.expiry,
            user: userData,
        };
    } else {
        await fetch('/account/logout', {
            method: 'POST',
        });
        return {
            token: null,
            expiry: null,
            user: null,
        }
    }
}
