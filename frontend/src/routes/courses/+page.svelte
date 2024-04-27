<script lang="ts">
    import CourseBriefCard from "$lib/course-card/CourseBriefCard.svelte";
    import type {Course, Metadata} from "$lib/types";
    import Pagination from "$lib/pagination/Pagination.svelte";
    import {goto} from "$app/navigation";
    import Title from "$lib/section/Title.svelte";
    import OutlineButton from "$lib/button/OutlineButton.svelte";

    export let data: { courses: Course[], metadata: Metadata, token: string, expiry: string };

    const previous = () => {
        if (data.metadata.current_page > data.metadata.first_page) {
            goto(`?page=${data.metadata.current_page - 1}&page_size=${data.metadata.page_size}`);
        }
    };
    const next = () => {
        if (data.metadata.current_page < data.metadata.last_page) {
            goto(`?page=${data.metadata.current_page + 1}&page_size=${data.metadata.page_size}`);
        }
    };

    const jump = (e: { detail: { page: number } }) => {
        if (e.detail.page >= data.metadata.first_page &&
            e.detail.page <= data.metadata.last_page &&
            e.detail.page !== data.metadata.current_page) {
            goto(`?page=${e.detail.page}&page_size=${data.metadata.page_size}`);
        }
    };

</script>

<div class="flex flex-col gap-y-4">
    <Title>All courses ({data.metadata.total})</Title>
    {#if data.courses.length > 0}
        {#each data.courses as course (course.id)}
            <CourseBriefCard course={course}/>
        {/each}

        <div class="mx-auto">
            <Pagination metadata={data.metadata} icon on:next={next} on:previous={previous} on:jump={jump}/>
        </div>
    {:else if data.metadata.current_page > 1}
        <OutlineButton on:click={previous}>
            No courses on page {data.metadata.current_page}. Click here to go back.
        </OutlineButton>
    {:else}
        <p class="text-sm">
            No courses found.
        </p>
    {/if}
</div>
