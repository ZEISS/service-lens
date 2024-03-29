'use client'

import { Cross2Icon } from '@radix-ui/react-icons'
import { Table } from '@tanstack/react-table'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { DataTableViewOptions } from '@/components/data-table-view-options'
import { DataTableFacetedFilter } from './data-table-faceted-filter'

interface DataTableToolbarProps<TData> {
  table: Table<TData>
  options?: DataTableToolbarOptions
}

export interface DataTableFacetedFilterCreationAttributes {
  column?: string
  title?: string
  options: {
    label: string
    value: string
    icon?: React.ComponentType<{ className?: string }>
  }[]
}

export type DataTableToolbarOptions = {
  filterColumnName?: string
  filterColumnPlaceholder?: string
  actions?: React.ReactNode[]
  facetFilters?: DataTableFacetedFilterCreationAttributes[]
}

export function DataTableToolbar<TData>({
  table,
  options
}: DataTableToolbarProps<TData>) {
  const isFiltered = table.getState().columnFilters.length > 0

  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <Input
          placeholder="Filter workloads..."
          value={(table.getColumn('name')?.getFilterValue() as string) ?? ''}
          onChange={event =>
            table.getColumn('name')?.setFilterValue(event.target.value)
          }
          className="h-8 w-[150px] lg:w-[250px]"
        />
        {options?.facetFilters?.map(
          (filter, idx) =>
            table.getColumn(filter.column ?? '') && (
              <DataTableFacetedFilter
                key={idx}
                column={table.getColumn(filter.column ?? '')}
                title={filter.title ?? ''}
                options={filter.options}
              />
            )
        )}
        {isFiltered && (
          <Button
            variant="ghost"
            onClick={() => table.resetColumnFilters()}
            className="h-8 px-2 lg:px-3"
          >
            Reset
            <Cross2Icon className="ml-2 h-4 w-4" />
          </Button>
        )}
      </div>
      {options?.actions?.map((action, idx) => <>{action}</>)}
      <DataTableViewOptions table={table} />
    </div>
  )
}
