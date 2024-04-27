<script lang="ts">
    import SignupForm from "$lib/auth/SignupForm.svelte";

    export let data: {nextLink: string};
</script>

<div class="flex-1 flex items-center justify-center">
    <div class="max-w-[540px] w-full">
        <SignupForm nextLink={data.nextLink}/>
    </div>
</div>
