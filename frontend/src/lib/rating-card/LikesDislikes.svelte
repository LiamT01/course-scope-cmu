<script lang="ts">
    import {ThumbsDownOutline, ThumbsDownSolid, ThumbsUpOutline, ThumbsUpSolid} from "flowbite-svelte-icons";
    import GhostButton from "$lib/button/GhostButton.svelte";
    import {twMerge} from "tailwind-merge";

    export let netLikes: number;
    export let likedByUser: boolean;
    export let dislikedByUser: boolean;
    export let onLike: () => void;
    export let onDislike: () => void;
</script>

<div class={twMerge("text-sm flex gap-x-2 items-center", $$props.class)}>
    <GhostButton class="w-10 h-10" on:click={onLike} size="xs">
        {#if likedByUser}
            <ThumbsUpSolid class="w-4 h-4"/>
        {:else}
            <ThumbsUpOutline class="w-4 h-4"/>
        {/if}
    </GhostButton>
    {netLikes}
    <GhostButton class="ml-4 w-10 h-10" on:click={onDislike} size="xs">
        {#if dislikedByUser}
            <ThumbsDownSolid class="w-4 h-4"/>
        {:else}
            <ThumbsDownOutline class="w-4 h-4"/>
        {/if}
    </GhostButton>
</div>
