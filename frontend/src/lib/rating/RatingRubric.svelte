<script lang="ts">
    import RatingStats from "$lib/rating/RatingStats.svelte";
    import type {RatingStatsT} from "$lib/types";
    import round, {roundFixed} from "$lib/util/round";
    import {twMerge} from "tailwind-merge";
    import PrimaryButton from "$lib/button/PrimaryButton.svelte";
    import Star from "$lib/rating/Star.svelte";
    import RatingIcons from "$lib/rating/RatingIcons.svelte";
    import Heart from "$lib/rating/Heart.svelte";
    import MindBlown from "$lib/rating/MindBlown.svelte";
    import Clock from "$lib/rating/Clock.svelte";
    import Laugh from "$lib/rating/Laugh.svelte";

    export let stats: RatingStatsT;
    export let onContributeRating: () => void = () => {
    };

    export let key: string;

    const mapToPercentage = (stats: Record<string, number>) => {
        const total = Object.values(stats).reduce((a, b) => a + b, 0);
        return Object.entries(stats)
            .map(([k, v]) => ({name: k, value: v}))
            .map(({name, value}) => ({label: name, percentage: total > 0 ? round(value / total * 100, 0) : 0}));
    }

</script>

{#if stats}
    <div class={twMerge("flex flex-col gap-y-4", $$props.class)}>

        <div class="w-full flex flex-col items-center gap-y-2">

            <p class="flex gap-x-2 ml-10">
                <span class="text-5xl font-medium">{roundFixed(stats.avg_overall, 1)}</span>
                <RatingStats key="course-overall-rating-{key}" rubricName="overall rating"
                             stats={mapToPercentage(stats.overall)}/>
            </p>
            <div class="ml-6 flex justify-center items-center gap-x-2">
                <RatingIcons icon={Star} rating={stats.avg_overall} size={24} text="sm"/>
                <span class="text-sm">
                    ({stats.rating_count})
                </span>
            </div>

            <div class="w-full flex flex-col items-center sm:flex-row lg:flex-col justify-around">
                <div class="grid gap-x-4 grid-cols-[72px_minmax(0,1fr)_max-content] text-sm font-medium items-center">
                    <span>Teaching:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={Heart} rating={stats.avg_teaching}/>
                        <span>{roundFixed(stats.avg_teaching, 1)}</span>
                    </span>
                    <RatingStats key="course-teaching-quality-{key}" rubricName="teaching quality"
                                 stats={mapToPercentage(stats.teaching)}/>

                    <span>Materials:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={Heart} rating={stats.avg_materials}/>
                        <span>{roundFixed(stats.avg_materials, 1)}</span>
                    </span>
                    <RatingStats key="course-materials-quality-{key}" rubricName="materials quality"
                                 stats={mapToPercentage(stats.materials)}/>

                    <span>Value:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={Heart} rating={stats.avg_value}/>
                        <span>{roundFixed(stats.avg_value, 1)}</span>
                    </span>
                    <RatingStats key="course-amount-value-{key}" rubricName="amount of value"
                                 stats={mapToPercentage(stats.value)}/>
                </div>

                <div class="grid gap-x-4 grid-cols-[72px_minmax(0,1fr)_max-content] text-sm font-medium items-center">
                    <span>Difficulty:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={MindBlown} rating={stats.avg_difficulty}/>
                        <span>{roundFixed(stats.avg_difficulty, 1)}</span>
                    </span>
                    <RatingStats key="course-difficulty-{key}" rubricName="difficulty"
                                 stats={mapToPercentage(stats.difficulty)}/>

                    <span>Workload:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={Clock} rating={stats.avg_workload}/>
                        <span>{roundFixed(stats.avg_workload, 1)}</span>
                    </span>
                    <RatingStats key="course-workload-{key}" rubricName="workload"
                                 stats={mapToPercentage(stats.workload)}/>

                    <span>Grading:</span>
                    <span class="flex gap-x-4">
                        <RatingIcons icon={Laugh} rating={stats.avg_grading}/>
                        <span>{roundFixed(stats.avg_grading, 1)}</span>
                    </span>
                    <RatingStats key="course-grading-{key}" rubricName="grading"
                                 stats={mapToPercentage(stats.grading)}/>
                </div>
            </div>
        </div>

        <div class="flex flex-col gap-y-4">
            <!--            <OutlineButton size="sm">-->
            <!--                Taking / Took it (0)-->
            <!--            </OutlineButton>-->
            <PrimaryButton on:click={onContributeRating} size="sm">
                Rate it!
            </PrimaryButton>
        </div>
    </div>
{/if}
