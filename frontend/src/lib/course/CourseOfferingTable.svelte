<script lang="ts">
    import {Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell} from 'flowbite-svelte';
    import type {Offering} from "$lib/types.js";

    export let courseOfferings: Offering[] | null = null;
    export let instructorOfferings: Offering[] | null = null;
</script>

<Table hoverable={true}>
    <TableHead>
        <TableHeadCell>Semester</TableHeadCell>
        <TableHeadCell>Year</TableHeadCell>
        {#if instructorOfferings}
            <TableHeadCell>Number</TableHeadCell>
            <TableHeadCell>Name</TableHeadCell>
        {/if}
        <TableHeadCell>Location</TableHeadCell>
        {#if courseOfferings}
            <TableHeadCell>Instructors</TableHeadCell>
        {/if}
    </TableHead>
    <TableBody>
        {#if courseOfferings}
            {#each courseOfferings as courseOffering (courseOffering.id)}
                <TableBodyRow>
                    <TableBodyCell>{courseOffering.semester}</TableBodyCell>
                    <TableBodyCell>{courseOffering.year}</TableBodyCell>
                    <TableBodyCell>{courseOffering.location}</TableBodyCell>
                    <TableBodyCell>
                        {#each courseOffering.instructors as instructor, instructorIdx (instructor.id)}
                            <a class="underline underline-offset-4 decoration-gray-400 hover:decoration-gray-900"
                               href="/instructors/{instructor.id}">
                                {instructor.name}</a>{instructorIdx < courseOffering.instructors.length - 1 ? "; " : ""}
                        {/each}
                    </TableBodyCell>
                </TableBodyRow>
            {/each}
        {/if}
        {#if instructorOfferings}
            {#each instructorOfferings as instructorOffering (instructorOffering.id)}
                <TableBodyRow>
                    <TableBodyCell>{instructorOffering.semester}</TableBodyCell>
                    <TableBodyCell>{instructorOffering.year}</TableBodyCell>
                    <TableBodyCell>
                        <a class="underline underline-offset-4 decoration-gray-400 hover:decoration-gray-900"
                           href="/courses/{instructorOffering.course.id}">{instructorOffering.course.number}</a>
                    </TableBodyCell>
                    <TableBodyCell>
                        <a class="underline underline-offset-4 decoration-gray-400 hover:decoration-gray-900"
                           href="/courses/{instructorOffering.course.id}">{instructorOffering.course.name}</a>
                    </TableBodyCell>
                    <TableBodyCell>{instructorOffering.location}</TableBodyCell>
                </TableBodyRow>
            {/each}
        {/if}
    </TableBody>
</Table>