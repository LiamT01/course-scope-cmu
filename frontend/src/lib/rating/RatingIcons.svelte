<script lang="ts">
	import { roundFixed } from '$lib/util/round';
	import Star from '$lib/rating/Star.svelte';
	import type { ComponentType } from 'svelte';
	import generateID from '$lib/util/generateID';
	import { twMerge } from 'tailwind-merge';

	export let icon: ComponentType = Star;
	export let rating: number;
	export let size: number = 20;
	export let text: string = 'sm';
	export let showText: boolean = false;

	export let divClass: string = 'flex items-center';
	export let total: number = 5;
	export let partialId: string = 'partialStar' + generateID();
	export let count: boolean = false;

	export let edit: boolean = false;

	// generate unique id for full star and gray star
	const fullStarId: string = generateID();
	const grayStarId: string = generateID();
	let fullStars = Math.floor(rating);
	let rateDifference = rating - fullStars;
	let percentRating = Math.round(rateDifference * 100);
	let grayStars: number = total - (fullStars + Math.ceil(rateDifference));
	$: {
		fullStars = Math.floor(rating);
		rateDifference = rating - fullStars;
		percentRating = Math.round(rateDifference * 100);
		grayStars = total - (fullStars + Math.ceil(rateDifference));
	}
</script>

<div class="flex flex-row items-center">
	{#key rating}
		<div class={twMerge(divClass, $$props.class)}>
			{#if count}
				{#if edit}
					<svelte:component this={icon} class="cursor-pointer" on:click fillPercent={100} {size} />
				{:else}
					<svelte:component this={icon} fillPercent={100} {size} />
				{/if}

				<p class="ms-2 text-sm font-bold text-gray-900 dark:text-white">{rating}</p>
				<slot />
			{:else}
				{#each Array.from(Array(fullStars).keys()) as index}
					{#if edit}
						<svelte:component
							this={icon}
							class="cursor-pointer"
							on:click={() => (rating = index + 1)}
							{size}
							fillPercent={100}
							id={fullStarId}
						/>
					{:else}
						<svelte:component this={icon} {size} fillPercent={100} id={fullStarId} />
					{/if}
				{/each}
				{#if percentRating}
					{#if edit}
						<svelte:component
							this={icon}
							class="cursor-pointer"
							on:click={() => (rating = fullStars + 1)}
							{size}
							fillPercent={percentRating}
							id={partialId}
						/>
					{:else}
						<svelte:component this={icon} {size} fillPercent={percentRating} id={partialId} />
					{/if}
				{/if}
				{#each Array.from(Array(grayStars).keys()) as index}
					{#if edit}
						<svelte:component
							this={icon}
							class="cursor-pointer"
							on:click={() => (rating = total - grayStars + 1 + index)}
							{size}
							fillPercent={0}
							id={grayStarId}
						/>
					{:else}
						<svelte:component this={icon} {size} fillPercent={0} id={grayStarId} />
					{/if}
				{/each}
				{#if $$slots.text}
					<slot name="text" />
				{/if}
			{/if}
		</div>
	{/key}
	{#if showText}
		<span class="ml-[20px] text-{text} font-medium text-gray-600 dark:text-gray-400">
			{roundFixed(rating, 1)} / 5.0
		</span>
	{/if}
</div>
