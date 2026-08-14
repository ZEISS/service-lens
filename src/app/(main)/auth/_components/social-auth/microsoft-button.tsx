"use client"

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { signIn } from "@/lib/auth-client"

export function MicrosoftButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  return (
    <Button
      variant="secondary"
      className={cn(className)}
      onClick={() => signIn.social({ provider: "microsoft" })}
      {...props}
    >
      Continue with Microsoft
    </Button>
  )
}
