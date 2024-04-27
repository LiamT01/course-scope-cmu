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
    <path d="M12 6C6.5 1 1 8 5.8 13l6.2 7 6.2-7C23 8 17.5 1 12 6Z"
          fill="url(#{id})" stroke={strokeColor}
          stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>
</svg>
