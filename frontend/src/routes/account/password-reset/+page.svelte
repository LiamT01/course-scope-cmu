<script lang="ts">
    import PasswordResetForm from "$lib/auth/PasswordResetForm.svelte";

    export let data: { token: string | null };
</script>

<div class="flex-1 flex items-center justify-center">
    <div class="max-w-[320px] w-full">
        <PasswordResetForm token={data.token}/>
    </div>
</div>
