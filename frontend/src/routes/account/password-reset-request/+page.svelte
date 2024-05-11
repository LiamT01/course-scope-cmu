<script lang="ts">
	import PasswordResetRequestForm from '$lib/auth/PasswordResetRequestForm.svelte';

	export let data: { nextLink: string };
</script>

<div class="flex flex-1 items-center justify-center">
	<div class="w-full max-w-[320px]">
		<PasswordResetRequestForm nextLink={data.nextLink} />
	</div>
</div>
