// src/styles/theme.ts
export const theme = {
    colors: {
        primary: "#3b82f6",
        secondary: "#64748b",
        danger: "#ef4444",
        background: "#f8fafc",
        text: "#1e293b",
        white: "#ffffff",
    },
    Container: {
        default: "600px",
        form: "450px",
        mobile: "90%",
    },
    fontSize: {
        small: "0.875rem",
        medium: "1rem",
        large: "1.25rem",
    },
};

// TypeScript用の型定義
export type Theme = typeof theme;
