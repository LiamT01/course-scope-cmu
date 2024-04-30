<script lang="ts">
    import {clickOutside} from "$lib/actions/clickOutside";
    import {fade} from "svelte/transition";
    import {Dropdown} from "flowbite-svelte";
    import {InfoCircleOutline} from "flowbite-svelte-icons";
    import {twMerge} from "tailwind-merge";
    import GhostButton from "$lib/button/GhostButton.svelte";

    export let rubricName: string;
    export let stats: { label: string, percentage: number }[];
    export let key: string;

    let dropdownOpen = false;

    const handleClickOutside = () => dropdownOpen && (dropdownOpen = false);
</script>

<div class={twMerge("relative", $$props.class)}>
    <GhostButton class="w-10 h-10" id="rating-stats-{key}" size="xs">
        <InfoCircleOutline class="w-4 h-4"/>
    </GhostButton>
    <div use:clickOutside={{callback: handleClickOutside, excludedElement: document.getElementById(`rating-stats-${key}`)}}>
        <Dropdown bind:open={dropdownOpen} class="font-medium p-4 space-y-4 w-64 border rounded-lg" placement="bottom"
                  triggeredBy="#rating-stats-{key}">
            <p class="text-sm">{rubricName.slice(0, 1).toUpperCase() + rubricName.slice(1)} distribution</p>
            <div class="grid grid-cols-[max-content_minmax(0,1fr)_max-content] items-center gap-x-4 gap-y-4 text-xs"
                 transition:fade={{duration: 100}}>
                {#each stats as {label, percentage}}
                    <p class="dark:text-gray-500">{label.slice(0, 1).toUpperCase() + label.slice(1)}</p>
                    <div class="h-5 bg-gray-100 rounded dark:bg-gray-700">
                        <div class="h-5 bg-[#ffd3b6] rounded" style="width: {percentage}%"/>
                    </div>
                    <span class="w-8 dark:text-gray-500">{percentage}%</span>
                {/each}
            </div>
        </Dropdown>
    </div>
</div>
