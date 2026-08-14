"use client"

import { siGoogle } from "simple-icons"

import { SimpleIcon } from "@/components/simple-icon"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { signIn } from "@/lib/auth-client"

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
