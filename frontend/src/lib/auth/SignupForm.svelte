<script lang="ts">
    import {Helper, Input, Label} from "flowbite-svelte";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import {twMerge} from "tailwind-merge";
    import type {FormSubmitResponse} from "$lib/auth/authFetch";
    import {submitSignUpFormWithinPage} from "$lib/auth/authFetch";
    import {goto} from "$app/navigation";

    export let nextLink: string | null = '';

    type color = "base" | "red" | "green"
    let andrewIDColor: color = "base";
    let usernameColor: color = "base";
    let passwordColor: color = "base";
    let repeatPasswordColor: color = "base";

    const onSubmit = async (e: Event) => {
        const response: FormSubmitResponse = await submitSignUpFormWithinPage(e);
        const errorFields = response.errorFields;

        if (errorFields.includes("andrew_id")) {
            andrewIDColor = "red";
        } else {
            andrewIDColor = "base";
        }

        if (errorFields.includes("username")) {
            usernameColor = "red";
        } else {
            usernameColor = "base";
        }

        if (errorFields.includes("password")) {
            passwordColor = "red";
        } else {
            passwordColor = "base";
        }

        if (errorFields.includes("repeat_password")) {
            repeatPasswordColor = "red";
        } else {
            repeatPasswordColor = "base";
        }

        if (response.success && nextLink !== null) {
            await goto(nextLink);
        }
    };
</script>

<form class={twMerge("space-y-4", $$props.style)} on:submit={onSubmit}>
    <h3 class="text-xl font-medium dark:text-white">Create an account</h3>
    <Label class="space-y-2">
        <span>Andrew ID</span>
        <Input color={andrewIDColor} name="andrew_id" placeholder="andrew123" required type="text"/>
        <Helper class="text-xs mt-2">
            Your Andrew ID will never be displayed or shared.
        </Helper>
    </Label>
    <Label class="space-y-2">
        <span>Username</span>
        <Input color={usernameColor} name="username" placeholder="SuperCoolName" required type="text"/>
        <Helper class="text-xs mt-2">
            Your username must be unique. 1-30 characters. Letters, digits and @.+-_ only.
        </Helper>
    </Label>
    <Label class="space-y-2">
        <span>Password</span>
        <Input color={passwordColor} name="password" placeholder="••••••••" required type="password"/>
        <Helper class="text-xs mt-2">
            Your password must be 8-64 characters and contain at least:
            (1) One digit;
            (2) One lowercase letter;
            (3) One uppercase letter;
            (4) One special character.
        </Helper>
    </Label>
    <Label class="space-y-2">
        <span>Repeat password</span>
        <Input color={repeatPasswordColor} name="repeat_password" placeholder="••••••••" required type="password"/>
    </Label>
    <PrimaryButton class="w-full" size="sm" type="submit">Sign up</PrimaryButton>
</form>