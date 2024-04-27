<script lang="ts">
    import type {Course, CourseOffering} from "$lib/types";
    import {BookOpenOutline, BuildingOutline} from "flowbite-svelte-icons";
    import CourseOfferingTable from "$lib/course/CourseOfferingTable.svelte";
    import OutlineButton from "$lib/button/OutlineButton.svelte";

    export let course: Course;
    export let courseOfferings: CourseOffering[];
    let showMore: boolean = false;

    const toggleShowMore = () => {
        showMore = !showMore;
    }
</script>

<div class="flex flex-col gap-y-4" id="course-{course.id}">
    <div class="flex flex-col gap-y-4">
        <!--Course info-->
        <div class="flex flex-wrap text-xs   gap-x-4 gap-y-1">
            <div class="flex items-center gap-x-2">
                <BuildingOutline class="w-4 h-4"/>
                <span>{course.department}</span>
            </div>
            <div class="flex items-center gap-x-2">
                <BookOpenOutline class="w-4 h-4"/>
                <span>{course.units} units</span>
            </div>
        </div>

        <p class="text-xl font-medium">
            <span class="mr-2">{course.number}</span>
            <span>{course.name}</span>
        </p>

<!--        <div class="leading-relaxed text-sm whitespace-pre-wrap {showMore ? '' : 'line-clamp-6'}">-->
<!--            {course.description}-->
<!--        </div>-->
<!--        {#if showMore}-->
<!--            <OutlineButton size="sm" on:click={toggleShowMore}>Show less</OutlineButton>-->
<!--        {:else}-->
<!--            <OutlineButton size="sm" on:click={toggleShowMore}>Show more</OutlineButton>-->
<!--        {/if}-->
    </div>

    <div class="flex flex-col gap-y-2 border-y">
        <CourseOfferingTable {courseOfferings}/>
    </div>
</div>
