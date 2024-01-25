'use client'

import { ColumnDef } from '@tanstack/react-table'
import { Checkbox } from '@/components/ui/checkbox'
import { Lens } from '@/db/models/lens'
import { DataTableColumnHeader } from '@/components/data-table-column-header'
import { buttonVariants } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import Link from 'next/link'
import { getRouteAsPath, teamsWorkloadLensPath } from '@/lib/utils'

export const columns: ColumnDef<Lens>[] = [
  {
    id: 'select',
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={value => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="translate-y-[2px]"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={value => row.toggleSelected(!!value)}
        aria-label="Select row"
        className="translate-y-[2px]"
      />
    ),
    enableSorting: false,
    enableHiding: false
  },
  {
    accessorKey: 'id',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="ID" />
    ),
    cell: ({ row }) => <div>{row.getValue('id')}</div>,
    enableSorting: false
  },
  {
    accessorKey: 'name',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Name" />
    ),
    cell: ({ row }) => {
      const team = row.original.workloads[0].teams[0].slug
      const id = row.original.workloads[0].id
      const lensId: string = row.getValue('id')
      const params = { id, lensId, team }
      const href = getRouteAsPath(teamsWorkloadLensPath, params)

      return (
        <Link
          href={href}
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            'hover:bg-transparent hover:underline',
            'px-0',
            'justify-start'
          )}
          passHref
        >
          <div className="flex space-x-2">
            <span className="max-w-[500px] truncate font-medium">
              {row.getValue('name')}
            </span>
          </div>
        </Link>
      )
    }
  },
  {
    accessorKey: 'version',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Version" />
    ),
    cell: ({ row }) => {
      return (
        <div className="flex space-x-2">
          <span className="max-w-[500px] truncate font-medium">
            {row.getValue('version')}
          </span>
        </div>
      )
    }
  }
]
