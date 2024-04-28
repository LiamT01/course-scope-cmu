<script lang="ts">
    import CourseDetailCard from "$lib/course-card/CourseDetailCard.svelte";
    import RatingRubric from "$lib/rating/RatingRubric.svelte";
    import RatingEditForm from "$lib/rating/RatingEditForm.svelte";
    import type {Course, CourseOffering, Metadata, Rating, RatingStats} from "$lib/types";
    import RatingCard from "$lib/rating-card/RatingCard.svelte";
    import Pagination from "$lib/pagination/Pagination.svelte";
    import {goto} from "$app/navigation";
    import SearchFilters from "$lib/search/SearchFilters.svelte";
    import makeQueryParams from "$lib/util/makeQueryParams";
    import {userStore} from "$lib/auth/stores";
    import {openLogInModal} from "$lib/modal/stores";
    import Title from "$lib/section/Title.svelte";
    import {LinkOutline} from "flowbite-svelte-icons";
    import GhostButton from "$lib/button/GhostButton.svelte";
    import OutlineButton from "$lib/button/OutlineButton.svelte";

    export let data: {
        course: Course,
        courseOfferings: CourseOffering[],
        ratings: Rating[],
        stats: RatingStats,
        sort: string,
        semester: string,
        year: string,
        instructorIDs: string[],
        overall: string,
        metadata: Metadata,
    };

    let offeringOptions: { value: number, name: string }[] = [
        ...data.courseOfferings.map(offering => {
            const instructorNamesName = offering.instructors.map(instructor => instructor.name).join("; ");
            return {
                value: offering.id,
                name: `${offering.semester} ${offering.year} (${instructorNamesName})`,
            }
        })
    ];

    let semesterOptions: { value: string, name: string }[] = [
        {value: "all", name: "All"},
        ...data.courseOfferings
            .map(offering => offering.semester)
            .filter((semester, index, self) => self.indexOf(semester) === index)
            .sort((a, b) => a.localeCompare(b))
            .map(semester => ({value: semester, name: semester}))
    ];

    let yearOptions: { value: string, name: string }[] = [
        {value: "all", name: "All"},
        ...data.courseOfferings
            .map(offering => offering.year)
            .filter((year, index, self) => self.indexOf(year) === index)
            .sort((a, b) => b - a)
            .map(year => ({value: year.toString(), name: year.toString()}))
    ];

    let instructorOptions: { value: string, name: string }[] = data.courseOfferings
        .map(offering => offering.instructors)
        .flat()
        .map(instructor => ({value: instructor.id.toString(), name: instructor.name}));

    let reportAbuseModal = false;

    const openReportAbuse = () => {
        reportAbuseModal = true;
    }

    const closeReportAbuse = () => {
        reportAbuseModal = false;
    }

    const onContributeRating = () => {
        if ($userStore) {
            const elementTop = document.getElementById("rating-edit-form")?.getBoundingClientRect().top ?? 0;
            const y = elementTop - document.body.getBoundingClientRect().top;
            window.scrollTo({top: y, behavior: "smooth"});
        } else {
            $openLogInModal = true;
        }
    }

    const scrollToFirstRating = () => {
        const documentTop = document.body.getBoundingClientRect().top;
        // const rubricContainerTop = document.getElementById("main-rating-rubric-container")?.getBoundingClientRect().top ?? 0;
        const ratingsContainerTop = document.getElementById("course-ratings-container")?.getBoundingClientRect().top ?? 0;
        // if (rubricContainerTop !== 0) {
        //     window.scrollTo({top: rubricContainerTop - documentTop, behavior: "smooth"});
        // } else if (ratingsContainerTop !== 0) {
        window.scrollTo({top: ratingsContainerTop - documentTop, behavior: "smooth"});
        // }
    }

    const onApplyFilters = () => {
        const allQuery = makeQueryParams({
            semester: data.semester,
            year: data.year,
            instructorIDs: data.instructorIDs,
            overall: data.overall,
            sort: data.sort,
            page: 1,
            pageSize: data.metadata.page_size,
        });
        const currentScrollY = window.scrollY;
        goto(`?${allQuery}`)
            .then(() => scrollTo({top: currentScrollY}))
        // .then(scrollAfterApplyFilters);
    }

    // const onSubmitSuccess = () => {
    //     const allQuery = makeQueryParams({
    //         semester: data.semester,
    //         year: data.year,
    //         instructorIDs: data.instructorIDs,
    //         overall: data.overall,
    //         sort: data.sort,
    //         page: 1,
    //         pageSize: data.metadata.page_size,
    //     });
    //     const currentScrollY = window.scrollY;
    //     goto(`?${allQuery}`)
    //         .then(() => scrollTo({top: currentScrollY}))
    //         .then(scrollToFirstRating)
    //         .then(() => console.log("After submit success"))
    // }

    const previous = () => {
        if (data.metadata.current_page > data.metadata.first_page) {
            const allQuery = makeQueryParams({
                semester: data.semester,
                year: data.year,
                instructorIDs: data.instructorIDs,
                overall: data.overall,
                sort: data.sort,
                page: data.metadata.current_page - 1,
                pageSize: data.metadata.page_size,
            });
            const currentScrollY = window.scrollY;
            goto(`?${allQuery}`)
                .then(() => scrollTo({top: currentScrollY}))
                .then(scrollToFirstRating);
        }

    };
    const next = () => {
        if (data.metadata.current_page < data.metadata.last_page) {
            const allQuery = makeQueryParams({
                semester: data.semester,
                year: data.year,
                instructorIDs: data.instructorIDs,
                overall: data.overall,
                sort: data.sort,
                page: data.metadata.current_page + 1,
                pageSize: data.metadata.page_size,
            });
            const currentScrollY = window.scrollY;
            goto(`?${allQuery}`)
                .then(() => scrollTo({top: currentScrollY}))
                .then(scrollToFirstRating);
        }
    };

    const jump = (e: { detail: { page: number } }) => {
        if (e.detail.page >= data.metadata.first_page &&
            e.detail.page <= data.metadata.last_page &&
            e.detail.page !== data.metadata.current_page) {
            const allQuery = makeQueryParams({
                semester: data.semester,
                year: data.year,
                instructorIDs: data.instructorIDs,
                overall: data.overall,
                sort: data.sort,
                page: e.detail.page,
                pageSize: data.metadata.page_size,
            });
            const currentScrollY = window.scrollY;
            goto(`?${allQuery}`)
                .then(() => scrollTo({top: currentScrollY}))
                .then(scrollToFirstRating);
        }
    }
</script>

<!--Use flex-col with gap-y-4 instead of space-y-4 to avoid mis-positioning of modal backdrop-->
<div class="flex flex-col gap-y-6">

    <div class="grid grid-cols-1 lg:grid-cols-[300px_minmax(0,1fr)] gap-x-8 gap-y-8">
        <aside class="hidden lg:block">
            <div class="lg:sticky lg:z-20 lg:top-0 flex flex-col gap-y-4 pt-4">
                <RatingRubric key="aside" {onContributeRating} stats={data.stats}/>
            </div>
        </aside>

        <main class="relative flex flex-col gap-y-8">
            <div class="space-y-4">
                <Title>Course info</Title>
                <CourseDetailCard course={data.course} courseOfferings={data.courseOfferings}/>
            </div>

            <div class="flex flex-col gap-y-4 lg:hidden" id="main-rating-rubric-container">
                <Title>Rating stats</Title>
                <RatingRubric key="main" {onContributeRating} stats={data.stats}/>
            </div>


            <div class="space-y-4" id="search-filters">
                <Title>Search filters</Title>
                <SearchFilters bind:instructors={data.instructorIDs}
                               bind:overall={data.overall} bind:semester={data.semester}
                               bind:sort={data.sort}
                               bind:year={data.year}
                               {instructorOptions}
                               onSubmit={onApplyFilters}
                               {semesterOptions}
                               {yearOptions}/>
            </div>


            {#if data.ratings?.length > 0}
                <div id="course-ratings-container" class="flex flex-col gap-y-4">
                    <Title>User ratings ({data.metadata.total})</Title>
                    {#each data.ratings as rating (rating.id)}
                        <!--                                            <RatingDetailCard rating={rating} onReportAbuse={openReportAbuse}/>-->
                        <RatingCard {rating}/>
                    {/each}
                    <Pagination class="place-self-center" metadata={data.metadata} on:next={next} on:previous={previous}
                                on:jump={jump} icon/>
                </div>
            {:else if data.metadata.current_page > 1}
                <OutlineButton on:click={previous}>
                    No ratings on page {data.metadata.current_page}. Click here to go back.
                </OutlineButton>
            {:else}
                <p class="text-sm text-center">
                    No ratings found.
                </p>
            {/if}
        </main>
    </div>

    {#if $userStore}
        <div class="space-y-4" id="rating-edit-form">
            <Title>New rating</Title>
            <RatingEditForm {offeringOptions} on:success={scrollToFirstRating}/>
        </div>
    {/if}
</div>

<!--<Modal bind:open={reportAbuseModal} class="space-y-4">-->
<!--    <form action="#" class="flex flex-col gap-y-4">-->
<!--        <h3 class="text-xl font-medium text-gray-900 dark:text-white">Sign in to our platform</h3>-->
<!--        <Label class="space-y-2">-->
<!--            <span>Email</span>-->
<!--            <Input name="email" placeholder="name@company.com" required type="email"/>-->
<!--        </Label>-->
<!--        <Label class="space-y-2">-->
<!--            <span>Your password</span>-->
<!--            <Input name="password" placeholder="•••••" required type="password"/>-->
<!--        </Label>-->
<!--        <div class="flex items-start">-->
<!--            <Checkbox>Remember me</Checkbox>-->
<!--        </div>-->
<!--    </form>-->
<!--    <div class="flex gap-x-4">-->
<!--        <OutlineButton class="flex-1" on:click={closeReportAbuse} size="sm">Cancel</OutlineButton>-->
<!--        <PrimaryButton class="flex-1" on:click={closeReportAbuse} size="sm">Submit</PrimaryButton>-->
<!--    </div>-->
<!--</Modal>-->
