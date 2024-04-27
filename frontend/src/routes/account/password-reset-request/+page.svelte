<script lang="ts">
    import PasswordResetRequestForm from "$lib/auth/PasswordResetRequestForm.svelte";

    export let data: { nextLink: string };
</script>

<div class="flex-1 flex items-center justify-center">
    <div class="max-w-[320px] w-full">
        <PasswordResetRequestForm nextLink={data.nextLink}/>
    </div>
</div>
