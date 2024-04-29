import {error} from "@sveltejs/kit";
import {fetchWithinLoad} from "$lib/auth/fetchWrappers";
import {apiBaseUrl} from "$lib/constants";

export const load = async ({fetch, params, data}) => {
    const instructorID = parseInt(params.id);

    const instructorResponse = await fetchWithinLoad(
        fetch,
        `${apiBaseUrl}/instructors/${instructorID}`,
        {
            token: data.token,
            expiry: data.expiry,
        }
    );
    const instructor = await instructorResponse.json();

    if (!instructor) {
        throw error(404, "Instructor not found");
    }

    return {
        instructor,
    }
}