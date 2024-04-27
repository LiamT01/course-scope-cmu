<script lang="ts">
    import type {Rating} from '$lib/types';
    import {Avatar} from 'flowbite-svelte';
    import {format, formatDistance, isToday} from 'date-fns';
    import PrimaryButton from '$lib/button/PrimaryButton.svelte';
    import LikesDislikes from '$lib/rating-card/LikesDislikes.svelte';
    import RatingIcons from '$lib/rating/RatingIcons.svelte';

    export let rating: Rating;
    export let onLike: () => void = () => {
    };
    export let onDislike: () => void = () => {
    };
    export let onReportAbuse: () => void;

    const formatDate = (date: Date) => {
        if (isToday(date)) {
            return formatDistance(date, new Date(), {addSuffix: true});
        } else {
            return format(date, 'yyyy-MM-dd HH:mm');
        }
    };
</script>

<div class="flex flex-col gap-x-4 gap-y-2 rounded-lg border p-4 shadow" id="rating-{rating.id}">
    <div class="flex flex-col gap-x-8 gap-y-2 sm:flex-row sm:items-center">
        <!--Author avatar-->
        <div class="flex items-center gap-x-2">
            <Avatar src={rating.avatarURL}/>
            <div class="text-sm">
                <p>{rating.username}</p>
                <time class="text-xs" dateTime={rating.updatedAt.toString()}>
                    {formatDate(rating.updatedAt)}
                </time>
            </div>
        </div>
        <RatingIcons rating={rating.rubric.overall}/>
    </div>

    <!--Rating Rubric-->
    <div class="flex gap-x-8 border-y py-2 text-sm font-medium text-gray-600 dark:text-gray-400">
        <div
                class="grid grid-flow-col grid-rows-3 gap-x-6 gap-y-2 sm:gap-x-16 md:grid-flow-row md:grid-cols-3 md:grid-rows-none"
        >
			<span class="grid grid-cols-[72px_max-content]">
				<span>Teaching:</span>
				<span>{rating.rubric.teaching} / 5.0</span>
			</span>
            <span class="grid grid-cols-[72px_max-content]">
				<span>Materials:</span>
				<span>{rating.rubric.materials} / 5.0</span>
			</span>
            <span class="grid grid-cols-[72px_max-content]">
				<span>Value:</span>
				<span>{rating.rubric.value} / 5.0</span>
			</span>
            <span class="grid grid-cols-[72px_max-content]">
				<span>Difficulty:</span>
				<span
                >{rating.rubric.difficulty.slice(0, 1).toUpperCase() +
                rating.rubric.difficulty.slice(1)}</span
                >
			</span>
            <span class="grid grid-cols-[72px_max-content]">
				<span>Workload:</span>
				<span
                >{rating.rubric.workload.slice(0, 1).toUpperCase() +
                rating.rubric.workload.slice(1)}</span
                >
			</span>
            <span class="grid grid-cols-[72px_max-content]">
				<span>Grading:</span>
				<span
                >{rating.rubric.grading.slice(0, 1).toUpperCase() + rating.rubric.grading.slice(1)}</span
                >
			</span>
        </div>
    </div>

    <!--Comment-->
    <div class="flex flex-col gap-y-2">
        <div class="space-y-2">
            {#each rating.comment.split('\n') as line}
                <p class="whitespace-pre-wrap text-sm">{line}</p>
            {/each}
        </div>

        <div class="flex flex-row justify-between">
            <LikesDislikes
                    dislikedByUser={rating.dislikedByUser}
                    dislikes={rating.dislikes}
                    likedByUser={rating.likedByUser}
                    likes={rating.likes}
                    {onDislike}
                    {onLike}
            />
            <PrimaryButton on:click={onReportAbuse} size="sm">Report abuse</PrimaryButton>
        </div>
    </div>
</div>
