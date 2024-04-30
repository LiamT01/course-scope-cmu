<script lang="ts">
    import type {Instructor, Metadata} from "$lib/types";
    import InstructorBriefCard from "$lib/instructor-card/InstructorBriefCard.svelte";
    import Pagination from "$lib/pagination/Pagination.svelte";
    import {goto} from "$app/navigation";
    import Title from "$lib/section/Title.svelte";
    import OutlineButton from "$lib/button/OutlineButton.svelte";
    import makeQueryParams from "$lib/util/makeQueryParams";

    export let data: {
        instructors: Instructor[],
        metadata: Metadata,
        name: string,
        sort: string,
    };

    const previous = () => {
        if (data.metadata.current_page > data.metadata.first_page) {
            const allQuery = makeQueryParams({
                name: data.name,
                sort: data.sort,
                page: data.metadata.current_page - 1,
                pageSize: data.metadata.page_size,
            });
            goto(`?${allQuery}`);
        }
    };
    const next = () => {
        if (data.metadata.current_page < data.metadata.last_page) {
            const allQuery = makeQueryParams({
                name: data.name,
                sort: data.sort,
                page: data.metadata.current_page + 1,
                pageSize: data.metadata.page_size,
            });
            goto(`?${allQuery}`);
        }
    };

    const jump = (e: { detail: { page: number } }) => {
        if (e.detail.page >= data.metadata.first_page &&
            e.detail.page <= data.metadata.last_page &&
            e.detail.page !== data.metadata.current_page) {
            const allQuery = makeQueryParams({
                name: data.name,
                sort: data.sort,
                page: e.detail.page,
                pageSize: data.metadata.page_size,
            });
            goto(`?${allQuery}`);
        }
    };

</script>

<svelte:head>
    <title>All Instructors - CourseScope CMU</title>
    <meta name="description" content="All instructors teaching at Carnegie Mellon University."/>
</svelte:head>

<div class="flex flex-col gap-y-4">
    <Title>All instructors ({data.metadata.total})</Title>
    {#if data.instructors.length > 0}
        {#each data.instructors as instructor (instructor.id)}
            <InstructorBriefCard {instructor}/>
        {/each}

        <div class="mx-auto">
            <Pagination metadata={data.metadata} icon on:next={next} on:previous={previous} on:jump={jump}/>
        </div>
    {:else if data.metadata.current_page > 1}
        <OutlineButton on:click={previous}>
            No instructors on page {data.metadata.current_page}. Click here to go back.
        </OutlineButton>
    {:else}
        <p class="text-sm">
            No instructors found.
        </p>
    {/if}
</div>
