<script>
	import { Button } from 'flowbite-svelte';
	import { twMerge } from 'tailwind-merge';
</script>

<Button
	{...$$restProps}
	class={twMerge(
		'rounded-lg border-gray-300 bg-gray-600 text-white hover:bg-gray-700 dark:hover:bg-gray-700',
		$$props.class
	)}
	color="none"
	on:click
>
	<slot />
</Button>
