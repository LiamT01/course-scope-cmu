<script lang="ts">
    import {Button} from "flowbite-svelte";
    import {twMerge} from "tailwind-merge";
</script>

<Button
        {...$$restProps}
        class={twMerge(
            "text-gray-900 dark:text-gray-400 dark:hover:bg-gray-700 hover:bg-gray-100 rounded-lg",
            $$props.class,
        )}
        color="none"
        on:click
>
    <slot/>
</Button>
