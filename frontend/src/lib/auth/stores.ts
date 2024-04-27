import {writable} from "svelte/store";
import type {User} from "$lib/types";

export const tokenStore = writable<string | null>(null);
export const expiryStore = writable<string | null>(null);

export const userStore = writable<User | null>(null);

export const isTokenValid = (expiry: string | null | undefined): boolean => {
    const expiryDate = expiry ? new Date(expiry) : null;
    return !!(expiryDate && expiryDate > new Date());
}

export const isTokenExpired = (expiry: string | null | undefined): boolean => {
    const expiryDate = expiry ? new Date(expiry) : null;
    return !!(expiryDate && expiryDate < new Date());
}
