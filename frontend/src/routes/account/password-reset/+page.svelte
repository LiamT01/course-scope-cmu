<script lang="ts">
	import PasswordResetForm from '$lib/auth/PasswordResetForm.svelte';

	export let data: { token: string | null };
</script>

<div class="flex flex-1 items-center justify-center">
	<div class="w-full max-w-[320px]">
		<PasswordResetForm token={data.token} />
	</div>
</div>
