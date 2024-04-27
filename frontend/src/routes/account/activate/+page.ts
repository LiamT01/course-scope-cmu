export const load = async ({url}) => {
    const token = url.searchParams.get("token");

    return {
        token,
    }
}
