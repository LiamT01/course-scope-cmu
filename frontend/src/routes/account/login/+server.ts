import type {RequestHandler} from './$types';
import {apiBaseUrl} from "$lib/constants";

interface TokenResponse {
    token: string;
    expiry: string;
}

export const POST: RequestHandler = async ({request, fetch, cookies}) => {
    // Parse the form data from the request
    const formData = await request.formData();
    const requestBody: { [key: string]: any } = {};

    // Convert the form data to a JSON object
    for (const [key, value] of formData) {
        requestBody[key] = value;
    }

    const response = await fetch(`${apiBaseUrl}/tokens/auth`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody)
    });
    const data = await response.json() as TokenResponse;

    if (response.ok) {
        cookies.set('token', data.token, {
            httpOnly: true,
            secure: true,
            sameSite: 'none',
            path: '/',
            expires: new Date(data.expiry),
        });
        cookies.set('expiry', data.expiry, {
            httpOnly: true,
            secure: true,
            sameSite: 'none',
            path: '/',
            expires: new Date(data.expiry),
        });

        return new Response(JSON.stringify({
            "message": "Success",
            "detail": "Successfully logged in."
        }), {status: response.status});
    } else {
        return new Response(JSON.stringify(data), {status: response.status});
    }
};
