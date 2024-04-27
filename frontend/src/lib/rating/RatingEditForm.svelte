<script lang="ts">
    import {Label, Select, Textarea, Toggle} from 'flowbite-svelte';
    import {twMerge} from 'tailwind-merge';
    import OutlineButton from '$lib/button/OutlineButton.svelte';
    import PrimaryButton from '$lib/button/PrimaryButton.svelte';
    import Markdown from '$lib/markdown/markdown.svelte';
    import {createEventDispatcher, onMount} from "svelte";
    import Star from "$lib/rating/Star.svelte";
    import Heart from "$lib/rating/Heart.svelte";
    import MindBlown from "$lib/rating/MindBlown.svelte";
    import Clock from "$lib/rating/Clock.svelte";
    import Laugh from "$lib/rating/Laugh.svelte";
    import RatingIcons from "$lib/rating/RatingIcons.svelte";
    import {submitRatingWithinPage} from "$lib/auth/authFetch";
    import {invalidateAll} from "$app/navigation";

    const dispatch = createEventDispatcher();

    export let offeringOptions: { value: number, name: string }[];
    export let ratingID: number | undefined = undefined;

    const onSubmit: (e: Event) => void = async (e) => {
        const response = await submitRatingWithinPage(e, {
            rating_id: ratingID,
            offering_id: offeringID,
            overall,
            teaching,
            materials,
            value,
            difficulty,
            workload,
            grading,
            comment,
        });

        if (response.success) {
            onClear();
            await invalidateAll();
            dispatch('success');
        }
    };

    export let offeringID: number = 0;
    export let overall: number = 0;
    export let teaching: number = 0;
    export let materials: number = 0;
    export let value: number = 0;
    export let difficulty: number = 0;
    export let workload: number = 0;
    export let grading: number = 0;
    export let comment: string = '';

    let markdownPreview: boolean = true;

    const onClear = () => {
        offeringID = 0;
        overall = 0;
        teaching = 0;
        materials = 0;
        value = 0;
        difficulty = 0;
        workload = 0;
        grading = 0;
        comment = '';
    }

    function syncScroll() {
        const comment = document.getElementById('comment')
        const preview = document.getElementById('preview')
        if (comment && preview) {
            preview.scrollTop = comment.scrollTop * (preview.scrollHeight - preview.clientHeight) / (comment.scrollHeight - comment.clientHeight);
        }
    }

    onMount(() => {
        const comment = document.getElementById('comment');
        comment?.addEventListener('scroll', syncScroll);
        return () => {
            comment?.removeEventListener('scroll', syncScroll);
        }
    })
</script>

<!--Use overflow-hidden to prevent Select from overflowing on iOS Safari-->
<form {...$$restProps} class={twMerge('flex flex-col gap-y-4', $$props.class)}
      on:submit={onSubmit}>
    <div class="space-y-4">
        <Label class="grid grid-cols-[120px_minmax(0,1fr)] items-center">
            <span class="whitespace-nowrap">Course Offering</span>
            <Select bind:value={offeringID}
                    class="overflow-hidden text-xs"
                    items={offeringOptions}
                    placeholder=""
                    size="sm"/>
        </Label>

        <div class="text-sm font-medium space-y-4">
            <div class="grid grid-cols-[76px_max-content] items-center">
                <span>Overall:</span>
                <div class="flex gap-x-4">
                    <RatingIcons bind:rating={overall} edit icon={Star} size={24}/>
                </div>
            </div>

            <div class="grid grid-rows-6 sm:grid-rows-3 sm:grid-flow-col md:grid-rows-none md:grid-flow-row md:grid-cols-3 gap-y-4">
                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Teaching:</span>
                    <RatingIcons bind:rating={teaching} edit icon={Heart} size={24}/>
                </div>

                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Materials:</span>
                    <RatingIcons bind:rating={materials} edit icon={Heart} size={24}/>
                </div>

                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Value:</span>
                    <RatingIcons bind:rating={value} edit icon={Heart} size={24}/>
                </div>

                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Difficulty:</span>
                    <RatingIcons bind:rating={difficulty} edit icon={MindBlown} size={24}/>
                </div>

                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Workload:</span>
                    <RatingIcons bind:rating={workload} edit icon={Clock} size={24}/>
                </div>

                <div class="grid grid-cols-[76px_max-content] items-center">
                    <span>Grading:</span>
                    <RatingIcons bind:rating={grading} edit icon={Laugh} size={24}/>
                </div>
            </div>
        </div>

        <div class="flex flex-col gap-y-4">
            <div class="space-y-4">
                <Label for="comment">Comment</Label>
                <div class="grid {markdownPreview ? 'md:grid-cols-2' : ''} gap-4">
                    <Textarea bind:value={comment}
                              class="overflow-auto"
                              id="comment"
                              name="comment"
                              placeholder="You can use Markdown here."
                              rows="20"/>
                    {#if markdownPreview}
                        <div id="preview"
                             class="max-h-[420px] rounded-lg border p-4 overflow-auto">
                            <Markdown class="" text={comment}/>
                        </div>
                    {/if}
                </div>
            </div>

            <div class="flex justify-between gap-x-4">
                <div class="flex gap-x-8">
                    <Toggle bind:checked={markdownPreview} size="small">Preview</Toggle>
                    <OutlineButton on:click={onClear} size="sm">Reset</OutlineButton>
                </div>
                <PrimaryButton size="sm" type="submit">Submit</PrimaryButton>
            </div>
        </div>
    </div>
</form>
