<script lang="ts">
	import type { Instructor, Metadata, Offering, RatingStatsT } from '$lib/types';
	import round, { roundFixed } from '$lib/util/round';
	import RatingIcons from '$lib/rating/RatingIcons.svelte';
	import { onMount } from 'svelte';
	import CourseOfferingTable from '$lib/course/CourseOfferingTable.svelte';
	import { fetchWithinPage } from '$lib/auth/fetchWrappers';
	import { apiBaseUrl, listOfferingsPageSize } from '$lib/constants';

	export let instructor: Instructor;

	let stats: RatingStatsT | null;

	onMount(async () => {
		stats = await fetchStats();
	});

	async function fetchStats() {
		const statsResponse = await fetchWithinPage(
			`${apiBaseUrl}/ratings/stats?instructor_ids=${instructor.id}`
		);
		return await statsResponse.json();
	}

	let instructorOfferings: Offering[] | null;

	onMount(async () => {
		instructorOfferings = await fetchInstructorOfferings();
	});

	async function fetchInstructorOfferings() {
		const instructorOfferingsResponse = await fetchWithinPage(
			`${apiBaseUrl}/offerings?instructor_ids=${instructor.id}&page_size=${listOfferingsPageSize}`
		);
		const instructorOfferingsData = (await instructorOfferingsResponse.json()) as {
			items: Offering[];
			metadata: Metadata;
		};
		return instructorOfferingsData.items;
	}
</script>

<div class="flex flex-col gap-y-4 border-y py-4 sm:py-8" id="instructor-{instructor.id}">
	<div class="flex flex-col gap-y-4">
		<p class="text-lg font-medium">
			<a href="/instructors/{instructor.id}">{instructor.name}</a>
		</p>

		{#if stats}
			<RatingIcons size={20} text="sm" rating={round(stats.avg_overall, 1)} />
			<div class="flex gap-x-8 text-sm font-medium text-gray-600 dark:text-gray-400">
				<div
					class="grid grid-flow-col grid-rows-3 gap-x-6 gap-y-2 sm:grid-flow-row sm:grid-cols-3 sm:grid-rows-none sm:gap-x-16"
				>
					<span class="grid grid-cols-[72px_max-content]">
						<span>Teaching:</span>
						<span>{roundFixed(stats.avg_teaching, 1)} / 5.0</span>
					</span>
					<span class="grid grid-cols-[72px_max-content]">
						<span>Materials:</span>
						<span>{roundFixed(stats.avg_materials, 1)} / 5.0</span>
					</span>
					<span class="grid grid-cols-[72px_max-content]">
						<span>Value:</span>
						<span>{roundFixed(stats.avg_value, 1)} / 5.0</span>
					</span>
				</div>
			</div>
		{/if}
	</div>

	<div class="flex flex-col gap-y-2 border-y">
		<CourseOfferingTable {instructorOfferings} />
	</div>
</div>
