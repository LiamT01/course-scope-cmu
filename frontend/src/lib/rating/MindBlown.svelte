<script lang="ts">
    import generateID from "$lib/util/generateID";

    export let fillPercent: number = 100;
    export let fillColor: string = '#ffd3b6';
    export let strokeColor: string = '#ff8b94';
    export let size: number = 24;
    export let ariaLabel: string = 'heart';
    export let id: string = generateID();
    export let role = 'img';
</script>

<svg {...$$restProps} aria-label={ariaLabel} class={$$props.class} height={size} on:click {role}
     viewBox="0 0 24 24" width={size}>
    <defs>
        <linearGradient {id}>
            {#if fillPercent !== 100}
                <stop offset="0%" stop-color={fillColor}/>
                <stop offset="{fillPercent}%" stop-color={fillColor}/>
                <stop offset="{fillPercent}%" stop-color="transparent"/>
                <stop offset="100%" stop-color="transparent"/>
            {:else}
                <stop offset="0%" stop-color={fillColor}/>
                <stop offset="100%" stop-color={fillColor}/>
            {/if}
        </linearGradient>
    </defs>
    <path d="M13.8 9h2.7a2 2 0 1 0-1-3.7 1.9 1.9 0 0 0-.4-1.4 2 2 0 0 0-3.1-.2 2 2 0 0 0-3.5 1.6 2 2 0 1 0-1 3.7h3.4m3-1v3.6M10.1 9v2.6M15 16h0m-6 0h0m12-4.4a1.8 1.8 0 0 1-1.8 1.3 1.9 1.9 0 0 1-1.8-1.3 1.9 1.9 0 0 1-3.6 0 1.8 1.8 0 0 1-1.7 1.3 2 2 0 0 1-2-1.3 1.8 1.8 0 0 1-1.6 1.3 2 2 0 0 1-2-1.3 1.9 1.9 0 0 1-1.8 1.3A1.8 1.8 0 0 1 3 11.6v.4a9 9 0 0 0 18 0v-.4Z"
          fill="url(#{id})" stroke={strokeColor}
          stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>
    <path d="M12 17a2 2 0 0 1 2 2h-4a2 2 0 0 1 2-2Z" fill={strokeColor}
          stroke-linecap="round" stroke-linejoin="round"/>
</svg>
