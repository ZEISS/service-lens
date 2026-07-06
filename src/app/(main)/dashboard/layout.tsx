import type { ReactNode } from "react"

import { RootLayout } from "@/components/layout/root"

export default async function Layout({ children }: Readonly<{ children: ReactNode }>) {
  return <RootLayout>{children}</RootLayout>
}
