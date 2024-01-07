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
import { useRouter } from 'next/navigation'
import { Lens } from '@/db/models/lens'
import { useAction } from '@/trpc/client'
import { rhfActionDeleteLens } from '@/app/dashboard/lenses/actions/lens.action'

interface ActionsDropdownProps {
  lens?: Lens | null
}

export function ActionsDropdown({ lens }: ActionsDropdownProps) {
  const mutation = useAction(rhfActionDeleteLens)
  const router = useRouter()
  const handleOnClickDelete = async () => {
    await mutation.mutate(lens?.id ?? '')
    router.replace('/dashboard/lenses')
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline">
          <DotsHorizontalIcon className="h-4 w-4" />
          <span className="sr-only">Create new solution</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        <DropdownMenuItem onClick={handleOnClickDelete}>
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
