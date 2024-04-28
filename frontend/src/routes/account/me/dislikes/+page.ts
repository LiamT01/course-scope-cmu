import {fetchWithinLoad} from "$lib/auth/fetchClient";
import type {Metadata, Rating} from "$lib/types";
import {apiBaseUrl} from "$lib/constants";

export const load = async ({fetch, url, data}) => {
    if (!data.token || !data.expiry) {
        return {
            ratings: [],
            metadata: {
                total: 0,
                page: 1,
                page_size: 10,
            },
        }
    }

    const page: number = parseInt(url.searchParams.get('page') ?? '1');
    const pageSize: number = parseInt(url.searchParams.get('page_size') ?? '10');
    const ratingsResponse = await fetchWithinLoad(
        fetch,
        `${apiBaseUrl}/ratings/my-dislikes?page=${page}&page_size=${pageSize}&sort=-updated_at`,
        {
            token: data.token,
            expiry: data.expiry,
        }
    );
    const ratingsData: {items: Rating[], metadata: Metadata} = await ratingsResponse.json()

    return {
        ratings: ratingsData.items,
        metadata: ratingsData.metadata,
    }
}
