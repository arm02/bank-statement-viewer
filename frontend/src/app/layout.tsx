import "../styles/globals.css";
import type { Metadata } from "next";

export const metadata: Metadata = {
    title: "Bank Statement Viewer",
    description: "Full-stack Go + Next.js assignment",
    icons: {
        icon: "/favicon.ico",
    },
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <body>{children}</body>
        </html>
    );
}
