<script lang="ts">
	import { twMerge } from 'tailwind-merge';
	import { createEventDispatcher, setContext } from 'svelte';
	import PaginationItem from '$lib/pagination/PaginationItem.svelte';
	import type { Metadata } from '$lib/types';
	import { ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';

	export let activeClass: string =
		'text-gray-600 bg-gray-200 dark:border-gray-700 dark:bg-gray-700 dark:text-white';
	export let normalClass: string =
		'text-gray-500 bg-white hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white';
	export let ulClass: string = 'inline-flex -space-x-px rtl:space-x-reverse items-center flex-wrap';
	export let table: boolean = false;
	export let large: boolean = false;
	export let ariaLabel: string = 'Page navigation';
	export let metadata: Metadata;
	export let numNeighbours: number = 2;

	let pages: { name: number; active: boolean }[] = [];
	$: {
		const start = Math.max(metadata.current_page - numNeighbours, metadata.first_page);
		const end = Math.min(metadata.current_page + numNeighbours + 1, metadata.last_page + 1);
		pages = [
			...Array.from({ length: end - start }, (_, k) => k + start).map((page) => ({
				name: page,
				active: page === metadata.current_page
			}))
		];

		// If there is discontinuity at the beginning or end,
		// add '...' to the list plus the first or last page
		if (pages[0].name > metadata.first_page) {
			if (pages[0].name > metadata.first_page + 1) {
				pages.unshift({
					name: -1,
					active: false
				});
			}
			pages.unshift({
				name: metadata.first_page,
				active: false
			});
		}

		if (pages[pages.length - 1].name < metadata.last_page) {
			if (pages[pages.length - 1].name < metadata.last_page - 1) {
				pages.push({
					name: -1,
					active: false
				});
			}
			pages.push({
				name: metadata.last_page,
				active: false
			});
		}
	}

	const dispatch = createEventDispatcher();

	setContext<boolean>('group', true);
	setContext<boolean>('table', table);

	const previous = () => {
		dispatch('previous');
	};
	const next = () => {
		dispatch('next');
	};

	const jump = (page: number) => {
		dispatch('jump', { page });
	};
</script>

<nav aria-label={ariaLabel} class={$$props.class}>
	<ul
		class={twMerge(
			ulClass,
			table && 'dark divide-x divide-gray-700 dark:divide-gray-700 rtl:divide-x-reverse',
			$$props.class
		)}
	>
		<li>
			<PaginationItem
				class={table ? 'rounded-l' : 'rounded-s-lg'}
				{large}
				{normalClass}
				on:click={previous}
			>
				<span class="sr-only">Previous</span>
				<ChevronLeftOutline class="h-2.5 w-2.5" />
			</PaginationItem>
		</li>
		{#each pages as { name, active }}
			<li>
				<PaginationItem
					on:click={() => jump(name)}
					{large}
					{active}
					{activeClass}
					{normalClass}
					on:blur
					on:change
					on:focus
					on:keydown
					on:keypress
					on:keyup
					on:mouseenter
					on:mouseleave
					on:mouseover
				>
					{name !== -1 ? name : '...'}
				</PaginationItem>
			</li>
		{/each}
		<li>
			<PaginationItem
				class={table ? 'rounded-r' : 'rounded-e-lg'}
				{large}
				{normalClass}
				on:click={next}
			>
				<span class="sr-only">Next</span>
				<ChevronRightOutline class="h-2.5 w-2.5" />
			</PaginationItem>
		</li>
	</ul>
</nav>
