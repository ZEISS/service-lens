'use client'

import { ColumnDef } from '@tanstack/react-table'
import { Checkbox } from '@/components/ui/checkbox'
import Link from 'next/link'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import { Workload } from '@/db/models/workload'
import { DataTableColumnHeader } from '@/components/data-table-column-header'
import { DataTableRowActions } from '@/components/workloads/data-rows-actions'
import { getRouteAsPath, teamsWorkloadUrlPath } from '@/lib/utils'
import { type Profile } from '@/db/models/profile'

export const columns: ColumnDef<Workload>[] = [
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
    cell: ({ row }) => <div className="w-[80px]">{row.getValue('id')}</div>,
    enableSorting: false,
    enableHiding: false
  },
  {
    accessorKey: 'name',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Name" />
    ),
    cell: ({ row }) => {
      const team = row.original.teams[0].slug
      const id: string = row.getValue('id')
      const url = getRouteAsPath(teamsWorkloadUrlPath, {
        id,
        team
      })

      return (
        <Link
          href={url}
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            'hover:bg-transparent hover:underline',
            'px-0',
            'justify-start'
          )}
          passHref
        >
          <div className="flex space-x-2">
            {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
            <span className="max-w-[500px] truncate font-medium">
              {row.getValue('name')}
            </span>
          </div>
        </Link>
      )
    }
  },
  {
    accessorKey: 'profile',
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Profile" />
    ),
    cell: ({ row }) => {
      //   const label = labels.find(label => label.value === row.original.)
      const profile: Profile = row.getValue('profile')

      return (
        <Link
          href=""
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            'hover:bg-transparent hover:underline',
            'px-0',
            'justify-start'
          )}
          passHref
        >
          <div className="flex space-x-2">
            {/* {label && <Badge variant="outline">{label.label}</Badge>} */}
            <span className="max-w-[500px] truncate font-medium">
              {row.original?.profile?.name}
            </span>
          </div>
        </Link>
      )
    }
  },
  // {
  //   accessorKey: 'environment',
  //   header: ({ column }) => (
  //     <DataTableColumnHeader column={column} title="Environment" />
  //   ),
  //   cell: ({ row }) => {
  //     return (
  //       <div className="flex space-x-2 items-center">
  //         {/* {status.icon && (
  //           <status.icon className="mr-2 h-4 w-4 text-muted-foreground" />
  //         )} */}
  //         {row.original?.environments?.map(env => (
  //           <Button
  //             key={env.id}
  //             variant="outline"
  //             size="sm"
  //             className="h-8 border-dashed"
  //           >
  //             {env.name}
  //           </Button>
  //         ))}
  //       </div>
  //     )
  //   },
  //   filterFn: (row, id, value) => {
  //     return value.includes(row.getValue(id))
  //   }
  // },
  {
    id: 'actions',
    cell: ({ row }) => <DataTableRowActions row={row} />
  }
]
