import {fetchWithinLoad} from "$lib/auth/fetchClient";
import type {Metadata, Rating, UserStats} from "$lib/types";
import {userStore} from "$lib/auth/stores";
import {get} from "svelte/store";

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
        `/api/ratings/my?page=${page}&page_size=${pageSize}&sort=-updated_at`,
        {
            token: data.token,
            expiry: data.expiry,
        }
    );
    const ratingsData: {items: Rating[], metadata: Metadata} = await ratingsResponse.json()

    const userStatsResponse = await fetchWithinLoad(
        fetch,
        `/api/users/stats/me`,
        {
            token: data.token,
            expiry: data.expiry,
        }
    );
    const userStatsData: UserStats = await userStatsResponse.json()

    return {
        ratings: ratingsData.items,
        metadata: ratingsData.metadata,
        userStats: userStatsData,
    }
}
