<script lang="ts">
	import { Helper, Input, Label } from 'flowbite-svelte';
	import PrimaryButton from '$lib/button/PrimaryButton.svelte';
	import { twMerge } from 'tailwind-merge';
	import {
		type FormSubmitResponse,
		submitPasswordResetFormWithinPage
	} from '$lib/auth/authFetchClient';
	import { goto } from '$app/navigation';

	export let token: string | null = '';

	type color = 'base' | 'red' | 'green';
	let passwordColor: color = 'base';
	let repeatPasswordColor: color = 'base';

	const onSubmit = async (e: Event) => {
		const response: FormSubmitResponse = await submitPasswordResetFormWithinPage(e, token);
		const errorFields = response.errorFields;

		if (errorFields.includes('password')) {
			passwordColor = 'red';
		} else {
			passwordColor = 'base';
		}

		if (errorFields.includes('repeat_password')) {
			repeatPasswordColor = 'red';
		} else {
			repeatPasswordColor = 'base';
		}

		if (response.success) {
			await goto('/');
		}
	};
</script>

<form class={twMerge('space-y-4', $$props.style)} on:submit={onSubmit}>
	<h3 class="text-xl font-medium dark:text-white">Reset your password</h3>
	<Label class="space-y-2">
		<span>Password</span>
		<Input color={passwordColor} name="password" placeholder="••••••••" required type="password" />
		<Helper class="mt-2 text-xs">
			Your password must be 8-64 characters and contain at least: (1) One digit; (2) One lowercase
			letter; (3) One uppercase letter; (4) One special character.
		</Helper>
	</Label>
	<Label class="space-y-2">
		<span>Repeat password</span>
		<Input
			color={repeatPasswordColor}
			name="repeat_password"
			placeholder="••••••••"
			required
			type="password"
		/>
	</Label>
	<PrimaryButton class="w-full" size="sm" type="submit">Reset</PrimaryButton>
</form>
