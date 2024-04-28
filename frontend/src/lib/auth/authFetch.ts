import {fetchWithinPage} from "$lib/auth/fetchClient";
import {openLogInModal} from "$lib/modal/stores";
import {goto, invalidateAll} from "$app/navigation";
import {expiryStore, tokenStore, userStore} from "$lib/auth/stores";
import {get} from "svelte/store";
import {addErrorToast, addSuccessToast} from "$lib/toast/stores";
import {page} from "$app/stores";
import type {RatingIn} from "$lib/types";
import {apiBaseUrl} from "$lib/constants";

const andrewIDPattern = /^(?=.{1,20}$)[a-z]+[0-9]*$/;
const usernamePattern = /^[\w.@+-]{1,30}$/;
const passwordPattern = /((?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[\W]).{8,64})/;


const convertFormToJSON = (form: HTMLFormElement): { [key: string]: any } => {
    const formData = new FormData(form);
    const json: { [key: string]: any } = {};
    for (const [key, value] of formData) {
        json[key] = value;
    }
    return json;
}

const validateField = (field: string, pattern: RegExp): boolean => {
    return pattern.test(field);
}

const validateAndrewId = (andrewID: string): boolean => {
    return validateField(andrewID, andrewIDPattern);
}

const validateUsername = (username: string): boolean => {
    return validateField(username, usernamePattern);
}

const validatePassword = (password: string): boolean => {
    return validateField(password, passwordPattern);
}

const validateRepeatPassword = (password: string, repeatPassword: string): boolean => {
    return password === repeatPassword;
}

export const selectFormFields = (form: HTMLFormElement, field: string) => {
    return form.querySelector(`input[name="${field}"]`) as HTMLInputElement;
}

export const submitLoginFormWithinPage = async (e: Event) => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    submitButton.disabled = true;
    submitButton.textContent = "Logging in...";

    const response = await fetchWithinPage('/account/login', {
        method: 'POST',
        body: formData,
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        openLogInModal.set(false);
        form.reset();
    } else {
        // Reset password field
        const passwordField = selectFormFields(form, 'password');
        passwordField.value = "";
    }

    await invalidateAll();
}

export const logOutWithinPage = async () => {
    if (!get(userStore)) {
        return
    }

    await fetchWithinPage('/account/logout', {
        method: 'POST',
    })

    userStore.set(null);
    tokenStore.set(null);
    expiryStore.set(null);

    await invalidateAll();

    if (get(page).url.pathname.startsWith("/account/me")) {
        await goto("/");
    }
}

export interface FormSubmitResponse {
    success: boolean,
    errorFields: string[],
}

export const submitSignUpFormWithinPage = async (e: Event): Promise<FormSubmitResponse> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    const requestBody = convertFormToJSON(form);

    const andrewID = requestBody['andrew_id'] as string;
    const username = requestBody['username'] as string;
    const password = requestBody['password'] as string;
    const repeatPassword = requestBody['repeat_password'] as string;

    // Validate fields: andrew_id, username, password, repeat_password
    if (!validateAndrewId(andrewID)) {
        addErrorToast("Andrew ID is invalid.");
        return {success: false, errorFields: ["andrew_id"]};
    }
    if (!validateUsername(username)) {
        addErrorToast("Username is invalid.");
        return {success: false, errorFields: ["username"]};
    }
    if (!validatePassword(password)) {
        addErrorToast("Password is invalid.");
        return {success: false, errorFields: ["password"]};
    }
    if (!validateRepeatPassword(password, repeatPassword)) {
        addErrorToast("Passwords do not match.");
        return {success: false, errorFields: ["password", "repeat_password"]};
    }

    submitButton.disabled = true;
    submitButton.textContent = "Creating account...";

    const response = await fetchWithinPage(`${apiBaseUrl}/users`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody),
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        openLogInModal.set(false);
        addSuccessToast("Your account has been created. Please check your Andrew email for account activation.");
        form.reset();
    }

    return {success: response.ok, errorFields: []};
}

export const submitPasswordResetRequestLoggedInWithinPage = async (e: Event): Promise<boolean> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    submitButton.disabled = true;
    submitButton.textContent = "Sending email...";

    const response = await fetchWithinPage(`${apiBaseUrl}/tokens/password-reset`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Password reset email has been sent. Please check your Andrew email.");
        form.reset();
    }

    return response.ok;
}

export const submitActivationLinkRequestLoggedInWithinPage = async (e: Event): Promise<boolean> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    submitButton.disabled = true;
    submitButton.textContent = "Sending email...";

    const response = await fetchWithinPage(`${apiBaseUrl}/tokens/activation`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Account activation email has been sent. Please check your Andrew email.");
        form.reset();
    }

    return response.ok;
}

export const submitPasswordResetRequestFormWithinPage = async (e: Event): Promise<FormSubmitResponse> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    const requestBody = convertFormToJSON(form);

    // Validate andrew_id
    const andrewID = requestBody['andrew_id'] as string;
    if (!validateAndrewId(andrewID)) {
        addErrorToast("Andrew ID is invalid.");
        return {success: false, errorFields: ["andrew_id"]};
    }

    submitButton.disabled = true;
    submitButton.textContent = "Sending email...";

    const response = await fetchWithinPage(`${apiBaseUrl}/tokens/password-reset`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody),
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Password reset email has been sent. Please check your Andrew email.");
        form.reset();
    }

    return {success: response.ok, errorFields: []};
}

export const activateAccountWithinPage = async (token: string | null): Promise<boolean> => {
    if (!token || token.length !== 26) {
        addErrorToast("Your account activation token is invalid.");
        return false;
    }

    const response = await fetchWithinPage(`${apiBaseUrl}/users/activated/me`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({token}),
    })

    if (response.ok) {
        addSuccessToast("Your account has been activated.");
        return true;
    }

    return false;
}

export const submitPasswordResetFormWithinPage = async (e: Event, token: string | null): Promise<FormSubmitResponse> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    const requestBody = convertFormToJSON(form);

    const password = requestBody['password'] as string;
    const repeatPassword = requestBody['repeat_password'] as string;

    if (!token || token.length !== 26) {
        addErrorToast("Your password reset token is invalid.");
        return {success: false, errorFields: []};
    }

    if (!validatePassword(password)) {
        addErrorToast("Password is invalid.");
        return {success: false, errorFields: ["password"]};
    }

    if (!validateRepeatPassword(password, repeatPassword)) {
        addErrorToast("Passwords do not match.");
        return {success: false, errorFields: ["password", "repeat_password"]};
    }

    submitButton.disabled = true;
    submitButton.textContent = "Resetting password...";

    const response = await fetchWithinPage(`${apiBaseUrl}/users/password/me`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({...requestBody, token}),
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Your password has been reset.");
        form.reset();
    }

    return {success: response.ok, errorFields: []};
}

export const submitUsernameChangeFormWithinPage = async (e: Event): Promise<FormSubmitResponse> => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    const requestBody = convertFormToJSON(form);

    const username = requestBody['username'] as string;

    if (!validateUsername(username)) {
        addErrorToast("Username is invalid.");
        return {success: false, errorFields: ["username"]};
    }

    submitButton.disabled = true;
    submitButton.textContent = "Changing username...";

    const response = await fetchWithinPage(`${apiBaseUrl}/users/username/me`, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody),
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Your username has been changed.");
        form.reset();
        await invalidateAll();
    }

    return {success: response.ok, errorFields: []};
}

const validateOfferingID = (offeringID: number): boolean => {
    return offeringID > 0;
}

const validateOverallRating = (overallRating: number): boolean => {
    return overallRating >= 1 && overallRating <= 5;
}

const validateTeachingRating = (teachingRating: number): boolean => {
    return teachingRating >= 1 && teachingRating <= 5;
}

const validateMaterialsRating = (materialsRating: number): boolean => {
    return materialsRating >= 1 && materialsRating <= 5;
}

const validateValueRating = (valueRating: number): boolean => {
    return valueRating >= 1 && valueRating <= 5;
}

const validateDifficultyRating = (difficultyRating: number): boolean => {
    return difficultyRating >= 1 && difficultyRating <= 5;
}

const validateWorkloadRating = (workloadRating: number): boolean => {
    return workloadRating >= 1 && workloadRating <= 5;
}

const validateGradingRating = (gradingRating: number): boolean => {
    return gradingRating >= 1 && gradingRating <= 5;
}

const validateComment = (comment: string): boolean => {
    return comment.length > 0 && comment.length <= 10000;
}


const validateRatingID = (ratingID: number): boolean => {
    return ratingID > 0;
}

export const submitRatingWithinPage = async (e: Event, rating: RatingIn): Promise<FormSubmitResponse> => {
    e.preventDefault();

    if (rating.rating_id && !validateRatingID(rating.rating_id)) {
        addErrorToast("Rating ID is invalid.");
        return {success: false, errorFields: ["rating_id"]};
    }

    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    if (!validateOfferingID(rating.offering_id)) {
        addErrorToast("Course offering is invalid.");
        return {success: false, errorFields: ["offering_id"]};
    }

    if (!validateOverallRating(rating.overall)) {
        addErrorToast("Overall rating is invalid.");
        return {success: false, errorFields: ["overall"]};
    }

    if (!validateTeachingRating(rating.teaching)) {
        addErrorToast("Teaching rating is invalid.");
        return {success: false, errorFields: ["teaching"]};
    }

    if (!validateMaterialsRating(rating.materials)) {
        addErrorToast("Materials rating is invalid.");
        return {success: false, errorFields: ["materials"]};
    }

    if (!validateValueRating(rating.value)) {
        addErrorToast("Value rating is invalid.");
        return {success: false, errorFields: ["value"]};
    }

    if (!validateDifficultyRating(rating.difficulty)) {
        addErrorToast("Difficulty rating is invalid.");
        return {success: false, errorFields: ["difficulty"]};
    }

    if (!validateWorkloadRating(rating.workload)) {
        addErrorToast("Workload rating is invalid.");
        return {success: false, errorFields: ["workload"]};
    }

    if (!validateGradingRating(rating.grading)) {
        addErrorToast("Grading rating is invalid.");
        return {success: false, errorFields: ["grading"]};
    }

    if (!validateComment(rating.comment)) {
        addErrorToast("Comment is invalid. Length must be between 1 and 10000 characters.");
        return {success: false, errorFields: ["comment"]};
    }

    submitButton.disabled = true;
    submitButton.textContent = rating.rating_id ? "Updating..." : "Submitting...";

    const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${rating.rating_id || ''}`, {
        method: rating.rating_id ? 'PUT' : 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(rating),
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast(rating.rating_id ? "Your rating has been updated." : "Your rating has been submitted.");
    }

    return {success: response.ok, errorFields: []};
}


export const deleteRatingWithinPage = async (e: Event, ratingID: number): Promise<boolean> => {
    e.preventDefault();

    const form = e.target as HTMLFormElement;

    // Get submit button and disable it
    const submitButton = form.querySelector('button[type="submit"]') as HTMLButtonElement;
    const buttonLabel = submitButton.textContent;

    if (!validateRatingID(ratingID)) {
        addErrorToast("Rating ID is invalid.");
        return false;
    }

    submitButton.disabled = true;
    submitButton.textContent = "Deleting...";

    const response = await fetchWithinPage(`${apiBaseUrl}/ratings/${ratingID}`, {
        method: 'DELETE',
    })

    submitButton.disabled = false;
    submitButton.textContent = buttonLabel;

    if (response.ok) {
        addSuccessToast("Your rating has been deleted.");
        await invalidateAll();
    }

    return response.ok;
}