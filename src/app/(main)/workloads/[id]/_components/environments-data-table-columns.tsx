import Link from "next/link"

import type { ColumnDef } from "@tanstack/react-table"

import { DataTableColumnHeader } from "@/components/data-table/data-table-column-header"
import { EnvironmentsTableRowActions } from "./environments-data-table-actions"
import { Button } from "@/components/ui/button"
import type { TEnvironment } from "@/db/schema"

export const environmentColumns: ColumnDef<TEnvironment>[] = [
  {
    accessorKey: "name",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Name" />,
    cell: ({ row }) => {
      return (
        <Button variant="link" className="w-fit px-0 text-left text-foreground" asChild>
          <Link href={`/environments/${row.original.id}`}>{row.original.name}</Link>
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
  {
    id: "actions",
    cell: ({ row }) => <EnvironmentsTableRowActions workloadId={row.original.id} environmentId={row.original.id} />,
    enableSorting: false,
  },
]
