<script lang="ts">
	import type { WritingEntry } from '$lib/data/writings.js';

	interface Props {
		item: WritingEntry | null;
		type: 'plot' | 'tech';
		onClose: () => void;
	}

	let { item, onClose }: Props = $props();

	const BULLET = /^\s*[-*•]\s+/;

	type BodyBlock = { type: 'list'; items: string[] } | { type: 'paragraph'; text: string };

	function parseBody(body: string): BodyBlock[] {
		const lines = body.split('\n');
		const blocks: BodyBlock[] = [];
		let i = 0;
		while (i < lines.length) {
			const line = lines[i];
			if (BULLET.test(line)) {
				const items: string[] = [];
				while (i < lines.length && BULLET.test(lines[i])) {
					items.push(lines[i].replace(BULLET, '').trim());
					i++;
				}
				if (items.length) blocks.push({ type: 'list', items });
			} else {
				const para: string[] = [];
				while (i < lines.length && !BULLET.test(lines[i])) {
					para.push(lines[i]);
					i++;
				}
				const text = para.join('\n').trim();
				if (text) blocks.push({ type: 'paragraph', text });
			}
		}
		return blocks;
	}

	function handleBackdropClick(e: MouseEvent) {
		if ((e.target as HTMLElement).dataset.backdrop === 'true') onClose();
	}
</script>

{#if item}
	<div
		data-backdrop="true"
		role="dialog"
		aria-modal="true"
		aria-labelledby="modal-title"
		class="fixed inset-0 z-[100] flex items-center justify-center bg-black/50 p-4 font-jost"
		onclick={handleBackdropClick}
	>
		<div
			class="relative flex max-h-[90vh] w-full max-w-2xl flex-col rounded-2xl bg-white shadow-lg"
			onclick={(e) => e.stopPropagation()}
		>
			<div class="flex shrink-0 items-center justify-between border-b border-black/10 px-6 py-4">
				<h2 id="modal-title" class="text-lg font-bold tracking-tight text-black/90">
					{item.title}
				</h2>
				<button
					type="button"
					class="rounded p-1 text-black/60 transition-colors hover:bg-black/10 hover:text-[#ff3e00]"
					aria-label="Close"
					onclick={onClose}
				>
					<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<path d="M18 6 6 18" />
						<path d="m6 6 12 12" />
					</svg>
				</button>
			</div>
			<div class="flex shrink-0 flex-wrap gap-2 px-6 py-2">
				{#each item.tags as tag}
					<span class="rounded-full bg-black/10 px-3 py-0.5 text-xs font-medium text-black/70">
						{tag}
					</span>
				{/each}
			</div>
			<div class="min-h-0 flex-1 overflow-y-auto px-6 py-4 text-sm text-black/80">
				{#each parseBody(item.body) as block}
					{#if block.type === 'list'}
						<ul class="my-2 list-disc pl-5 space-y-1">
							{#each block.items as listItem}
								<li>{listItem}</li>
							{/each}
						</ul>
					{:else}
						<p class="whitespace-pre-wrap my-2">{block.text}</p>
					{/if}
				{/each}
			</div>
		</div>
	</div>
{/if}
