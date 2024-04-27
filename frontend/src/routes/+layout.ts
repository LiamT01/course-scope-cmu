import {expiryStore, isTokenValid, tokenStore, userStore} from "$lib/auth/stores";
import {addErrorToast} from "$lib/toast/stores";
import {fetchWithinLoad} from "$lib/auth/fetchClient";
import type {User} from "$lib/types";

export const ssr = false;

export const load = async ({data, fetch}) => {
    if (!data.token || !data.expiry || !isTokenValid(data.expiry)) {
        return;
    }


    const userResponse = await fetchWithinLoad(
        fetch,
        "/api/users/me",
        {
            token: data.token,
            expiry: data.expiry,
        }
    )

    if (userResponse.ok) {
        tokenStore.set(data.token);
        expiryStore.set(data.expiry);

        const userData: User = await userResponse.json();
        userStore.set(userData);

        if (!userData.activated) {
            addErrorToast(
                "Your account is not activated. You can request a new activation link in the dashboard."
            );
        }
    }
}
