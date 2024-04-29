import type {Cookies} from "@sveltejs/kit";
import {isTokenValid} from "$lib/auth/tokenValidity";

// Returns valid tokens, or null if invalid
export const getTokenFromCookies = async ({cookies}: { cookies: Cookies }) => {
    const token = cookies.get("token") ?? null;
    const expiry = cookies.get("expiry") ?? null;

    if (isTokenValid(expiry)) {
        return {
            token,
            expiry,
        }
    } else {
        return {
            token: null,
            expiry: null,
        }
    }
}
