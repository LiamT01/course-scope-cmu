<script lang="ts">
    import type {Instructor, RatingStats} from "$lib/types";
    import RatingIcons from "$lib/rating/RatingIcons.svelte";
    import round, {roundFixed} from "$lib/util/round";
    import {onMount} from "svelte";
    import {fetchWithinPage} from "$lib/auth/fetchClient";
    import {apiBaseUrl} from "$lib/constants";

    export let instructor: Instructor;

    let stats: RatingStats | null;

    onMount(async () => {
        stats = await fetchStats();
    });

    async function fetchStats() {
        const statsResponse = await fetchWithinPage(`${apiBaseUrl}/ratings/stats?instructor_ids=${instructor.id}`)
        return await statsResponse.json();
    }
</script>

<div class="py-4 sm:py-8 border-y flex flex-col gap-4" id="instructor-{instructor.id}">
    <p class="text-lg font-medium">
        <a href="/instructors/{instructor.id}">{instructor.name}</a>
    </p>

    {#if stats}
        <RatingIcons size={20} text="sm" rating={round(stats.avg_overall, 1)}/>
        <div class="flex gap-x-8 text-sm font-medium text-gray-600 dark:text-gray-400">
            <div class="grid grid-rows-3 grid-flow-col sm:grid-rows-none sm:grid-cols-3 sm:grid-flow-row gap-y-2 gap-x-6 sm:gap-x-16">
                    <span class="grid grid-cols-[72px_max-content]">
                        <span>Teaching:</span>
                        <span>{roundFixed(stats.avg_teaching, 1)} / 5.0</span>
                    </span>
                <span class="grid grid-cols-[72px_max-content]">
                        <span>Materials:</span>
                        <span>{roundFixed(stats.avg_materials, 1)} / 5.0</span>
                    </span>
                <span class="grid grid-cols-[72px_max-content]">
                        <span>Value:</span>
                        <span>{roundFixed(stats.avg_value, 1)} / 5.0</span>
                    </span>
            </div>
        </div>
    {/if}
</div>