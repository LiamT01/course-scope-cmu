import makeQueryParams from '$lib/util/makeQueryParams';
import { fetchWithinLoad } from '$lib/auth/fetchWrappers';
import type { Instructor, Metadata } from '$lib/types';
import { apiBaseUrl } from '$lib/constants';

export const load = async ({ fetch, url }) => {
	const name = url.searchParams.get('name') ?? '';
	const page = url.searchParams.get('page') ?? '1';
	const pageSize = url.searchParams.get('page_size') ?? '10';
	const sort = url.searchParams.get('sort') ?? 'name';

	const query = makeQueryParams({ name, page, pageSize, sort });
	const instructorsResponse = await fetchWithinLoad(fetch, `${apiBaseUrl}/instructors?${query}`);
	const instructorsData: { items: Instructor[]; metadata: Metadata } =
		await instructorsResponse.json();
	return {
		instructors: instructorsData.items,
		metadata: instructorsData.metadata,
		name,
		sort
	};
};
