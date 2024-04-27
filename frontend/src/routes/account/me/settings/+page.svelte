<script lang="ts">
    import {userStore} from "$lib/auth/stores.js";
    import UsernameChangeForm from "$lib/auth/UsernameChangeForm.svelte";
    import Title from "$lib/section/Title.svelte";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import {twMerge} from "tailwind-merge";
    import {
        submitActivationLinkRequestLoggedInWithinPage,
        submitPasswordResetRequestLoggedInWithinPage
    } from "$lib/auth/authFetch";
</script>

{#if $userStore}
    <div class="space-y-8">
        {#if !$userStore?.activated}
            <div class="space-y-4">
                <Title>Activate account</Title>
                <p class="text-sm">Your account is not activated. Please activate your account to access
                    all features.</p>
                <form class={twMerge("space-y-4", $$props.style)}
                      on:submit={submitActivationLinkRequestLoggedInWithinPage}>
                    <PrimaryButton class="w-full sm:max-w-72" size="sm" type="submit">Request a new activation link
                    </PrimaryButton>
                </form>
            </div>
        {/if}

        <div class="space-y-4 sm:max-w-72">
            <Title>Change username</Title>
            <UsernameChangeForm/>
        </div>

        <div class="space-y-4">
            <Title>Reset password</Title>
            <p class="text-sm">If you forgot your password, you can request a password reset link.</p>
            <div class="space-y-4 sm:max-w-72">
                <form class={twMerge("space-y-4", $$props.style)}
                      on:submit={submitPasswordResetRequestLoggedInWithinPage}>
                    <PrimaryButton class="w-full" size="sm" type="submit">Request a reset link</PrimaryButton>
                </form>
            </div>
        </div>
    </div>
{/if}
