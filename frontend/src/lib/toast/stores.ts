import { type Writable, writable } from 'svelte/store';
import generateID from '$lib/util/generateID';

interface ToastIn {
	type: 'success' | 'error' | 'info';
	message: string;
	dismissible?: boolean;
	timeout?: number;
}

interface Toast extends ToastIn {
	id: string;
}

export const toasts: Writable<Toast[]> = writable([]);

export const addToast = (toast: ToastIn) => {
	const id = generateID();

	const defaults = {
		id,
		type: 'info',
		dismissible: true,
		timeout: 3000
	};

	toasts.update((all) => [{ ...defaults, ...toast }, ...all]);

	if (toast.timeout) setTimeout(() => dismissToast(id), toast.timeout);
};

export const dismissToast = (id: string) => {
	toasts.update((all) => all.filter((t) => t.id !== id));
};

const calculateTimeout = (message: string) => {
	return Math.min(2000 + message.length * 15, 10000);
};

export const addErrorToast = (message: string, dismissible: boolean = true) => {
	const timeout = calculateTimeout(message);
	addToast({ type: 'error', message, dismissible, timeout });
};

export const addSuccessToast = (message: string, dismissible: boolean = true) => {
	const timeout = calculateTimeout(message);
	addToast({ type: 'success', message, dismissible, timeout });
};
