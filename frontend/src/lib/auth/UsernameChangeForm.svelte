<script lang="ts">
    import {Helper, Input, Label} from "flowbite-svelte";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import {twMerge} from "tailwind-merge";
    import {type FormSubmitResponse, submitUsernameChangeFormWithinPage} from "$lib/auth/authFetch";
    import {goto} from "$app/navigation";
    import {userStore} from "$lib/auth/stores";

    type color = "base" | "red" | "green"
    let usernameColor: color = "base";

    const onSubmit = async (e: Event) => {
        const response: FormSubmitResponse = await submitUsernameChangeFormWithinPage(e);
        const errorFields = response.errorFields;

        if (errorFields.includes("username")) {
            usernameColor = "red";
        } else {
            usernameColor = "base";
        }
    };
</script>

<form class={twMerge("space-y-4", $$props.style)} on:submit={onSubmit}>
    <Label class="space-y-2">
        <span>Current: {$userStore?.username}</span>
        <Input color={usernameColor} name="username" placeholder="SuperCoolName" required type="text"/>
        <Helper class="text-xs mt-2">
            Your username must be unique. 1-30 characters. Letters, digits and @.+-_ only.
        </Helper>
    </Label>
    <PrimaryButton class="w-full" size="sm" type="submit">Change</PrimaryButton>
</form>
