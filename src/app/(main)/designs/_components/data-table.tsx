"use client"
"use no memo"

import * as React from "react"

import { Plus } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

import { DataTable } from "@/components/data-table/data-table"
import { DataTablePagination } from "@/components/data-table/data-table-pagination"
import { DataTableViewOptions } from "@/components/data-table/data-table-view-options"
import { withDndColumn } from "@/components/data-table/table-utils"
import type { getDesigns } from "@/db/queries/designs"
import { useDataTable } from "@/hooks/use-data-table"
import type { QueryKeys } from "@/types/data-table"
import { AddDesignModal } from "./add-design-modal"
import { designColumns } from "./columns"

interface DesignTableProps {
  promises: Promise<[Awaited<ReturnType<typeof getDesigns>>]>
  queryKeys?: Partial<QueryKeys>
}

export function DesignDataTable({ promises, queryKeys }: DesignTableProps) {
  const columns = withDndColumn(designColumns)
  const [{ data, pageCount }] = React.use(promises)

  const { table } = useDataTable({
    data,
    columns,
    pageCount,
    queryKeys,
    initialState: {
      sorting: [{ id: "createdAt", desc: true }],
      columnPinning: { right: ["actions"] },
    },
    getRowId: (row) => row.id,
    shallow: false,
    clearOnDefault: true,
  })

  return (
    <Tabs defaultValue="all" className="w-full flex-col justify-start gap-6">
      <div className="flex items-center justify-between">
        <Label htmlFor="view-selector" className="sr-only">
          View
        </Label>
        <Select defaultValue="all">
          <SelectTrigger className="flex @4xl/main:hidden w-fit" size="sm" id="view-selector">
            <SelectValue placeholder="Select a view" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All</SelectItem>
            <SelectItem value="most-recent">Most Recent</SelectItem>
          </SelectContent>
        </Select>
        <TabsList className="@4xl/main:flex hidden **:data-[slot=badge]:size-5 **:data-[slot=badge]:rounded-full **:data-[slot=badge]:bg-muted-foreground/30 **:data-[slot=badge]:px-1">
          <TabsTrigger value="all">All</TabsTrigger>
          <TabsTrigger value="most-recent" disabled={true}>
            Most Recent
          </TabsTrigger>
        </TabsList>
        <div className="flex items-center gap-2">
          <DataTableViewOptions table={table} />
          <Button variant="outline" size="sm" className="hidden lg:flex">
            <Plus />
            <span className="lg:inline">Add Section</span>
          </Button>
          <AddDesignModal />
        </div>
      </div>
      <TabsContent value="all" className="relative flex flex-col gap-4 overflow-auto">
        <div className="overflow-hidden rounded-lg border">
          <DataTable table={table} columns={columns} />
        </div>
        <DataTablePagination table={table} />
      </TabsContent>
      <TabsContent value="most-recent" className="flex flex-col">
        <div className="aspect-video w-full flex-1 rounded-lg border border-dashed" />
      </TabsContent>
    </Tabs>
  )
}
