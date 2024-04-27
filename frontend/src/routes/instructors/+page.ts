import makeQueryParams from "$lib/util/makeQueryParams";
import {fetchWithinLoad} from "$lib/auth/fetchClient";

export const load = async ({fetch, url, data}) => {
    try {
        const name = url.searchParams.get('name') ?? '';
        const page = url.searchParams.get('page') ?? '1';
        const pageSize = url.searchParams.get('page_size') ?? '10';
        const sort = url.searchParams.get('sort') ?? 'name';

        const query = makeQueryParams({name, page, pageSize, sort});
        const instructorsResponse = await fetchWithinLoad(
            fetch,
            `/api/instructors?${query}`,
            {
                token: data.token,
                expiry: data.expiry,
            }
        )
        const instructorsData = await instructorsResponse.json()
        return {
            instructors: instructorsData.items,
            metadata: instructorsData.metadata,
        }
    } catch (e) {
        console.error(e);
        return {
            instructors: [],
            metadata: {},
        }
    }
}