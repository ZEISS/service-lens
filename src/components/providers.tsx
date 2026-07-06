"use client"

import { TooltipProvider } from "@/components/ui/tooltip"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import type { ThemeProviderProps } from "next-themes"
import { NuqsAdapter } from "nuqs/adapters/next/app"
import { useState } from "react"

export function QueryProvider({ children, ...props }: ThemeProviderProps) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000, // 1 minute
            gcTime: 5 * 60 * 1000, // 5 minutes
          },
        },
      }),
  )

  return (
    <QueryClientProvider client={queryClient} {...props}>
      <TooltipProvider delayDuration={120}>
        <NuqsAdapter>{children}</NuqsAdapter>
      </TooltipProvider>
    </QueryClientProvider>
  )
}
