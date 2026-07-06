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
import type { getProfiles } from "@/db/queries/profiles"
import { useDataTable } from "@/hooks/use-data-table"
import type { QueryKeys } from "@/types/data-table"
import { AddProfileModal } from "./add-profile-modal"
import { profileColumns } from "./columns"

interface ProfileTableProps {
  promises: Promise<[Awaited<ReturnType<typeof getProfiles>>]>
  queryKeys?: Partial<QueryKeys>
}

export function ProfileDataTable({ promises, queryKeys }: ProfileTableProps) {
  const columns = withDndColumn(profileColumns)
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
          </SelectContent>
          <SelectContent>
            <SelectItem value="recent">Most Recent</SelectItem>
          </SelectContent>
        </Select>
        <TabsList className="@4xl/main:flex hidden **:data-[slot=badge]:size-5 **:data-[slot=badge]:rounded-full **:data-[slot=badge]:bg-muted-foreground/30 **:data-[slot=badge]:px-1">
          <TabsTrigger value="all">All</TabsTrigger>
          <TabsTrigger value="recent">Most Recent</TabsTrigger>
        </TabsList>
        <div className="flex items-center gap-2">
          <DataTableViewOptions table={table} />
          <Button variant="outline" size="sm">
            <Plus />
            <span className="hidden lg:inline">Add Section</span>
          </Button>
          <AddProfileModal />
        </div>
      </div>
      <TabsContent value="all" className="relative flex flex-col gap-4 overflow-auto">
        <div className="overflow-hidden rounded-lg border">
          <DataTable table={table} columns={columns} />
        </div>
        <DataTablePagination table={table} />
      </TabsContent>
    </Tabs>
  )
}
