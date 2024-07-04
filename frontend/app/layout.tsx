import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { TooltipProvider } from "@/components/ui/tooltip";
import { AuthProvider } from "@/context/authContext";
import { Suspense } from "react";
import Loading from "@/app/loading";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "eve's tracker",
  description: "All-in-one tracker for midwife data",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AuthProvider>
        <TooltipProvider>
          <Suspense fallback={<Loading />}>
            <body className={inter.className}>{children}</body>
          </Suspense>
        </TooltipProvider>
      </AuthProvider>
    </html>
  );
}
