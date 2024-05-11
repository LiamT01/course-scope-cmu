import makeQueryParams from '$lib/util/makeQueryParams';
import type { Course, Metadata, Offering } from '$lib/types';
import { fetchWithinLoad } from '$lib/auth/fetchWrappers';
import { apiBaseUrl, listOfferingsPageSize } from '$lib/constants';

export const load = async ({ fetch, params, url, data }) => {
	const courseID = parseInt(params.id);

	// API doesn't support "all" or "" for these fields
	const semester = url.searchParams.get('semester') ?? 'all';
	const year = url.searchParams.get('year') ?? 'all';
	const instructor_ids = url.searchParams.getAll('instructor_ids');
	const overall = url.searchParams.get('overall') ?? 'all';

	// Compatible with API
	const sort = url.searchParams.get('sort') ?? '-updated_at';
	const page: number = parseInt(url.searchParams.get('page') ?? '1');
	const pageSize: number = parseInt(url.searchParams.get('page_size') ?? '10');

	const courseResponse = await fetchWithinLoad(fetch, `${apiBaseUrl}/courses/${courseID}`);
	const course: Course = await courseResponse.json();

	const courseOfferingsResponse = await fetchWithinLoad(
		fetch,
		`${apiBaseUrl}/offerings?course_id=${courseID}&page_size=${listOfferingsPageSize}`
	);
	const courseOfferingsData: { items: Offering[]; metadata: Metadata } =
		await courseOfferingsResponse.json();
	const courseOfferings = courseOfferingsData.items;

	const allRatingsQuery = makeQueryParams({
		courseID,
		semester,
		year,
		instructorIDs: instructor_ids,
		overall,
		sort,
		page,
		pageSize
	});

	const ratingsResponse = await fetchWithinLoad(fetch, `${apiBaseUrl}/ratings?${allRatingsQuery}`, {
		token: data.token,
		expiry: data.expiry
	});
	const ratingsData = await ratingsResponse.json();
	const ratings = ratingsData.items;
	const metadata = ratingsData.metadata;

	const statsAllQuery = makeQueryParams({
		courseID,
		semester,
		year,
		instructorIDs: instructor_ids,
		overall
	});
	const statsResponse = await fetchWithinLoad(
		fetch,
		`${apiBaseUrl}/ratings/stats?${statsAllQuery}`
	);
	const stats = await statsResponse.json();

	return {
		course,
		courseOfferings,
		ratings,
		stats,
		sort,
		semester,
		year,
		instructorIDs: instructor_ids,
		overall,
		metadata
	};
};
