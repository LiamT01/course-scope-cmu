<script lang="ts">
	import { EnvelopeOutline, GithubSolid } from 'flowbite-svelte-icons';

	let year = new Date().getFullYear();
</script>

<div class="mt-4 border-t border-gray-300 bg-gray-50 p-8">
	<div class="flex flex-col items-center gap-y-4 text-gray-600">
		<p class="text-center">{year} CourseScope CMU</p>
		<p class="flex items-center justify-center gap-x-8 whitespace-nowrap dark:text-white">
			<a href="mailto:course.scope.cmu@gmail.com">
				<EnvelopeOutline class="h-6 w-6" />
			</a>
			<a href="https://github.com/liamt01/course-scope-cmu" target="_blank">
				<GithubSolid class="h-6 w-6" />
			</a>
		</p>
	</div>
</div>
