import makeQueryParams from '$lib/util/makeQueryParams';
import { fetchWithinLoad } from '$lib/auth/fetchWrappers';
import type { Course, Metadata } from '$lib/types';
import { apiBaseUrl } from '$lib/constants';

export const load = async ({ fetch, url }) => {
	const number = url.searchParams.get('number') ?? '';
	const name = url.searchParams.get('name') ?? '';
	const page = url.searchParams.get('page') ?? '1';
	const pageSize = url.searchParams.get('page_size') ?? '10';
	const sort = url.searchParams.get('sort') ?? 'number';

	const query = makeQueryParams({ number, name, page, pageSize, sort });

	const coursesResponse = await fetchWithinLoad(fetch, `${apiBaseUrl}/courses?${query}`);

	const coursesData: { items: Course[]; metadata: Metadata } = await coursesResponse.json();
	return {
		courses: coursesData.items,
		metadata: coursesData.metadata,
		number,
		name,
		sort
	};
};
