<script lang="ts">
	import UserSideBar from '$lib/dashboard/SideBar.svelte';
	import { page } from '$app/stores';
	import DrawerSideBar from '$lib/dashboard/DrawerSideBar.svelte';

	const titles: { [key: string]: string } = {
		'/account/me': 'Profile',
		'/account/me/likes': 'Likes',
		'/account/me/dislikes': 'Dislikes',
		'/account/me/settings': 'Settings'
	};
</script>

<div class="grid grid-cols-1 gap-x-4 gap-y-4 lg:grid-cols-[258px_minmax(0,1fr)]">
	<div class="flex items-center gap-x-2 lg:hidden">
		<DrawerSideBar />
		<span class="font-semibold">{titles[$page.url.pathname]}</span>
	</div>
	<aside class="hidden lg:block">
		<div class="lg:sticky lg:top-0 lg:z-20">
			<UserSideBar />
		</div>
	</aside>

	<main class="lg:pt-4">
		<slot />
	</main>
</div>
