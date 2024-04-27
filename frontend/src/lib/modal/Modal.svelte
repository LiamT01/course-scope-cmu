<script lang="ts">
    import {fade} from 'svelte/transition'
    import {onMount} from "svelte";
    import {clickOutside} from "$lib/actions/clickOutside";
    import focusTrap from "$lib/actions/focusTrap";
    import {twMerge} from "tailwind-merge";
    import {CloseOutline} from "flowbite-svelte-icons";
    import GhostButton from "$lib/button/GhostButton.svelte";

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
        }
    });

    const handleKeys = (e: KeyboardEvent) => {
        if (e.key === 'Escape') return hide(e);
    }

    const isScrollable = (e: HTMLElement): boolean[] => [
        e.scrollWidth > e.clientWidth && ['scroll', 'auto'].indexOf(getComputedStyle(e).overflowX) >= 0,
        e.scrollHeight > e.clientHeight && ['scroll', 'auto'].indexOf(getComputedStyle(e).overflowY) >= 0
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

<div transition:fade={{duration: 100}} class="w-screen h-screen fixed inset-0 z-50 bg-gray-900 bg-opacity-50 dark:bg-opacity-80"/>
<!--    <div use:prepareFocus use:focusTrap on:wheel|preventDefault|nonpassive on:touchmove|preventDefault|nonpassive-->
    <div use:prepareFocus use:focusTrap
         class="fixed justify-center items-center top-0 start-0 end-0 h-modal md:inset-0 md:h-full z-50 w-full p-4 flex"
         tabindex="-1" aria-modal="true" role="dialog">
    <!--        <div use:clickOutside={{callback: hide, excludedElement: null}}-->
        <div
             class={twMerge("border rounded-lg bg-white shadow p-4 relative max-w-md w-full max-h-full", $$props.class)}>
            <GhostButton size="xs" class="w-8 h-8 absolute top-4 end-4 cursor-pointer" on:click={hide}>
                <CloseOutline class="w-3 h-3" />
            </GhostButton>
            <slot/>
    </div>
</div>
