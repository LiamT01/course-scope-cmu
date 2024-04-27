<script lang="ts">
    import { twMerge } from 'tailwind-merge';
    import { getContext } from 'svelte';

    export let active: boolean = false;
    export let activeClass: string = 'text-blue-600 border border-gray-300 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white';
    export let normalClass: string = 'text-gray-500 bg-white hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white';
    export let large: boolean = false;

    const group = getContext<boolean>('group');
    const table = getContext<boolean>('table');

    let defaultClass: string;
    $: defaultClass = twMerge(
        'flex items-center justify-center font-medium',
        large ? 'h-7 w-7 sm:h-10 sm:w-10 text-base' : 'h-7 w-7 sm:h-10 sm:w-10 px-3 text-sm',
        group ? '' : table ? 'rounded' : 'rounded-lg',
        // table || 'border border-gray-300 dark:border-gray-700 dark:bg-gray-800',
        table ? '' : 'border',
        active ? activeClass : normalClass,
        $$props.class
    );
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!--<svelte:element this={href ? 'a' : 'button'} {href} class={defaultClass} on:blur on:change on:click on:focus on:keydown on:keypress on:keyup on:mouseenter on:mouseleave on:mouseover role={href ? 'button' : undefined}>-->
<!--    <slot />-->
<!--</svelte:element>-->

<svelte:element this='button' class={defaultClass} on:blur on:change on:click on:focus on:keydown on:keypress on:keyup on:mouseenter on:mouseleave on:mouseover>
    <slot />
</svelte:element>

