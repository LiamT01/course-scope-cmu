<script lang="ts">
	import { ChevronUpSolid } from 'flowbite-svelte-icons';
	import { fade } from 'svelte/transition';
	import { twMerge } from 'tailwind-merge';
	import { onMount } from 'svelte';
	import OutlineButton from '$lib/button/OutlineButton.svelte';

	let isVisible = false;

	const scrollToTop = () => {
		window.scrollTo({
			top: 0,
			behavior: 'smooth'
		});
	};

	const handleScroll = () => {
		const triggerY = 500;
		const maxScrollY = document.body.scrollHeight - window.innerHeight;
		isVisible = window.scrollY > triggerY && window.scrollY < maxScrollY;
	};

	onMount(() => {
		window.addEventListener('scroll', handleScroll);
		handleScroll(); // Initialize to set the correct state

		return () => {
			window.removeEventListener('scroll', handleScroll);
		};
	});
</script>

{#if isVisible}
	<div transition:fade={{ duration: 100 }} class="fixed bottom-32 right-8 z-10">
		<OutlineButton size="sm" on:click={scrollToTop} class={twMerge('h-10 w-10')}>
			<ChevronUpSolid class="h-4 w-4" />
		</OutlineButton>
	</div>
{/if}
