<script lang="ts">
	import { fade } from 'svelte/transition';
	import { Button, Dropdown, DropdownDivider, DropdownItem, Search } from 'flowbite-svelte';
	import { ChevronDownSolid } from 'flowbite-svelte-icons';
	import { twMerge } from 'tailwind-merge';
	import { clickOutside } from '$lib/actions/clickOutside';

	export let searchItems: { label: string }[] = [];
	export let value: string = '';
	export let selectIndex: number = 0;
	export let key: string;
	export let placeholder: string;

	let dropdownOpen = false;

	const handleSelectCategory = (index: number) => {
		selectIndex = index;
		dropdownOpen = false;
	};

	const handleClickOutside = () => dropdownOpen && (dropdownOpen = false);
</script>

<form {...$$restProps} on:submit>
	<div class="relative">
		<Button
			class="h-10 w-28 whitespace-nowrap rounded-e-none border border-e-0 border-gray-300"
			color="light"
			id="category-menu-{key}"
		>
			{searchItems[selectIndex].label}
			<ChevronDownSolid class="ms-1.5 h-2 w-2" />
		</Button>
		<div
			use:clickOutside={{
				callback: handleClickOutside,
				excludedElement: document.getElementById(`category-menu-${key}`)
			}}
		>
			<Dropdown
				bind:open={dropdownOpen}
				class="rounded-lg border"
				placement="bottom"
				triggeredBy="#category-menu-{key}"
			>
				<div transition:fade={{ duration: 100 }}>
					{#each searchItems as { label }, index}
						<DropdownItem
							on:click={() => handleSelectCategory(index)}
							class={twMerge(
								selectIndex === index ? 'text-gray-900' : 'text-gray-400',
								'hover:text-gray-600'
							)}
						>
							{label}
						</DropdownItem>

						{#if index < searchItems.length - 1}
							<DropdownDivider />
						{/if}
					{/each}
				</div>
			</Dropdown>
		</div>
	</div>

	<Search bind:value class="h-10 rounded-s-none" on:input {placeholder} size="md" />
</form>
