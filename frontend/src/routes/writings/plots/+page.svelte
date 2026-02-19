<script lang="ts">
	import WritingsList from '$lib/components/WritingsList.svelte';
	import { plotWritings } from '$lib/data/writings.js';
</script>

<div class="relative flex min-h-screen w-full flex-col items-center pt-36 pb-12 font-jost lg:px-52 px-30 text-justify">
	<h1 class="mb-8 text-2xl font-bold tracking-tight text-black/90">Plot Writings</h1>
	<p class="text-white"> Lately I have been expressing interest in writing plots to stories. My strengths appear to be the topics of the tragedy of "the human experience". I've never been much of a reader or writer and therefore these are just plots and not complete stories but for some reason these appear to flow out of my mind at pace occassionally </p>
	<WritingsList items={plotWritings} type="plot" />
</div>
