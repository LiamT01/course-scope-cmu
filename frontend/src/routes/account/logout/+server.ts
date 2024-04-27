import type {RequestHandler} from './$types';
import {isTokenValid} from "$lib/auth/stores";

export const POST: RequestHandler = async ({fetch, cookies}) => {
    const token = cookies.get('token') ?? null;
    const expiry = cookies.get('expiry') ?? null;

    cookies.delete('token', {path: '/'});
    cookies.delete('expiry', {path: '/'});

    if (isTokenValid(expiry)) {
        await fetch('/api/tokens/auth/me', {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
        });
    }

    return new Response(null, {status: 204});
};
