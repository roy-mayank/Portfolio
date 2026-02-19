export interface WritingEntry {
	id: string;
	title: string;
	tags: string[];
	body: string;
}

export const plotWritings: WritingEntry[] = [
	{
		id: 'story-1',
		title: 'Doom Patterns',
		tags: ['realism', 'tragedy','awareness'],
		body: `In Progress.`
	},
	{
		id: 'story-2',
		title: 'Broken interstellar',
		tags: ['tragedy', 'sci-fi'],
		body: 'In Progress.'
	}
];

export const techBlogs: WritingEntry[] = [
	{
		id: 'sample-tech',
		title: 'Sample',
		tags: ['svelte', 'firebase'],
		body: 'In Progress.'
	}
];
