<script lang="ts">
    import type {Instructor} from "$lib/types";
    import InstructorDetailCard from "$lib/instructor-card/InstructorDetailCard.svelte";

    export let data: { instructor: Instructor }
</script>

<InstructorDetailCard instructor={data.instructor}/>
