<script lang="ts">
	import generateID from '$lib/util/generateID';

	export let fillPercent: number = 100;
	export let fillColor: string = '#ffd3b6';
	export let strokeColor: string = '#ff8b94';
	export let size: number = 24;
	export let ariaLabel: string = 'heart';
	export let id: string = generateID();
	export let role = 'img';
</script>

<svg
	{...$$restProps}
	aria-label={ariaLabel}
	class={$$props.class}
	height={size}
	on:click
	{role}
	viewBox="0 0 24 24"
	width={size}
>
	<defs>
		<linearGradient {id}>
			{#if fillPercent !== 100}
				<stop offset="0%" stop-color={fillColor} />
				<stop offset="{fillPercent}%" stop-color={fillColor} />
				<stop offset="{fillPercent}%" stop-color="transparent" />
				<stop offset="100%" stop-color="transparent" />
			{:else}
				<stop offset="0%" stop-color={fillColor} />
				<stop offset="100%" stop-color={fillColor} />
			{/if}
		</linearGradient>
	</defs>
	<path
		d="M11 5.1a1 1 0 0 1 2 0l1.7 4c.1.4.4.6.8.6l4.5.4a1 1 0 0 1 .5 1.7l-3.3 2.8a1 1 0 0 0-.3 1l1 4a1 1 0 0 1-1.5 1.2l-3.9-2.3a1 1 0 0 0-1 0l-4 2.3a1 1 0 0 1-1.4-1.1l1-4.1c.1-.4 0-.8-.3-1l-3.3-2.8a1 1 0 0 1 .5-1.7l4.5-.4c.4 0 .7-.2.8-.6l1.8-4Z"
		fill="url(#{id})"
		stroke={strokeColor}
		stroke-linecap="round"
		stroke-linejoin="round"
		stroke-width="1.5"
	/>
</svg>
