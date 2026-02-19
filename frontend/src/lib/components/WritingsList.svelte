<script lang="ts">
	import type { WritingEntry } from '$lib/data/writings.js';
	import WritingModal from './WritingModal.svelte';

	interface Props {
		items: WritingEntry[];
		type: 'plot' | 'tech';
	}

	let { items, type }: Props = $props();

	let selectedItem = $state<WritingEntry | null>(null);

	function openModal(item: WritingEntry) {
		selectedItem = item;
	}

	function closeModal() {
		selectedItem = null;
	}
</script>

<div class="grid w-full max-w-4xl grid-cols-1 gap-4 p-6 sm:grid-cols-2 md:gap-6">
	{#each items as item (item.id)}
		<button
			type="button"
			class="rounded-2xl border border-black/10 bg-white p-5 text-left shadow-sm transition-shadow hover:shadow-md focus:outline-none focus:ring-2 focus:ring-[#ff3e00]/50"
			onclick={() => openModal(item)}
		>
			<h3 class="font-bold tracking-tight text-black/90">{item.title}</h3>
			<div class="mt-2 flex flex-wrap gap-2">
				{#each item.tags as tag}
					<span class="rounded-full bg-black/10 px-2.5 py-0.5 text-xs font-medium text-black/70">
						{tag}
					</span>
				{/each}
			</div>
		</button>
	{/each}
</div>

<WritingModal item={selectedItem} type={type} onClose={closeModal} />
