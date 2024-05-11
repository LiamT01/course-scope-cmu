<script lang="ts">
	import {
		Avatar,
		Dropdown,
		DropdownDivider,
		DropdownHeader,
		DropdownItem,
		Navbar,
		NavBrand
	} from 'flowbite-svelte';
	import { fade, slide } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { clickOutside } from '$lib/actions/clickOutside';
	import Burger from '$lib/button/Burger.svelte';
	import NavLinks from '$lib/header/NavLinks.svelte';
	import SearchInput from '$lib/search/SearchInput.svelte';
	import { openLogInModal } from '../modal/stores';
	import { logOutWithinPage } from '$lib/auth/authFetchClient';
	import {
		ArrowLeftToBracketOutline,
		ArrowRightToBracketOutline,
		ChartOutline
	} from 'flowbite-svelte-icons';
	import type { User } from '$lib/types';

	export let token: string | null;
	export let expiry: string | null;
	export let user: User | null;

	const links: { label: string; url: string }[] = [
		{ label: 'Home', url: '/' },
		{ label: 'Courses', url: '/courses' },
		{ label: 'Instructors', url: '/instructors' }
	];

	const searchItems: { label: string; value: string }[] = [
		{ label: 'Courses', value: 'courses' },
		{ label: 'Instructors', value: 'instructors' }
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
	};

	const formatCourseNumber = (query: string) => {
		// Check for 5 consecutive digits not necessarily surrounded by word boundaries
		const regex = /(\d{2})(\d{3})/;
		return query.replace(regex, '$1-$2');
	};

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
		await logOutWithinPage(token, expiry);
	};
</script>

<Navbar class="mb-4 flex w-full border-b border-gray-200 px-2 py-2.5 sm:px-4" let:NavContainer>
	<NavContainer class="gap-y-1">
		<NavBrand href="/">
			<img alt="CourseScope CMU Logo" class="mx-3 h-9" src="/logo.png" />
			<span class="self-center whitespace-nowrap text-lg font-medium dark:text-white"
				>CourseScope CMU</span
			>
		</NavBrand>

		<SearchInput
			bind:selectIndex={searchCategoryIndex}
			bind:value={searchQuery}
			class="mx-12 hidden flex-1 lg:flex"
			key="desktop"
			on:input={onSearchInput}
			on:submit={onSubmitSearch}
			{placeholder}
			{searchItems}
		/>

		<div class="flex items-center md:order-last">
			<Avatar
				class="mx-4 my-1 h-8 w-8 cursor-pointer md:my-0.5 md:ml-12 md:h-9 md:w-9"
				id="avatar-menu"
				src={user ? '/authenticated-avatar.png' : '/anonymous-avatar.png'}
			/>
			<Burger class="h-10 w-10 md:hidden" id="navbar-burger" on:click={toggleNav} />
		</div>

		<Dropdown
			bind:open={openDropdown}
			class="min-w-36 rounded-lg border"
			placement="bottom"
			triggeredBy="#avatar-menu"
		>
			{#if user}
				<div transition:fade={{ duration: 100 }}>
					<DropdownHeader>
						<span class="block text-sm">{user.username}</span>
					</DropdownHeader>
					<DropdownItem class="flex items-center gap-x-4" href="/account/me">
						<ChartOutline class="h-4 w-4" />
						Dashboard
					</DropdownItem>
					<DropdownDivider />
					<DropdownItem class="flex items-center gap-x-4" on:click={logOut}>
						<ArrowRightToBracketOutline class="h-4 w-4" />
						Log out
					</DropdownItem>
				</div>
			{:else}
				<div transition:fade={{ duration: 100 }}>
					<DropdownItem class="flex items-center gap-x-4" on:click={openLogIn}>
						<ArrowLeftToBracketOutline class="h-4 w-4" />
						Log in
					</DropdownItem>
				</div>
			{/if}
		</Dropdown>

		{#if !navHidden}
			<div
				use:clickOutside={{
					callback: handleClickOutside,
					excludedElement: document.getElementById('navbar-burger')
				}}
				transition:slide={{ duration: 200 }}
				class="order-last m-2 flex w-full flex-col gap-y-2 rounded-xl border border-gray-100 bg-white p-4 md:hidden"
			>
				<NavLinks class="rounded-lg p-2 hover:bg-gray-100" {links} on:click={toggleNav} />
			</div>
		{/if}

		<div class="hidden md:flex md:gap-x-12">
			<NavLinks {links} />
		</div>
	</NavContainer>
</Navbar>

<div class="container mx-auto mb-4 px-4 lg:hidden lg:max-w-[1024px]">
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
