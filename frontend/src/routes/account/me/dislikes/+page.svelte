<script lang="ts">
	import Pagination from '$lib/pagination/Pagination.svelte';
	import RatingCard from '$lib/rating-card/RatingCard.svelte';
	import type { Metadata, Rating, User } from '$lib/types';
	import { goto } from '$app/navigation';
	import OutlineButton from '$lib/button/OutlineButton.svelte';
	import Title from '$lib/section/Title.svelte';

	export let data: {
		metadata: Metadata;
		ratings: Rating[];
		token: string;
		expiry: string;
		user: User | null;
	};

	const previous = () => {
		if (data.metadata.current_page > data.metadata.first_page) {
			goto(`?page=${data.metadata.current_page - 1}&page_size=${data.metadata.page_size}`);
		}
	};
	const next = () => {
		if (data.metadata.current_page < data.metadata.last_page) {
			goto(`?page=${data.metadata.current_page + 1}&page_size=${data.metadata.page_size}`);
		}
	};

	const jump = (e: { detail: { page: number } }) => {
		if (
			e.detail.page >= data.metadata.first_page &&
			e.detail.page <= data.metadata.last_page &&
			e.detail.page !== data.metadata.current_page
		) {
			goto(`?page=${e.detail.page}&page_size=${data.metadata.page_size}`);
		}
	};
</script>

{#if data.user}
	<div class="flex flex-col gap-y-4">
		{#if data.ratings.length > 0}
			<Title>Disliked ratings ({data.metadata.total})</Title>

			{#each data.ratings as rating (rating.id)}
				<RatingCard
					token={data.token}
					expiry={data.expiry}
					user={data.user}
					{rating}
					showCourseDetail
				/>
			{/each}

			<div class="mx-auto">
				<Pagination
					metadata={data.metadata}
					icon
					on:next={next}
					on:previous={previous}
					on:jump={jump}
				/>
			</div>
		{:else if data.metadata.current_page > 1}
			<OutlineButton on:click={previous}>
				No ratings on page {data.metadata.current_page}. Click here to go back.
			</OutlineButton>
		{:else}
			<p class="text-center">No ratings yet.</p>
		{/if}
	</div>
{/if}
