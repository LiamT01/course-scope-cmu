<script lang="ts">
    import {EnvelopeOutline, GithubSolid} from "flowbite-svelte-icons";

    let year = new Date().getFullYear();
</script>

<div class="bg-gray-50 mt-4 p-8 border-t border-gray-300">
    <div class="flex flex-col items-center gap-y-4 text-gray-600">
        <p class="text-center">Copyright © {year} CourseScope CMU. All rights reserved.</p>
        <p class="flex justify-center items-center gap-x-4 whitespace-nowrap dark:text-white">
            <a href="mailto:course.scope.cmu@gmail.com">
                <EnvelopeOutline/>
            </a>
            <a href="https://github.com/liamt01/course-scope-cmu" target="_blank">
                <GithubSolid/>
            </a>
        </p>
    </div>
</div>