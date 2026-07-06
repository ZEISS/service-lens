"use client"

import * as React from "react"

import { Plus } from "lucide-react"
import type { z } from "zod"

import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useDataTable } from "@/hooks/use-data-table"

import type { QueryKeys } from "@/types/data-table"
import { DataTable as DataTableNew } from "../../../../../components/data-table/data-table"
import { DataTablePagination } from "../../../../../components/data-table/data-table-pagination"
import { DataTableViewOptions } from "../../../../../components/data-table/data-table-view-options"
import { withDndColumn } from "../../../../../components/data-table/table-utils"
import { dashboardColumns } from "./columns"
import type { sectionSchema } from "./schema"

interface DesignTableProps {
  queryKeys?: Partial<QueryKeys>
  data: z.infer<typeof sectionSchema>[]
}

export function DataTable({ data: initialData, queryKeys }: DesignTableProps) {
  const [data, setData] = React.useState(() => initialData)
  const columns = withDndColumn(dashboardColumns)
  const { table } = useDataTable({
    data,
    columns,
    queryKeys,
    pageCount: 1,
    getRowId: (row) => row.id.toString(),
    shallow: false,
    clearOnDefault: true,
  })

  return (
    <Tabs defaultValue="overview" className="w-full flex-col justify-start gap-6">
      <div className="flex items-center justify-between">
        <Label htmlFor="view-selector" className="sr-only">
          View
        </Label>
        <Select defaultValue="overview">
          <SelectTrigger className="flex @4xl/main:hidden w-fit" size="sm" id="view-selector">
            <SelectValue placeholder="Select a view" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="overview">Overview</SelectItem>
          </SelectContent>
        </Select>
        <TabsList className="@4xl/main:flex hidden **:data-[slot=badge]:size-5 **:data-[slot=badge]:rounded-full **:data-[slot=badge]:bg-muted-foreground/30 **:data-[slot=badge]:px-1">
          <TabsTrigger value="overview">Overview</TabsTrigger>
        </TabsList>
        <div className="flex items-center gap-2">
          <DataTableViewOptions table={table} />
          <Button variant="outline" size="sm">
            <Plus />
            <span className="hidden lg:inline">Add Section</span>
          </Button>
        </div>
      </div>
      <TabsContent value="outline" className="relative flex flex-col gap-4 overflow-auto">
        <div className="overflow-hidden rounded-lg border">
          <DataTableNew dndEnabled table={table} columns={columns} onReorder={setData} />
        </div>
        <DataTablePagination table={table} />
      </TabsContent>
    </Tabs>
  )
}
