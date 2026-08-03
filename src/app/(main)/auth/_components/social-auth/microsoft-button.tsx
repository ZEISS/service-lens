"use client";

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { signIn } from "@/lib/auth-client"

export function MicrosoftButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  const handleSignIn = async () => {
    const data = await signIn.social({
      provider: "microsoft",
      callbackURL: "/dashboard",
    })
  }

  return (
    <Button variant="secondary" className={cn(className)} onClick={handleSignIn} {...props}>
      Continue with Microsoft
    </Button>
  )
}
