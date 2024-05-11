<script lang="ts">
	import UsernameChangeForm from '$lib/auth/UsernameChangeForm.svelte';
	import Title from '$lib/section/Title.svelte';
	import PrimaryButton from '$lib/button/PrimaryButton.svelte';
	import { twMerge } from 'tailwind-merge';
	import Modal from '$lib/modal/Modal.svelte';
	import {
		submitActivationLinkRequestLoggedInWithinPage,
		submitPasswordResetRequestLoggedInWithinPage
	} from '$lib/auth/authFetchClient';
	import type { User } from '$lib/types';
	import { deleteAccountLoggedInWithinPage } from '$lib/auth/authFetchClient.js';
	import OutlineButton from '$lib/button/OutlineButton.svelte';
	import CriticalButton from '$lib/button/CriticalButton.svelte';

	export let data: { token: string | null; expiry: string | null; user: User | null };

	let openDeleteConfirmationModal: boolean = false;

	const onDelete = async (e: Event) => {
		const ok = await deleteAccountLoggedInWithinPage(e, data.token, data.expiry);
		if (ok) {
			openDeleteConfirmationModal = false;
		}
	};
</script>

{#if data.user}
	<div class="space-y-8">
		{#if !data.user?.activated}
			<div class="space-y-4">
				<Title>Activate account</Title>
				<p class="text-sm">
					Your account is not activated. Please activate your account to access all features.
				</p>
				<form
					class={twMerge('space-y-4', $$props.style)}
					on:submit={async (e) =>
						await submitActivationLinkRequestLoggedInWithinPage(e, data.token, data.expiry)}
				>
					<PrimaryButton class="w-full sm:max-w-72" size="sm" type="submit"
						>Request a new activation link
					</PrimaryButton>
				</form>
			</div>
		{/if}

		<div class="space-y-4 sm:max-w-72">
			<Title>Change username</Title>
			<UsernameChangeForm token={data.token} expiry={data.expiry} user={data.user} />
		</div>

		<div class="space-y-4">
			<Title>Reset password</Title>
			<p class="text-sm">If you forgot your password, you can request a password reset link.</p>
			<div class="space-y-4 sm:max-w-72">
				<form
					class={twMerge('space-y-4', $$props.style)}
					on:submit={async (e) =>
						await submitPasswordResetRequestLoggedInWithinPage(e, data.token, data.expiry)}
				>
					<PrimaryButton class="w-full" size="sm" type="submit">Request a reset link</PrimaryButton>
				</form>
			</div>
		</div>

		<div class="space-y-4">
			<Title>Delete account</Title>
			<p class="text-sm">
				This action is irreversible. All data associated with your account will be deleted.
			</p>
			<div class="space-y-4 sm:max-w-72">
				<CriticalButton
					on:click={() => (openDeleteConfirmationModal = true)}
					class="w-full"
					size="sm"
				>
					Delete account
				</CriticalButton>
			</div>
		</div>
	</div>

	{#if openDeleteConfirmationModal}
		<Modal bind:open={openDeleteConfirmationModal} class="w-96">
			<form class="space-y-8" on:submit={onDelete}>
				<h3 class="text-xl font-medium dark:text-white">Confirm deletion</h3>
				<p class="text-gray-600 dark:text-gray-400">
					Are you sure you want to delete your account? This action is irreversible.
				</p>
				<div class="grid grid-cols-2 gap-x-4">
					<OutlineButton on:click={() => (openDeleteConfirmationModal = false)}>
						Cancel
					</OutlineButton>
					<CriticalButton type="submit">Delete</CriticalButton>
				</div>
			</form>
		</Modal>
	{/if}
{/if}
