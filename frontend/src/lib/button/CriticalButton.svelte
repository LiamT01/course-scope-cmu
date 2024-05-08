<script>
    import {Button} from "flowbite-svelte";
    import {twMerge} from "tailwind-merge";
</script>

<Button
        {...$$restProps}
        class={twMerge(
            "bg-red-600 border-red-300 text-white dark:hover:bg-red-700 hover:bg-red-700 rounded-lg",
            $$props.class,
        )}
        color="none"
        on:click
>
    <slot/>
</Button>
