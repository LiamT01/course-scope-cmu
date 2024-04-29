export const isTokenValid = (expiry: string | null | undefined): boolean => {
    const expiryDate = expiry ? new Date(expiry) : null;
    return !!(expiryDate && expiryDate > new Date());
}

export const isTokenExpired = (expiry: string | null | undefined): boolean => {
    const expiryDate = expiry ? new Date(expiry) : null;
    return !!(expiryDate && expiryDate < new Date());
}
