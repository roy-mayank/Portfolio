<script lang="ts">
	import WritingsList from '$lib/components/WritingsList.svelte';
	import { techBlogs } from '$lib/data/writings.js';
</script>

<div class="relative flex min-h-screen w-full flex-col items-center pt-36 pb-12 font-jost">
	<h1 class="mb-8 text-2xl font-bold tracking-tight text-black/90">Tech blogs</h1>
	<WritingsList items={techBlogs} type="tech" />
</div>
