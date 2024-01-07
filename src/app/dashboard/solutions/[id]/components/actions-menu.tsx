'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Button } from '@/components/ui/button'
import { useAction } from '@/trpc/client'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
  DropdownMenuLabel
} from '@/components/ui/dropdown-menu'
import type { Solution } from '@/db/models/solution'
import { rhfDeleteSolutionAction } from './actions-menu.action'
import { rhfDeleteSolutionActionSchema } from './actions-menu.schema'
import { z } from 'zod'
import { useEffect } from 'react'
import { redirect } from 'next/navigation'

interface ActionsMenuProps {
  solution: Solution
}

export function ActionsMenu({ solution }: ActionsMenuProps) {
  const deleteMutation = useAction(rhfDeleteSolutionAction)
  const handleDelete = async (
    solutionId: z.infer<typeof rhfDeleteSolutionActionSchema>
  ) => await deleteMutation.mutateAsync(solutionId)

  useEffect(() => {
    deleteMutation.status === 'success' && redirect('/dashboard/solutions')
  }, [deleteMutation.status])

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant={'outline'}>
          <DotsHorizontalIcon className="h-4 w-4" />
          <span className="sr-only">Open actions</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        <DropdownMenuItem
          onClick={() => solution?.id && handleDelete(solution?.id)}
        >
          <span>Delete</span>
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
