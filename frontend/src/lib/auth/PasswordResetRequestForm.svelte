<script lang="ts">
    import {Input, Label} from "flowbite-svelte";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import {twMerge} from "tailwind-merge";
    import {type FormSubmitResponse, submitPasswordResetRequestFormWithinPage} from "$lib/auth/authFetch";
    import {goto} from "$app/navigation";

    export let nextLink: string | null = '';

    type color = "base" | "red" | "green"
    let andrewIDColor: color = "base";

    const onSubmit = async (e: Event) => {
        const response: FormSubmitResponse = await submitPasswordResetRequestFormWithinPage(e);
        const errorFields = response.errorFields;

        if (errorFields.includes("andrew_id")) {
            andrewIDColor = "red";
        } else {
            andrewIDColor = "base";
        }

        if (response.success && nextLink !== null) {
            await goto(nextLink);
        }
    };
</script>

<form class={twMerge("space-y-4", $$props.style)} on:submit={onSubmit}>
    <h3 class="text-xl font-medium dark:text-white">Reset your password</h3>
    <Label class="space-y-2">
        <span>Andrew ID</span>
        <Input color={andrewIDColor} name="andrew_id" placeholder="andrew123" required type="text"/>
    </Label>
    <PrimaryButton class="w-full" size="sm" type="submit">Request a reset link</PrimaryButton>
</form>
