<script lang="ts">
	import { twMerge } from 'tailwind-merge';
	import { Markdown } from 'svelte-exmarkdown';

	export let text: string;
</script>

<div
	class={twMerge(
		'prose prose-sm max-w-none space-y-2 leading-relaxed  text-gray-900 prose-headings:m-0 prose-h1:text-xl prose-h2:text-lg prose-h3:text-base prose-h4:text-sm prose-p:m-0 prose-a:m-0 prose-blockquote:m-0 prose-figure:m-0 prose-figcaption:m-0 prose-code:m-0 prose-pre:m-0 prose-ol:m-0 prose-ul:m-0 prose-li:m-0 prose-table:m-0 prose-thead:m-0 prose-tr:m-0 prose-th:m-0 prose-td:m-0 prose-img:m-0 prose-video:m-0 prose-hr:m-0 prose-lead:m-0',
		$$props.class
	)}
>
	<Markdown md={text} />
</div>
