<script>
	import { Button } from 'flowbite-svelte';
	import { twMerge } from 'tailwind-merge';
</script>

<Button
	{...$$restProps}
	class={twMerge(
		'rounded-lg border-red-300 bg-red-600 text-white hover:bg-red-700 dark:hover:bg-red-700',
		$$props.class
	)}
	color="none"
	on:click
>
	<slot />
</Button>
