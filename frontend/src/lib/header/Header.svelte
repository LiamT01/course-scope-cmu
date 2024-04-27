<script lang="ts">
    import {Avatar, Dropdown, DropdownDivider, DropdownHeader, DropdownItem, Navbar, NavBrand} from 'flowbite-svelte';
    import {fade, slide} from "svelte/transition";
    import {goto} from "$app/navigation";
    import {clickOutside} from "$lib/actions/clickOutside";
    import Burger from "$lib/button/Burger.svelte";
    import NavLinks from "$lib/header/NavLinks.svelte";
    import SearchInput from "$lib/search/SearchInput.svelte";
    import {openLogInModal} from "../modal/stores";
    import {userStore} from "$lib/auth/stores";
    import {logOutWithinPage} from "$lib/auth/authFetch";
    import {ArrowLeftToBracketOutline, ArrowRightToBracketOutline, ChartOutline} from "flowbite-svelte-icons";

    const links: { label: string, url: string }[] = [
        {label: 'Home', url: '/'},
        {label: 'Courses', url: '/courses'},
        {label: "Instructors", url: '/instructors'},
    ];

    const searchItems: { label: string, value: string }[] = [
        {label: 'Courses', value: 'courses'},
        {label: 'Instructors', value: 'instructors'},
    ];


    let navHidden: boolean = true;
    let searchCategoryIndex: number = 0;
    $: searchCategory = searchItems[searchCategoryIndex].value;
    let searchQuery: string = '';

    $: placeholder = searchCategory === 'courses' ? '12-345 Course Name' : 'Last Name, First Name';

    let openDropdown: boolean = false;

    const openLogIn = () => {
        openDropdown = false;
        $openLogInModal = true;
    }

    const formatCourseNumber = (query: string) => {
        // Check for 5 consecutive digits not necessarily surrounded by word boundaries
        const regex = /(\d{2})(\d{3})/;
        return query.replace(regex, '$1-$2');
    }

    const onSearchInput = () => {
        if (searchCategory === 'courses') {
            searchQuery = formatCourseNumber(searchQuery);
        }
    };

    const onSubmitSearch = (e: SubmitEvent) => {
        e.preventDefault();
        if (searchCategory === 'courses') {
            // Extract dd-ddd from the query
            const courseNumber: string = searchQuery?.match(/\d{2}-\d{3}/)?.at(0) ?? '';

            // Extract the rest of the query
            const courseName: string = searchQuery?.replace(courseNumber, '').trim();

            const numberQuery = courseNumber.length > 0 ? `number=${courseNumber}` : '';
            const nameQuery = courseName.length > 0 ? `name=${courseName}` : '';
            const query = [numberQuery, nameQuery].filter(Boolean).join('&');

            goto(`/courses?${query}`);
        } else {
            const nameQuery = searchQuery.length > 0 ? `name=${searchQuery}` : '';
            goto(`/instructors?${nameQuery}`);
        }
    };

    const toggleNav = () => {
        navHidden = !navHidden;
    };

    const handleClickOutside = () => !navHidden && (navHidden = true);

    const logOut = async () => {
        openDropdown = false;
        await logOutWithinPage();
    }
</script>

<Navbar class="flex px-2 border-b border-gray-200 sm:px-4 py-2.5 w-full mb-4"
        let:NavContainer>
    <NavContainer class="gap-y-1">
        <NavBrand href="/">
            <img alt="CourseScope CMU Logo" class="mx-3 h-9" src="/logo.png"/>
            <span class="self-center whitespace-nowrap text-lg font-medium dark:text-white">CourseScope CMU</span>
        </NavBrand>

        <SearchInput
                bind:selectIndex={searchCategoryIndex}
                bind:value={searchQuery}
                class="hidden lg:flex flex-1 mx-12"
                key="desktop"
                on:input={onSearchInput}
                on:submit={onSubmitSearch}
                {placeholder}
                {searchItems}
        />

        <div class="flex items-center md:order-last">
            <Avatar class="mx-4 md:ml-12 cursor-pointer my-1 md:my-0.5 w-8 md:w-9 h-8 md:h-9" id="avatar-menu"
                    src={$userStore? "/authenticated-avatar.png" : "/anonymous-avatar.png"}/>
            <Burger
                    class="md:hidden w-10 h-10"
                    id="navbar-burger"
                    on:click={toggleNav}
            />
        </div>

        <Dropdown bind:open={openDropdown} class="border rounded-lg min-w-36" placement="bottom"
                  triggeredBy="#avatar-menu">
            {#if $userStore}
                <div transition:fade={{duration: 100}}>
                    <DropdownHeader>
                        <span class="block text-sm">{$userStore.username}</span>
                    </DropdownHeader>
                    <DropdownItem class="flex gap-x-4 items-center" href="/account/me">
                        <ChartOutline class="w-4 h-4"/>
                        Dashboard
                    </DropdownItem>
                    <DropdownDivider/>
                    <DropdownItem class="flex gap-x-4 items-center" on:click={logOut}>
                        <ArrowRightToBracketOutline class="w-4 h-4"/>
                        Log out
                    </DropdownItem>
                </div>
            {:else}
                <div transition:fade={{duration: 100}}>
                    <DropdownItem class="flex gap-x-4 items-center" on:click={openLogIn}>
                        <ArrowLeftToBracketOutline class="w-4 h-4"/>
                        Log in
                    </DropdownItem>
                </div>
            {/if}
        </Dropdown>

        {#if !navHidden}
            <div
                    use:clickOutside={{
                            callback: handleClickOutside,
                            excludedElement: document.getElementById("navbar-burger"),
                        }}
                    transition:slide={{duration: 200}}
                    class="md:hidden bg-white flex flex-col gap-y-2 m-2 border border-gray-100 rounded-xl p-4 w-full order-last"
            >
                <NavLinks class="hover:bg-gray-100 rounded-lg p-2" links={links} on:click={toggleNav}/>
            </div>
        {/if}

        <div class="hidden md:flex md:gap-x-12">
            <NavLinks links={links}/>
        </div>
    </NavContainer>
</Navbar>

<div class="lg:hidden container mx-auto lg:max-w-[1024px] px-4 mb-4">
    <SearchInput
            bind:selectIndex={searchCategoryIndex}
            bind:value={searchQuery}
            class="flex"
            key="mobile"
            on:input={onSearchInput}
            on:submit={onSubmitSearch}
            {placeholder}
            {searchItems}
    />
</div>
