<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';

	let menuState = $state(true);
	function toggleMenu() {
		menuState = !menuState;
	}

	let writingsOpen = $state(false);

	function openWritings() {
		writingsOpen = true;
	}
	function closeWritings() {
		writingsOpen = false;
	}
	function toggleWritings() {
		writingsOpen = !writingsOpen;
	}
</script>

<div
	class="flex flex-col items-center font-jost"
	onmouseleave={closeWritings}
>
	<!-- Navbar bar only -->
	<div class="font-jost m-6 flex w-2/3 flex-col items-center justify-center rounded-2xl bg-white p-4 shadow-sm md:m-8 md:flex-row-reverse md:p-6 md:rounded-4xl md:gap-5">
		<nav class="md:w-full flex flex-row justify-center xl:justify-end">
			<ul class="flex flex-row items-center justify-between px-2 text-[10px] sm:gap-14 sm:text-xs md:justify-center md:gap-14">
				<li class="group relative flex items-center justify-center transition-transform hover:scale-110" aria-current={page.url.pathname === '/' ? 'page' : undefined}>
					<a href={resolve('/')} class="font-bold tracking-widest text-black/70 transition-colors hover:text-[#ff3e00]">
						ME
					</a>
					<div class="hidden md:group-aria-[current=page]:block absolute -top-6 left-1/2 -translate-x-1/2 w-0 h-0 border-x-[6px] border-x-transparent border-t-[6px] border-t-[#ff3e00]"></div>
				</li>

				<li class="group relative flex items-center justify-center transition-transform hover:scale-110" aria-current={page.url.pathname === '/moviegame' ? 'page' : undefined}>
					<a href={resolve('/moviegame')} class="font-bold tracking-widest text-black/70 transition-colors hover:text-[#ff3e00]">
						TMDB GAME
					</a>
					<div class="hidden md:group-aria-[current=page]:block absolute -top-6 left-1/2 -translate-x-1/2 w-0 h-0 border-x-[6px] border-x-transparent border-t-[6px] border-t-[#ff3e00]"></div>
				</li>

				<li class="group relative flex items-center justify-center transition-transform hover:scale-110" aria-current={page.url.pathname === '/collabdocs' ? 'page' : undefined}>
					<a href={resolve('/collabdocs')} class="font-bold tracking-widest text-black/70 transition-colors hover:text-[#ff3e00]">
						COLLABPAD
					</a>
					<div class="hidden md:group-aria-[current=page]:block absolute -top-6 left-1/2 -translate-x-1/2 w-0 h-0 border-x-[6px] border-x-transparent border-t-[6px] border-t-[#ff3e00]"></div>
				</li>

				<li
					class="group relative flex items-center justify-center transition-transform hover:scale-110"
					aria-current={page.url.pathname.startsWith('/writings') ? 'page' : undefined}
					onmouseenter={openWritings}
					onmouseleave={closeWritings}
				>
					<button
						type="button"
						class="font-bold tracking-widest text-black/70 transition-colors hover:text-[#ff3e00] md:cursor-default"
						onclick={toggleWritings}
						aria-expanded={writingsOpen}
						aria-haspopup="true"
					>
						MY WRITINGS
					</button>
					<div class="hidden md:group-aria-[current=page]:block absolute -top-6 left-1/2 -translate-x-1/2 w-0 h-0 border-x-[6px] border-x-transparent border-t-[6px] border-t-[#ff3e00]"></div>
					{#if writingsOpen}
						<div
							class="absolute left-1/2 top-full z-10 flex -translate-x-1/2 flex-row items-center gap-0 rounded-md bg-white px-0 py-1 shadow-md"
							onmouseenter={openWritings}
							role="menu"
						>
							<a
								href={resolve('/writings/plots')}
								class="whitespace-nowrap px-2.5 py-1 text-[10px] font-bold tracking-wider text-black/70 transition-colors hover:text-[#ff3e00]"
								role="menuitem"
							>
								Plot Writings
							</a>
							<span class="text-black/30" aria-hidden="true">|</span>
							<a
								href={resolve('/writings/tech')}
								class="whitespace-nowrap px-2.5 py-1 text-[10px] font-bold tracking-wider text-black/70 transition-colors hover:text-[#ff3e00]"
								role="menuitem"
							>
								Tech blogs
							</a>
						</div>
					{/if}
				</li>
			</ul>
		</nav>
	</div>
</div>

<!-- <style>
	header {
		display: flex;
		justify-content: space-between;
	}

	.corner {
		width: 3em;
		height: 3em;
	}

	.corner a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	.corner img {
		width: 2em;
		height: 2em;
		object-fit: contain;
	}

	nav {
		display: flex;
		justify-content: center;
		--background: rgba(255, 255, 255, 0.7);
	}

	svg {
		width: 2em;
		height: 3em;
		display: block;
	}

	path {
		fill: var(--background);
	}

	ul {
		position: relative;
		padding: 0;
		margin: 0;
		height: 3em;
		display: flex;
		justify-content: center;
		align-items: center;
		list-style: none;
		background: var(--background);
		background-size: contain;
	}

	li {
		position: relative;
		height: 100%;
	}

	li[aria-current='page']::before {
		--size: 6px;
		content: '';
		width: 0;
		height: 0;
		position: absolute;
		top: 0;
		left: calc(50% - var(--size));
		border: var(--size) solid transparent;
		border-top: var(--size) solid var(--color-theme-1);
	}

	nav a {
		display: flex;
		height: 100%;
		align-items: center;
		padding: 0 0.5rem;
		color: var(--color-text);
		font-weight: 700;
		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.1em;
		text-decoration: none;
		transition: color 0.2s linear;
	}

	a:hover {
		color: var(--color-theme-1);
	}
</style> -->
