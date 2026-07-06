import { RootLayout } from "@/components/layout/root"

export default function Layout({ children }: Readonly<{ children: React.ReactNode }>) {
  return <RootLayout>{children}</RootLayout>
}
