<script lang="ts">
	import type { Instructor } from '$lib/types';
	import InstructorDetailCard from '$lib/instructor-card/InstructorDetailCard.svelte';

	export let data: { instructor: Instructor };
</script>

<svelte:head>
	<title>{data.instructor.name} - CourseScope CMU</title>
	<meta
		name="description"
		content={`Check out ratings for course offerings taught by ${data.instructor.name} at Carnegie Mellon University. Get insights into course quality, instructor effectiveness, and more.`}
	/>
	<link rel="canonical" href={`https://course-scope-cmu.com/instructors/${data.instructor.id}`} />
</svelte:head>

<InstructorDetailCard instructor={data.instructor} />
