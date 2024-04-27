export const clickOutside = (node: HTMLElement, options: {
    callback: () => void,
    excludedElement: HTMLElement | null
}) => {
    const handleClick = (event: MouseEvent) => {
        if (!event?.target) return;

        // Check if the click is outside the node and not on the excluded element
        if (
            node &&
            !node.contains(event.target as Node) &&
            !options.excludedElement?.contains(event.target as Node) &&
            !event.defaultPrevented
        ) {
            options.callback();
        }
    };

    document.addEventListener('click', handleClick, true);

    return {
        destroy() {
            document.removeEventListener('click', handleClick, true);
        }
    };
};
