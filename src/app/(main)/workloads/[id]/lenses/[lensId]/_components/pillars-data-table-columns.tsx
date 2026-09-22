import Link from "next/link"

import type { ColumnDef } from "@tanstack/react-table"

import { DataTableColumnHeader } from "@/components/data-table/data-table-column-header"
import { Button } from "@/components/ui/button"
import type { TEnvironment } from "@/db/schema"

export const pillarsColumns = (workloadId: string, lensId: string): ColumnDef<TEnvironment>[] => [
  {
    accessorKey: "name",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Name" />,
    cell: ({ row }) => {
      return (
        <Button variant="link" className="w-fit px-0 text-left text-foreground" asChild>
          <Link href={`/workloads/${workloadId}/lenses/${lensId}/review?lensId=${row.original.id}`}>{row.original.name}</Link>
        </Button>
      )
    },
    enableSorting: false,
  },
  {
    accessorKey: "description",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Description" />,
    cell: ({ row }) => <div>{row.original.description}</div>,
    enableSorting: false,
  },
  {
    accessorKey: "createdAt",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Created At" />,
    cell: ({ row }) => <div>{row.original?.createdAt?.toDateString()}</div>,
    enableSorting: false,
  },
]
