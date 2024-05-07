<script>
    import {openLogInModal} from "$lib/modal/stores";
    import {submitLoginFormWithinPage} from "$lib/auth/authFetchClient";
    import {Helper, Input, Label} from "flowbite-svelte";
    import OutlineButton from "$lib/button/OutlineButton.svelte";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import Modal from "$lib/modal/Modal.svelte";
    import {goto} from "$app/navigation";
    import {page} from "$app/stores";

    const onForgotPassword = async () => {
        $openLogInModal = false;
        await goto(`/account/password-reset-request?next=${$page.url.pathname}`);
    }

    const onSingUp = async () => {
        $openLogInModal = false;
        await goto(`/account/signup?next=${$page.url.pathname}`);
    }
</script>

{#if $openLogInModal}
    <Modal bind:open={$openLogInModal}>
        <form on:submit={submitLoginFormWithinPage} class="flex flex-col gap-y-4">
            <h3 class="text-xl font-medium dark:text-white">Log in to CourseScope CMU</h3>
            <Label class="space-y-2">
                <span>Andrew ID</span>
                <Input name="andrew_id" placeholder="andrew123" required type="text"/>
            </Label>
            <Label class="space-y-2">
                <span>Password</span>
                <Input name="password" placeholder="••••••••" required type="password"/>
                <Helper class="text-xs mt-2">
                    This is NOT your Andrew password! Instead, please create a separate password by signing up below if you haven't already.
                </Helper>
            </Label>
            <div class="flex items-start">
                <a class="text-sm text-gray-600 hover:text-gray-900 dark:text-gray-500 dark:hover:text-gray-700"
                   href="javascript:" on:click={onForgotPassword}>
                    Forgot password?
                </a>
            </div>
            <PrimaryButton class="w-full" size="sm" type="submit">Log in</PrimaryButton>
            <OutlineButton class="w-full" size="sm" on:click={onSingUp}>Create an account</OutlineButton>
        </form>
    </Modal>
{/if}
