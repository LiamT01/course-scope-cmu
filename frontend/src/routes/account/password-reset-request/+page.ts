export const load = async ({url}) => {
    const nextLink = url.searchParams.get('next');
    return {
        nextLink,
    }
}