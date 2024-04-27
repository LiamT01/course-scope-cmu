<script lang="ts">
    import {userStore} from "$lib/auth/stores";
    import {format} from "date-fns";
    import {Avatar} from "flowbite-svelte";
    import {CalendarMonthOutline} from "flowbite-svelte-icons";
    import Pagination from "$lib/pagination/Pagination.svelte";
    import RatingCard from "$lib/rating-card/RatingCard.svelte";
    import type {Metadata, Rating, UserStats} from "$lib/types";
    import {goto} from "$app/navigation";
    import OutlineButton from "$lib/button/OutlineButton.svelte";

    export let data: { metadata: Metadata, ratings: Rating[], token: string, expiry: string, userStats: UserStats };

    const previous = () => {
        if (data.metadata.current_page > data.metadata.first_page) {
            goto(`?page=${data.metadata.current_page - 1}&page_size=${data.metadata.page_size}`);
        }
    };
    const next = () => {
        if (data.metadata.current_page < data.metadata.last_page) {
            goto(`?page=${data.metadata.current_page + 1}&page_size=${data.metadata.page_size}`);
        }
    };

    const jump = (e: { detail: { page: number } }) => {
        if (e.detail.page >= data.metadata.first_page &&
            e.detail.page <= data.metadata.last_page &&
            e.detail.page !== data.metadata.current_page) {
            goto(`?page=${e.detail.page}&page_size=${data.metadata.page_size}`);
        }
    };
</script>

{#if $userStore}
    <div class="flex flex-col gap-y-2">
        <div>
            <div class="px-4 pt-4 pb-6 flex flex-col sm:flex-row gap-x-8 gap-y-4 sm:items-center justify-between bg-gradient-to-b from-white via-cyan-100 via-30%">
                <div class="flex gap-x-6 items-center">
                    <Avatar class="ring-2 ring-white w-16 h-16" src="/authenticated-avatar.png"/>
                    <div>
                        <p class="font-medium text-lg">{$userStore.username}</p>
                        <p class="text-sm text-gray-500 flex items-center gap-x-2">
                            <CalendarMonthOutline class="w-4 h-4"/>
                            Joined {format($userStore.created_at, 'MMM d, yyyy')}</p>
                    </div>
                </div>

                <div class="flex-1 pb-2 gap-x-4 sm:py-0 justify-around flex text-center">
                    <div class="flex flex-col items-center">
                        <p>{data.metadata.total}</p>
                        <p class="text-sm text-gray-500">Ratings created</p>
                    </div>
                    <div class="flex flex-col items-center">
                        <p>{data.userStats?.courses_rated ?? 0}</p>
                        <p class="text-sm text-gray-500">Courses rated</p>
                    </div>
                    <div class="flex flex-col items-center">
                        <p>{data.userStats?.likes_received ?? 0}</p>
                        <p class="text-sm text-gray-500">Likes received</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="flex flex-col gap-y-4">
            {#if data.ratings.length > 0}
                {#each data.ratings as rating (rating.id)}
                    <RatingCard {rating} showCourseDetail/>
                {/each}

                <div class="mx-auto">
                    <Pagination metadata={data.metadata} icon on:next={next} on:previous={previous} on:jump={jump}/>
                </div>
            {:else if data.metadata.current_page > 1}
                <OutlineButton on:click={previous}>
                    No ratings on page {data.metadata.current_page}. Click here to go back.
                </OutlineButton>
            {:else}
                <p class="text-center">
                    No ratings yet. Go rate courses you have taken!
                </p>
            {/if}
        </div>
    </div>
{/if}
