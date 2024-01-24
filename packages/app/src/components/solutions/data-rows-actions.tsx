'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Row } from '@tanstack/react-table'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import { useAction } from '@/trpc/client'
import { rhfActionDeleteSolution } from '@/actions/solution.action'
import { RHfActionDeleteSolution } from '@/actions/solution.schema'

interface DataTableRowActionsProps<TData> {
  row: Row<TData>
}

export function DataTableRowActions<TData>({
  row
}: DataTableRowActionsProps<TData>) {
  const id = row.getValue('id') as string
  const mutation = useAction(rhfActionDeleteSolution)
  const deleteSolution = async (data: RHfActionDeleteSolution) => {
    await mutation.mutate(data)
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
        {/* <Link href={`/dashboard/solutions/${id}`}>
          <DropdownMenuItem>View</DropdownMenuItem>
        </Link>
        <DropdownMenuItem
          onClick={async () => await api.solutions.makeCopy.query(id)}
        >
          Make a copy
        </DropdownMenuItem>
        <DropdownMenuSeparator /> */}
        <DropdownMenuItem onClick={() => deleteSolution(id)}>
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
