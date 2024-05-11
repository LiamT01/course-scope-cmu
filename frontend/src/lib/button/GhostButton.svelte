<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { twMerge } from 'tailwind-merge';
</script>

<Button
	{...$$restProps}
	class={twMerge(
		'rounded-lg text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:bg-gray-700',
		$$props.class
	)}
	color="none"
	on:click
>
	<slot />
</Button>
