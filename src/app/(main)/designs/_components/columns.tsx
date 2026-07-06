import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox"
import type { TDesign } from "@/db/schema"
import type { ColumnDef } from "@tanstack/react-table"
import Link from "next/link"
import { DataTableColumnHeader } from "../../../../components/data-table/data-table-column-header"
import { DataTableRowActions } from "./data-rows-actions"

export const designColumns: ColumnDef<TDesign>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <div className="flex items-center justify-center">
        <Checkbox
          checked={table.getIsAllPageRowsSelected() || (table.getIsSomePageRowsSelected() && "indeterminate")}
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      </div>
    ),
    cell: ({ row }) => (
      <div className="flex items-center justify-center">
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      </div>
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "header",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Title" />,
    cell: ({ row }) => {
      // return <TableCellViewer item={row.original} />
      return (
        <Button variant="link" className="w-fit px-0 text-left text-foreground" asChild>
          <Link href={`/designs/${row.original.id}`}>{row.original.title}</Link>
        </Button>
      )
    },
    enableSorting: false,
  },
  {
    id: "actions",
    cell: ({ row }) => <DataTableRowActions row={row} />,
    enableSorting: false,
  },
]
