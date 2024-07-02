import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { TooltipProvider } from "@/components/ui/tooltip";
import { AuthProvider } from "@/context/authContext";

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
          <body className={inter.className}>{children}</body>
        </TooltipProvider>
      </AuthProvider>
    </html>
  );
}
