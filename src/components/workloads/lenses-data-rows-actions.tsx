'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Row } from '@tanstack/react-table'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import Link from 'next/link'
import { useAction } from '@/trpc/client'
import {
  rhfActionDeleteLens,
  rhfActionPushlishLens
} from '@/actions/lens.action'
import { useParams } from 'next/navigation'

interface DataTableRowActionsProps<TData> {
  row: Row<TData>
}

export function DataTableRowActions<TData>({
  row
}: DataTableRowActionsProps<TData>) {
  const id = row.getValue('id') as string
  const mutation = useAction(rhfActionDeleteLens)
  const deleteLens = async (lensId: string) => {
    await mutation.mutate(lensId)
  }
  const publishAction = useAction(rhfActionPushlishLens)
  const publishLens = async (lensId: string) => {
    await publishAction.mutate(lensId)
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
        <Link href={`/dashboard/lenses/${id}`} passHref>
          <DropdownMenuItem>View</DropdownMenuItem>
        </Link>
        <DropdownMenuItem onClick={() => publishLens(id)}>
          Publish
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={() => deleteLens(id)}>
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
