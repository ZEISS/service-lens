"use client"

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

export function MicrosoftButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  return (
    <Button variant="secondary" className={cn(className)} {...props}>
      Continue with Microsoft
    </Button>
  )
}
