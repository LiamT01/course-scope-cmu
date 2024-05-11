<script lang="ts">
	import type { Course, CourseOffering, Metadata, Offering, Rating, User } from '$lib/types';
	import { Avatar } from 'flowbite-svelte';
	import { format, formatDistance, isToday } from 'date-fns';
	import { EditOutline, TrashBinOutline, UserOutline } from 'flowbite-svelte-icons';
	import LikesDislikes from '$lib/rating-card/LikesDislikes.svelte';
	import RatingIcons from '$lib/rating/RatingIcons.svelte';
	import Markdown from '$lib/markdown/markdown.svelte';
	import GhostButton from '$lib/button/GhostButton.svelte';
	import CourseInfo from '$lib/course/CourseInfo.svelte';
	import Heart from '$lib/rating/Heart.svelte';
	import MindBlown from '$lib/rating/MindBlown.svelte';
	import Clock from '$lib/rating/Clock.svelte';
	import Laugh from '$lib/rating/Laugh.svelte';
	import { openLogInModal } from '$lib/modal/stores';
	import { fetchWithinPage } from '$lib/auth/fetchWrappers';
	import Modal from '$lib/modal/Modal.svelte';
	import RatingEditForm from '$lib/rating/RatingEditForm.svelte';
	import { onMount } from 'svelte';
	import OutlineButton from '$lib/button/OutlineButton.svelte';
	import { deleteRatingWithinPage } from '$lib/auth/authFetchClient';
	import { apiBaseUrl, listOfferingsPageSize } from '$lib/constants';
	import CriticalButton from '$lib/button/CriticalButton.svelte';

	export let token: string | null;
	export let expiry: string | null;
	export let user: User | null;

	export let rating: Rating;
	const onLike = async () => {
		if (!user) {
			$openLogInModal = true;
			return;
		}

		if (rating.liked_by_viewer) {
			const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${rating.id}/like`, {
				method: 'DELETE',
				token,
				expiry
			});

			if (response.ok) {
				rating.net_likes -= 1;
				rating.liked_by_viewer = false;
			}
		} else {
			const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${rating.id}/like`, {
				method: 'POST',
				token,
				expiry
			});

			if (response.ok) {
				rating.net_likes += 1;
				rating.liked_by_viewer = true;
			}
		}
	};

	const onDislike = async () => {
		if (!user) {
			$openLogInModal = true;
			return;
		}

		if (rating.disliked_by_viewer) {
			const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${rating.id}/dislike`, {
				method: 'DELETE',
				token,
				expiry
			});

			if (response.ok) {
				rating.net_likes += 1;
				rating.disliked_by_viewer = false;
			}
		} else {
			const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${rating.id}/dislike`, {
				method: 'POST',
				token,
				expiry
			});

			if (response.ok) {
				rating.net_likes -= 1;
				rating.disliked_by_viewer = true;
			}
		}
	};

	export let showCourseDetail: boolean = false;

	let course: Course | null = rating.offering.course;
	let courseOffering: CourseOffering | null = rating.offering;

	const formatDate = (date: Date) => {
		if (isToday(date)) {
			return formatDistance(date, new Date(), { addSuffix: true });
		} else {
			return format(date, 'HH:mm MMM d, yyyy');
		}
	};

	let openRatingEditModal: boolean = false;

	let offeringOptions: { value: number; name: string }[] = [];

	const fetchCourseOfferings = async () => {
		const courseOfferingsResponse = await fetchWithinPage(
			`${apiBaseUrl}/offerings?course_id=${rating.offering.course.id}&page_size=${listOfferingsPageSize}`
		);
		const courseOfferingsData: { items: Offering[]; metadata: Metadata } =
			await courseOfferingsResponse.json();
		const courseOfferings = courseOfferingsData.items;

		return [
			...courseOfferings.map((offering) => {
				const instructorNamesName = offering.instructors
					.map((instructor) => instructor.name)
					.join('; ');
				return {
					value: offering.id,
					name: `${offering.semester} ${offering.year} (${instructorNamesName})`
				};
			})
		];
	};

	if (user?.id === rating.user.id) {
		onMount(async () => {
			offeringOptions = await fetchCourseOfferings();
		});
	}

	let openDeleteConfirmationModal: boolean = false;

	const onDelete = async (e: Event) => {
		const ok = await deleteRatingWithinPage(e, rating.id, token, expiry);
		if (ok) {
			openDeleteConfirmationModal = false;
		}
	};

	const onSubmitRatingEditSuccess = async () => {
		openRatingEditModal = false;
	};
</script>

<div class="space-y-2 border-y py-4 sm:space-y-4 sm:py-8" id="rating-{rating.id}">
	<div class="flex flex-col items-start justify-between gap-y-2">
		<div class="flex flex-col gap-y-2">
			<div class="flex items-center gap-x-4">
				<Avatar
					src={user?.id === rating.user.id ? '/authenticated-avatar.png' : '/anonymous-avatar.png'}
				/>
				<div class="text-sm">
					<div class="flex flex-wrap items-center gap-x-2">
						<p class="font-medium">{rating.user.username}</p>
						<time class="text-xs text-gray-500" dateTime={rating.updated_at.toString()}>
							{formatDate(rating.updated_at)}
						</time>
					</div>
					<RatingIcons rating={rating.overall} size={20} text="sm" />
				</div>
			</div>
		</div>

		<div class="flex flex-wrap items-center gap-x-4">
			{#if showCourseDetail && course}
				<a class="text-md font-medium" href="/courses/{course.id}">
					<span class="mr-2">{course.number}</span>
					<span>{course.name}</span>
				</a>
			{/if}
			<div class="flex flex-wrap gap-x-4 gap-y-2">
				{#if courseOffering}
					<div class="ml-2 flex items-center gap-x-2 text-xs">
						<span>{courseOffering.semester} {courseOffering.year}</span>
					</div>
				{/if}
				<CourseInfo
					class={!showCourseDetail ? 'sm:hidden' : ''}
					{course}
					{courseOffering}
					key={rating.id.toString()}
					{showCourseDetail}
				/>
				{#if courseOffering && !showCourseDetail}
					{#each courseOffering.instructors as instructor (instructor.id)}
						<div class="hidden items-center gap-x-2 text-xs sm:flex">
							<UserOutline class="h-3 w-3" />
							<a
								class="underline decoration-gray-400 underline-offset-4 hover:decoration-gray-900"
								href="/instructors/{instructor?.id}"
							>
								{instructor?.name}
							</a>
						</div>
					{/each}
				{/if}
			</div>
		</div>
	</div>

	<div>
		<Markdown text={rating.comment} />
	</div>

	<div>
		<div
			class="grid max-w-[640px] grid-flow-col grid-rows-6 gap-y-2 text-xs font-medium text-gray-600 dark:text-gray-400 sm:grid-flow-row sm:grid-cols-3 sm:grid-rows-none sm:gap-x-8"
		>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Teaching:</span>
				<span><RatingIcons icon={Heart} rating={rating.teaching} size={16} /></span>
			</span>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Materials:</span>
				<span><RatingIcons icon={Heart} rating={rating.materials} size={16} /></span>
			</span>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Value:</span>
				<span><RatingIcons icon={Heart} rating={rating.value} size={16} /></span>
			</span>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Difficulty:</span>
				<span><RatingIcons icon={MindBlown} rating={rating.difficulty} size={16} /></span>
			</span>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Workload:</span>
				<span><RatingIcons icon={Clock} rating={rating.workload} size={16} /></span>
			</span>
			<span class="grid grid-cols-[64px_max-content] items-center">
				<span>Grading:</span>
				<span><RatingIcons icon={Laugh} rating={rating.grading} size={16} /></span>
			</span>
		</div>
	</div>

	<div class="flex flex-col justify-between gap-y-2 sm:flex-row">
		<div class="order-first flex w-full items-center justify-between gap-x-8 sm:order-none">
			<div class="flex items-center gap-x-8">
				<LikesDislikes
					dislikedByUser={rating.disliked_by_viewer}
					likedByUser={rating.liked_by_viewer}
					netLikes={rating.net_likes}
					{onDislike}
					{onLike}
				/>
				{#if user?.id === rating.user.id}
					<GhostButton on:click={() => (openRatingEditModal = true)} class="h-10 w-10" size="xs">
						<EditOutline class="h-4 w-4" />
					</GhostButton>
					<GhostButton
						on:click={() => (openDeleteConfirmationModal = true)}
						class="h-10 w-10"
						size="xs"
					>
						<TrashBinOutline class="h-4 w-4" />
					</GhostButton>
				{/if}
			</div>
			<!--            <GhostButton class="h-8 w-8" on:click={onReportAbuse} size="xs">-->
			<!--                <FlagOutline class="h-4 w-4"/>-->
			<!--            </GhostButton>-->
		</div>
	</div>
</div>

{#if user?.id === rating.user.id && openRatingEditModal && offeringOptions.length > 0}
	<Modal bind:open={openRatingEditModal} class="max-h-full max-w-full overflow-y-auto">
		<div class="space-y-4">
			<h3 class="text-xl font-medium dark:text-white">Update rating</h3>
			<RatingEditForm
				{token}
				{expiry}
				{offeringOptions}
				ratingID={rating.id}
				offeringID={rating.offering.id}
				overall={rating.overall}
				teaching={rating.teaching}
				materials={rating.materials}
				value={rating.value}
				difficulty={rating.difficulty}
				workload={rating.workload}
				grading={rating.grading}
				comment={rating.comment}
				on:success={onSubmitRatingEditSuccess}
			/>
		</div>
	</Modal>
{/if}

{#if user?.id === rating.user.id && openDeleteConfirmationModal}
	<Modal bind:open={openDeleteConfirmationModal} class="w-96">
		<form class="space-y-8" on:submit={onDelete}>
			<h3 class="text-xl font-medium dark:text-white">Confirm deletion</h3>
			<p class="text-gray-600 dark:text-gray-400">Are you sure you want to delete this rating?</p>
			<div class="grid grid-cols-2 gap-x-4">
				<OutlineButton on:click={() => (openDeleteConfirmationModal = false)}>Cancel</OutlineButton>
				<CriticalButton type="submit">Delete</CriticalButton>
			</div>
		</form>
	</Modal>
{/if}
