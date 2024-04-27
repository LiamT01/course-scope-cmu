<script lang="ts">
    import {clickOutside} from '$lib/actions/clickOutside';
    import {fade} from 'svelte/transition';
    import {Dropdown} from 'flowbite-svelte';
    import {
        BookOpenOutline,
        BuildingOutline,
        DotsHorizontalOutline,
        UserOutline
    } from 'flowbite-svelte-icons';
    import {twMerge} from 'tailwind-merge';
    import type {Course, CourseOffering} from '$lib/types';
    import GhostButton from "$lib/button/GhostButton.svelte";

    export let course: Course | null;
    export let courseOffering: CourseOffering | null;
    export let key: string;
    export let showCourseDetail: boolean = true;

    let dropdownOpen = false;

    const handleClickOutside = () => dropdownOpen && (dropdownOpen = false);
</script>

{#if course || courseOffering}
    <div class={twMerge('relative', $$props.class)}>
        <GhostButton class="h-8 w-8" id="course-info-{key}" size="xs">
            <DotsHorizontalOutline class="h-4 w-4"/>
        </GhostButton>

        <div
                use:clickOutside={{
			callback: handleClickOutside,
			excludedElement: document.getElementById(`course-info-${key}`)
		}}
        >
            <Dropdown
                    bind:open={dropdownOpen}
                    class="border rounded-lg w-56 space-y-2 p-4 font-medium"
                    placement="bottom"
                    triggeredBy="#course-info-{key}"
            >
                <div class="flex flex-col gap-y-2" transition:fade={{duration: 100}}>
                    <div class="flex flex-col gap-x-4 gap-y-4 text-xs">
                        {#if course && showCourseDetail}
                            <div class="flex items-center gap-x-2">
                                <BuildingOutline class="h-4 w-4"/>
                                <span>{course.department}</span>
                            </div>
                            <div class="flex items-center gap-x-2">
                                <BookOpenOutline class="h-4 w-4"/>
                                <span>{course.units} units</span>
                            </div>
                        {/if}
                        {#if courseOffering}
                            {#each courseOffering.instructors as instructor (instructor.id)}
                                <div class="flex items-center gap-x-2">
                                    <UserOutline class="h-4 w-4"/>
                                    <a
                                            class="underline decoration-gray-400 underline-offset-4 hover:decoration-gray-900"
                                            href="/instructors/{instructor?.id}"
                                    >
                                        {instructor?.name}
                                    </a>
                                </div>
                            {/each}
                        {/if}
                    </div>
                </div>
            </Dropdown>
        </div>
    </div>
{/if}
