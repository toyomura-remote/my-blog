const customMediaQuery = (minWidth: number) => {
    return `@media (min-width: ${minWidth}px)`;
};

export const media = {
    tablet: customMediaQuery(768),
    desktop: customMediaQuery(1024),
};