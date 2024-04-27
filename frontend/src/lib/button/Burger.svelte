<script lang="ts">
    import GhostButton from "$lib/button/GhostButton.svelte";
    import {BarsOutline} from "flowbite-svelte-icons";
</script>

<GhostButton {...$$restProps} on:click size="sm">
    <BarsOutline/>
</GhostButton>