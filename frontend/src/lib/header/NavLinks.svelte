<script lang="ts">
    import {page} from '$app/stores';
    import {twMerge} from "tailwind-merge";

    export let links: { label: string; url: string }[] = [];

    const ifCurrentUrl = (currentUrl: string, linkUrl: string) => {
        if (linkUrl == "/") {
            return currentUrl === "/";
        } else {
            return currentUrl.startsWith(linkUrl);
        }
    };
</script>

{#each links as link}
    <a
            {...$$restProps}
            href={link.url}
            on:click
            class={twMerge(
                ifCurrentUrl($page.url.pathname, link.url) ? "text-gray-900": "text-gray-400",
                "font-medium hover:text-gray-600",
                $$props.class,
            )}
    >
        {link.label}
    </a>
{/each}
