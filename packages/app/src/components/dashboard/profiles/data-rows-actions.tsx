'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Row } from '@tanstack/react-table'
import { Button } from '@/components/ui/button'
import { api } from '@/trpc/client'
import { rhfActionDeleteProfile } from '@/actions/profile.action'
import { useAction } from '@/trpc/client'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import Link from 'next/link'
import { useToast } from '@/components/ui/use-toast'

interface DataTableRowActionsProps<TData> {
  row: Row<TData>
}

async function deleteWorkload(id: string): Promise<void> {
  await api.deleteWorkload.query(id)
}

export function DataTableRowActions<TData>({
  row
}: DataTableRowActionsProps<TData>) {
  const id = row.getValue('id') as string
  const mutation = useAction(rhfActionDeleteProfile)
  const { toast } = useToast()

  const handleOnClickDelete = async (id: string) => {
    await mutation.mutateAsync(id)
    toast({
      title: 'Successfully deleted'
    })
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
        >
          <DotsHorizontalIcon className="h-4 w-4" />
          <span className="sr-only">Open menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        <Link href={`/dashboard/profiles/${id}`}>
          <DropdownMenuItem>View</DropdownMenuItem>
        </Link>
        <DropdownMenuItem>Make a copy</DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={() => handleOnClickDelete(id)}>
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
