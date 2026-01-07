<script lang="ts">
	import Header from './Header.svelte';
	import './layout.css';
	import "../app.css";


	let { children } = $props();
</script>

<div class="flex flex-col min-h-screen">
	<div class="fixed top-0 left-0 w-full z-50">
        <Header />
    </div>

	<main class="flex flex-col items-center h-screen pt-20">
		{@render children()}
	</main>

</div>

<style lang="postcss">
	@reference "tailwindcss";
	
</style>
