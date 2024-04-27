<script>
    import {CloseButton} from 'flowbite-svelte';
    import Drawer from "$lib/dashboard/Drawer.svelte";
    import {sineIn} from 'svelte/easing';
    import UserSideBar from "$lib/dashboard/SideBar.svelte";
    import GhostButton from "$lib/button/GhostButton.svelte";
    import {AdjustmentsHorizontalOutline} from "flowbite-svelte-icons";

    let hidden = true;
    let spanClass = 'flex-1 ms-3 whitespace-nowrap';
    let transitionParams = {
        x: -320,
        duration: 200,
        easing: sineIn
    };
</script>

<div class="text-center">
    <GhostButton class="w-10 h-10" on:click={() => (hidden = false)}>
        <AdjustmentsHorizontalOutline class="w-5 h-5"/>
    </GhostButton>
</div>

<Drawer bgOpacity="bg-opacity-50" bind:hidden={hidden} class="w-[288px]" id="drawer-sidebar" {transitionParams}
        transitionType="fly">
    <div class="flex items-center">
        <CloseButton class="mb-4 dark:text-white" on:click={() => (hidden = true)}/>
    </div>
    <UserSideBar on:click={() => (hidden = true)}/>
</Drawer>
