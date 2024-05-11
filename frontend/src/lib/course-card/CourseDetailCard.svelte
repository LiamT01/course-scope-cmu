<script lang="ts">
	import type { Course, CourseOffering } from '$lib/types';
	import { BookOpenOutline, BuildingOutline } from 'flowbite-svelte-icons';
	import CourseOfferingTable from '$lib/course/CourseOfferingTable.svelte';

	export let course: Course;
	export let courseOfferings: CourseOffering[];
</script>

<div class="flex flex-col gap-y-4" id="course-{course.id}">
	<div class="flex flex-col gap-y-4">
		<!--Course info-->
		<div class="flex flex-wrap gap-x-4 gap-y-1 text-xs">
			<div class="flex items-center gap-x-2">
				<BuildingOutline class="h-4 w-4" />
				<span>{course.department}</span>
			</div>
			<div class="flex items-center gap-x-2">
				<BookOpenOutline class="h-4 w-4" />
				<span>{course.units} units</span>
			</div>
		</div>

		<p class="text-xl font-medium">
			<span class="mr-2">{course.number}</span>
			<span>{course.name}</span>
		</p>
	</div>

	<div class="flex flex-col gap-y-2 border-y">
		<CourseOfferingTable {courseOfferings} />
	</div>
</div>
