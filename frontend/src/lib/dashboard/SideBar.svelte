<script>
    import {page} from '$app/stores';
    import {Sidebar, SidebarGroup, SidebarItem, SidebarWrapper} from 'flowbite-svelte';
    import {ThumbsDownSolid, ThumbsUpSolid, UserSettingsSolid, UserSolid} from 'flowbite-svelte-icons';

    $: activeUrl = $page.url.pathname;

    let activeClass = 'flex items-center p-2 text-base font-normal text-gray-900 bg-gray-200 dark:bg-gray-700 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700';
    let nonActiveClass = 'flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700';
</script>

<Sidebar {activeClass} {activeUrl} {nonActiveClass}>
    <SidebarWrapper class="bg-white">
        <SidebarGroup>
            <SidebarItem class="font-semibold" href="/account/me" label="Profile" on:click>
                <svelte:fragment slot="icon">
                    <UserSolid
                            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"/>
                </svelte:fragment>
            </SidebarItem>
            <SidebarItem class="font-semibold" href="/account/me/likes" label="Likes" on:click>
                <svelte:fragment slot="icon">
                    <ThumbsUpSolid
                            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"/>
                </svelte:fragment>
            </SidebarItem>
            <SidebarItem class="font-semibold" href="/account/me/dislikes" label="Dislikes" on:click>
                <svelte:fragment slot="icon">
                    <ThumbsDownSolid
                            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"/>
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>

        <SidebarGroup border>
            <SidebarItem class="font-semibold" href="/account/me/settings" label="Settings" on:click>
                <svelte:fragment slot="icon">
                    <UserSettingsSolid
                            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"/>
                </svelte:fragment>
            </SidebarItem>
        </SidebarGroup>
    </SidebarWrapper>
</Sidebar>