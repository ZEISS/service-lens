"use client"

import { siGoogle } from "simple-icons"

import { SimpleIcon } from "@/components/simple-icon"
import { Button } from "@/components/ui/button"
import { signIn } from "@/lib/auth-client"
import { cn } from "@/lib/utils"

export function GoogleButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  return (
    <Button
      variant="secondary"
      className={cn(className)}
      onClick={() => signIn.social({ provider: "google" })}
      {...props}
    >
      <SimpleIcon icon={siGoogle} className="size-4" />
      Continue with Google
    </Button>
  )
}
