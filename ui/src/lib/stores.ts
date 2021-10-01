import { writable } from 'svelte/store';

export const settings = writable<Object>({});

export function updateSettings(newSettings: Object) {
	settings.update((v) => {
		return {
			v,
			...newSettings
		};
	});
}

interface Head {
	title: string;
	description?: string;
}

export const view = writable<Head>({
	title: 'Home',
	description: 'Response is an open source CAD, MDT, and RMS for gaming communities.'
});

export function updateTitle(title: string) {
	view.update((value) => {
		return {
			...value,
			title
		};
	});
}
