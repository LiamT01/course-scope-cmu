<script lang="ts">
	import {
		ThumbsDownOutline,
		ThumbsDownSolid,
		ThumbsUpOutline,
		ThumbsUpSolid
	} from 'flowbite-svelte-icons';
	import GhostButton from '$lib/button/GhostButton.svelte';
	import { twMerge } from 'tailwind-merge';

	export let netLikes: number;
	export let likedByUser: boolean;
	export let dislikedByUser: boolean;
	export let onLike: () => void;
	export let onDislike: () => void;
</script>

<div class={twMerge('flex items-center gap-x-2 text-sm', $$props.class)}>
	<GhostButton class="h-10 w-10" on:click={onLike} size="xs">
		{#if likedByUser}
			<ThumbsUpSolid class="h-4 w-4" />
		{:else}
			<ThumbsUpOutline class="h-4 w-4" />
		{/if}
	</GhostButton>
	{netLikes}
	<GhostButton class="ml-4 h-10 w-10" on:click={onDislike} size="xs">
		{#if dislikedByUser}
			<ThumbsDownSolid class="h-4 w-4" />
		{:else}
			<ThumbsDownOutline class="h-4 w-4" />
		{/if}
	</GhostButton>
</div>
