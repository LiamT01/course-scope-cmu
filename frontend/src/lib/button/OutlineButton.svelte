<script lang="ts">
	import { twMerge } from 'tailwind-merge';
	import GhostButton from '$lib/button/GhostButton.svelte';
</script>

<GhostButton
	{...$$restProps}
	class={twMerge('border border-gray-300 bg-gray-50', $$props.class)}
	color="none"
	on:click
>
	<slot />
</GhostButton>
