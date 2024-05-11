<script lang="ts">
	import SignupForm from '$lib/auth/SignupForm.svelte';

	export let data: { nextLink: string };
</script>

<div class="flex flex-1 items-center justify-center">
	<div class="w-full max-w-[540px]">
		<SignupForm nextLink={data.nextLink} />
	</div>
</div>
