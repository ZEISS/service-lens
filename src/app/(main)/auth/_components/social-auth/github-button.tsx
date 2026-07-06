import { siGithub } from "simple-icons"

import { SimpleIcon } from "@/components/simple-icon"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

export function GitHubButton({ className, ...props }: React.ComponentProps<typeof Button>) {
  return (
    <Button variant="secondary" className={cn(className)} {...props}>
      <SimpleIcon icon={siGithub} className="size-4" />
      Continue with GitHub
    </Button>
  )
}
