<script lang="ts">
	import type { ProjectEntry } from '$lib/data/projects.js';

	interface Props {
		project: ProjectEntry | null;
		onClose: () => void;
	}

	let { project, onClose }: Props = $props();

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

{#if project}
	<div
		data-backdrop="true"
		role="dialog"
		aria-modal="true"
		aria-labelledby="project-modal-title"
		class="fixed inset-0 z-[100] flex items-center justify-center bg-black/50 p-4 font-jost"
		onclick={handleBackdropClick}
	>
		<div
			class="relative flex max-h-[90vh] w-full max-w-2xl flex-col rounded-2xl bg-white shadow-lg"
			onclick={(e) => e.stopPropagation()}
		>
			<div class="flex shrink-0 items-center justify-between border-b border-black/10 px-6 py-4">
				<h2 id="project-modal-title" class="text-lg font-bold tracking-tight text-black/90">
					{project.name}
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
			<div class="min-h-0 flex-1 overflow-y-auto px-6 py-4 text-sm text-black/80">
				{#each parseBody(project.longDescription) as block}
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
				{#if project.images && project.images.length > 0}
					<div class="mt-6 flex flex-col gap-4">
						{#each project.images as src}
							<div class="max-w-full overflow-hidden rounded-lg">
								<img
									src={src}
									alt=""
									class="max-h-64 w-full object-contain"
									loading="lazy"
								/>
							</div>
						{/each}
					</div>
				{/if}
			</div>
		</div>
	</div>
{/if}
