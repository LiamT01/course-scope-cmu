import {type Writable, writable} from "svelte/store";

export const openLogInModal: Writable<boolean> = writable(false);