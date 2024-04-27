<script lang="ts">
    import {twMerge} from "tailwind-merge";
    import GhostButton from "$lib/button/GhostButton.svelte";
</script>

<GhostButton
        {...$$restProps}
        class={twMerge(
            "border bg-gray-50 border-gray-300",
            $$props.class,
        )}
        color="none"
        on:click
>
    <slot/>
</GhostButton>
