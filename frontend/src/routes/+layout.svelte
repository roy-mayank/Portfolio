<script lang="ts">
	import Header from './Header.svelte';
	import './layout.css';
	import '../app.css';
	import { page } from '$app/state';
	import { browser } from '$app/environment';
	import { env } from '$env/dynamic/public';

	let { children } = $props();
	let viewCount = $state<number | null>(null);

	$effect(() => {
		if (!browser || !env.PUBLIC_GO_API_URL) return;
		const path = page.url.pathname;
		const base = env.PUBLIC_GO_API_URL.replace(/\/$/, '');
		fetch(`${base}/api/view?path=${encodeURIComponent(path)}`)
			.then((res) => res.json())
			.then((data: { count?: number }) => {
				if (typeof data?.count === 'number') viewCount = data.count;
			})
			.catch(() => {});
	});
</script>

<div class="flex flex-col min-h-screen">
	<div class="fixed top-0 left-0 w-full z-50">
		<Header />
	</div>

	<main class="flex flex-col items-center">
		{@render children()}
	</main>

	{#if browser && viewCount !== null}
		<p class="text-center text-xs text-white/50 py-2 font-jost">This page: {viewCount} view{viewCount === 1 ? '' : 's'}</p>
	{/if}
</div>

<style lang="postcss">
	@reference "tailwindcss";
	
</style>
