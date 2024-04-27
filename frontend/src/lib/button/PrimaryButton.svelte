<script>
    import {Button} from "flowbite-svelte";
    import {twMerge} from "tailwind-merge";
</script>

<Button
        {...$$restProps}
        class={twMerge(
            "bg-gray-600 border-gray-300 text-white dark:hover:bg-gray-700 hover:bg-gray-700 rounded-lg",
            $$props.class,
        )}
        color="none"
        on:click
>
    <slot/>
</Button>
