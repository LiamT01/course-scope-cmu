import {addErrorToast} from "$lib/toast/stores";

export default function displayErrors(detail: string | { [key: string]: string }) {
    if (typeof detail === "string") {
        addErrorToast(detail)
    } else {
        for (const [key, value] of Object.entries(detail)) {
            addErrorToast(`Field "${key}": ${value}`);
        }
    }
}