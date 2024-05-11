<script lang="ts">
	import PrimaryButton from '$lib/button/PrimaryButton.svelte';
	import { activateAccountWithinPage } from '$lib/auth/authFetchClient';
	import { goto } from '$app/navigation';

	export let data: { token: string | null };

	const onActivate = async () => {
		const ok = await activateAccountWithinPage(data.token);
		if (ok) {
			await goto('/');
		}
	};
</script>

<div class="flex flex-1 items-center justify-center">
	<PrimaryButton on:click={onActivate}>Activate your account</PrimaryButton>
</div>
