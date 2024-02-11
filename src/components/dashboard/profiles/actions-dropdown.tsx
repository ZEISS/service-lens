'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import type { Profile } from '@/db/models/profile'
import { useAction } from '@/trpc/client'
import { rhfActionDeleteProfile } from '@/actions/profile.action'
import { Separator } from '@/components/ui/separator'

interface ActionsDropdownProps {
  profile?: Profile | null
}

export function ActionsDropdown({ profile }: ActionsDropdownProps) {
  const mutation = useAction(rhfActionDeleteProfile)
  const router = useRouter()
  const handleOnClickDelete = async () => {
    await mutation.mutate(profile?.id ?? '')
    router.replace('/dashboard/profiles')
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant={'outline'}>
          <DotsHorizontalIcon className="h-4 w-4" />
          <span className="sr-only">Modify a profile</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        {profile && (
          <Link href={`/dashboard/profiles/${profile?.id}?editable=true`}>
            <DropdownMenuItem>
              Edit
              <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
            </DropdownMenuItem>
          </Link>
        )}
        <Separator />
        <DropdownMenuItem onClick={handleOnClickDelete}>
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
