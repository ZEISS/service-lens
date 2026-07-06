"use client"

import { signOut } from "@/lib/auth-client"
import { BadgeCheck, Bell, CreditCard, EllipsisVertical, LogOut } from "lucide-react"
import { redirect } from "next/navigation"
import { toast } from "sonner"

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { getInitials } from "@/lib/utils"
import type { User } from "better-auth"

export function TeamSwitcher({ user }: { user?: User }) {
  const onClick = async () => {
    await signOut({
      fetchOptions: {
        onError: (ctx) => {
          toast.error(ctx.error.message)
        },
        onSuccess: async () => {
          toast.success("Successfully logged out")
          redirect("/")
        },
      },
    })
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Avatar className="size-9 rounded-lg">
          <AvatarImage src={user?.image || undefined} alt={user?.name} />
          <AvatarFallback className="rounded-lg">{getInitials(user?.name || "")}</AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="min-w-56 space-y-1 rounded-lg" side="bottom" align="end" sideOffset={4}>
        <DropdownMenuItem>
          <Avatar className="h-8 w-8 rounded-lg grayscale">
            <Avatar className="size-9 rounded-lg">
              <AvatarImage src={user?.image || undefined} alt={user?.name} />
              <AvatarFallback className="rounded-lg">{getInitials(user?.name || "")}</AvatarFallback>
            </Avatar>
          </Avatar>
          <div className="grid flex-1 text-left text-sm leading-tight">
            <span className="truncate font-medium">{user?.name}</span>
            <span className="truncate text-muted-foreground text-xs">{user?.email}</span>
          </div>
          <EllipsisVertical className="ml-auto size-4" />
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
            <BadgeCheck />
            Account
          </DropdownMenuItem>
          <DropdownMenuItem>
            <CreditCard />
            Billing
          </DropdownMenuItem>
          <DropdownMenuItem>
            <Bell />
            Notifications
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={onClick}>
          <LogOut />
          Log out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
