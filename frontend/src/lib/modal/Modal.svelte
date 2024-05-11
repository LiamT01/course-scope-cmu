<script lang="ts">
	import { fade } from 'svelte/transition';
	import { onMount } from 'svelte';
	import focusTrap from '$lib/actions/focusTrap';
	import { twMerge } from 'tailwind-merge';
	import { CloseOutline } from 'flowbite-svelte-icons';
	import GhostButton from '$lib/button/GhostButton.svelte';

	export let open = false;

	const hide = (e?: Event) => {
		e?.preventDefault();
		open = false;
	};

	// Add listener to close modal on hitting escape
	onMount(() => {
		window.addEventListener('keydown', handleKeys);
		document.body.style.overflow = 'hidden';
		return () => {
			window.removeEventListener('keydown', handleKeys);
			document.body.style.overflow = '';
		};
	});

	const handleKeys = (e: KeyboardEvent) => {
		if (e.key === 'Escape') return hide(e);
	};

	const isScrollable = (e: HTMLElement): boolean[] => [
		e.scrollWidth > e.clientWidth && ['scroll', 'auto'].indexOf(getComputedStyle(e).overflowX) >= 0,
		e.scrollHeight > e.clientHeight &&
			['scroll', 'auto'].indexOf(getComputedStyle(e).overflowY) >= 0
	];

	function prepareFocus(node: HTMLElement) {
		const walker = document.createTreeWalker(node, NodeFilter.SHOW_ELEMENT);
		let n: Node | null;
		while ((n = walker.nextNode())) {
			if (n instanceof HTMLElement) {
				const el = n as HTMLElement;
				const [x, y] = isScrollable(el);
				if (x || y) el.tabIndex = 0;
			}
		}
		node.focus();
	}
</script>

<div
	class="fixed inset-0 z-50 h-screen w-screen bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
	transition:fade={{ duration: 100 }}
/>
<!--    <div use:prepareFocus use:focusTrap on:wheel|preventDefault|nonpassive on:touchmove|preventDefault|nonpassive-->
<div
	aria-modal="true"
	class="fixed end-0 start-0 top-0 z-50 flex h-modal w-full items-center justify-center p-4 md:inset-0 md:h-full"
	role="dialog"
	tabindex="-1"
	use:focusTrap
	use:prepareFocus
>
	<!--        <div use:clickOutside={{callback: hide, excludedElement: null}}-->
	<div
		class={twMerge(
			'relative max-h-full w-full max-w-md rounded-lg border bg-white p-4 shadow',
			$$props.class
		)}
	>
		<GhostButton class="absolute end-3 top-3 h-10 w-10 cursor-pointer" on:click={hide} size="xs">
			<CloseOutline class="h-3 w-3" />
		</GhostButton>
		<slot />
	</div>
</div>
