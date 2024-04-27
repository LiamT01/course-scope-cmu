<script lang="ts">
    import type {Course, RatingStats} from "$lib/types";
    import {BookOpenOutline, BuildingOutline,} from "flowbite-svelte-icons";
    import RatingIcons from "$lib/rating/RatingIcons.svelte";
    import round from "$lib/util/round";
    import {onMount} from "svelte";
    import {fetchWithinPage} from "$lib/auth/fetchClient";

    export let course: Course;

    let stats: RatingStats | null;
    onMount(async () => {
        stats = await fetchStats();
    })

    async function fetchStats() {
        const statsResponse = await fetchWithinPage(`/api/ratings/stats?course_id=${course.id}`);
        return await statsResponse.json();
    }
</script>

<div class="border-y py-4 sm:py-8 flex flex-col gap-4" id="course-{course.id}">
    <!--Course info-->
    <div class="flex flex-wrap text-xs gap-x-4 gap-y-1">
        <div class="flex items-center gap-x-2">
            <BuildingOutline class="w-4 h-4"/>
            <span>{course.department}</span>
        </div>
        <div class="flex items-center gap-x-2">
            <BookOpenOutline class="w-4 h-4"/>
            <span>{course.units} units</span>
        </div>
    </div>

    <a class="text-md font-medium" href="/courses/{course.id}">
        <span class="mr-2">{course.number}</span>
        <span>{course.name}</span>
    </a>

    {#if stats}
        <RatingIcons size={20} text="sm" rating={round(stats.avg_overall, 1)}/>
    {/if}

<!--    <div class="leading-relaxed text-sm whitespace-pre-wrap line-clamp-3">-->
<!--        {course.description}-->
<!--    </div>-->
</div>