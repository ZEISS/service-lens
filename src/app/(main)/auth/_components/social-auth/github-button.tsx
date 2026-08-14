"use client"

import { siGithub } from "simple-icons"

import { SimpleIcon } from "@/components/simple-icon"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { signIn } from "@/lib/auth-client"

export function GitHubButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  return (
    <Button
      variant="secondary"
      className={cn(className)}
      onClick={() => signIn.social({ provider: "github" })}
      {...props}
    >
      <SimpleIcon icon={siGithub} className="size-4" />
      Continue with GitHub
    </Button>
  )
}
