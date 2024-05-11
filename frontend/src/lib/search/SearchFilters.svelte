<script lang="ts">
	import { MultiSelect, Select } from 'flowbite-svelte';
	import { twMerge } from 'tailwind-merge';
	import PrimaryButton from '$lib/button/PrimaryButton.svelte';

	export let sort: string = 'newest';
	export let semester: string = 'all';
	export let year: string = 'all';
	export let instructors: string[] = [];

	export let semesterOptions: { value: string; name: string }[];
	export let yearOptions: { value: string; name: string }[];
	export let instructorOptions: { value: string; name: string }[];
	export let onSubmit: () => void;

	const sortByOptions = [
		{ value: '-updated_at', name: 'Newest' },
		{ value: 'updated_at', name: 'Oldest' },
		{ value: '-net_likes', name: 'Most liked' },
		{ value: 'net_likes', name: 'Least liked' },
		{ value: '-overall', name: 'Highest rated' },
		{ value: 'overall', name: 'Lowest rated' }
	];
</script>

<div
	class={twMerge(
		'flex flex-col justify-between gap-x-4 gap-y-4 sm:flex-row sm:items-end',
		$$props.class
	)}
>
	<!--    <div class="flex flex-col gap-y-2">-->
	<div class="flex justify-between gap-x-4 sm:gap-x-4">
		<div class="flex-1 space-y-2">
			<span class="whitespace-nowrap text-sm font-medium md:order-1">Sort by</span>
			<div class="mb-2 flex flex-wrap gap-x-4 gap-y-2 md:order-6 md:mb-0">
				<Select
					bind:value={sort}
					class="overflow-hidden text-xs sm:text-sm"
					items={sortByOptions}
					placeholder=""
					size="sm"
				/>
				<!--{#each sortByOptions as {value, name}}-->
				<!--    <Radio bind:group={sort} {value}>{name}</Radio>-->
				<!--{/each}-->
			</div>
		</div>

		<div class="space-y-2">
			<span class="whitespace-nowrap text-sm font-medium md:order-2">Semester</span>
			<Select
				bind:value={semester}
				class="mb-2 overflow-hidden text-xs sm:text-sm md:order-7 md:mb-0"
				items={semesterOptions}
				placeholder=""
				size="sm"
			/>
		</div>

		<div class="space-y-2">
			<span class="whitespace-nowrap text-sm font-medium md:order-3">Year</span>
			<Select
				bind:value={year}
				class="mb-2 overflow-hidden text-xs sm:text-sm md:order-8 md:mb-0"
				items={yearOptions}
				placeholder=""
				size="sm"
			/>
		</div>
	</div>

	<div
		class="flex flex-1 flex-col gap-x-2 gap-y-4 sm:flex-row sm:items-end sm:justify-between sm:gap-x-4"
	>
		<div class="flex-1 space-y-2">
			<span class="whitespace-nowrap text-sm font-medium md:order-4">Instructors ("OR"ed)</span>
			<MultiSelect
				bind:value={instructors}
				class="mb-2 md:order-9 md:mb-0"
				dropdownClass="text-sm"
				items={instructorOptions}
				on:change
				on:input
				size="sm"
			/>
		</div>

		<PrimaryButton class="md:order-10" on:click={onSubmit} size="sm">Apply</PrimaryButton>
	</div>

	<!--    </div>-->
	<!--    <Label class="flex flex-col gap-y-2">-->

	<!--    </Label>-->
	<!--    <Label class="flex flex-col gap-y-2">-->
	<!--        <span class="whitespace-nowrap">Year</span>-->
	<!--    </Label>-->
	<!--    <Label class="flex flex-col gap-y-2">-->

	<!--    </Label>-->
	<!--    <Label class="flex flex-col gap-y-2">-->
	<!--        <span class="whitespace-nowrap">Stars</span>-->
	<!--        <Select bind:value={overall} class="overflow-hidden" items={ratingOptions} placeholder="" size="sm"/>-->
	<!--    </Label>-->
</div>
