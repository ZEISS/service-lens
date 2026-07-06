"use client"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Input } from "@/components/ui/input"
import type { Row } from "@tanstack/react-table"
import { EllipsisVertical } from "lucide-react"
import Form from "next/form"
import { useActionState } from "react"
import { deleteLensAction } from "./data-rows-actions.action"

interface DataTableRowActionsProps<TData> {
  row: Row<TData>
}

export function DataTableRowActions<TDesign>({ row }: DataTableRowActionsProps<TDesign>) {
  const { id } = row
  const [state, formAction, pending] = useActionState(deleteLensAction, null)

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="flex size-8 text-muted-foreground data-[state=open]:bg-muted" size="icon">
          <EllipsisVertical />
          <span className="sr-only">Open menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-32">
        <DropdownMenuItem>Edit</DropdownMenuItem>
        <DropdownMenuItem>Make a copy</DropdownMenuItem>
        <DropdownMenuItem>Favorite</DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem variant="destructive">
          <Form action={formAction}>
            <Input id="id" name="id" value={row.id} hidden readOnly />
            <Button type="submit" disabled={pending} variant="ghost" size="sm" className="w-full">
              Trash
            </Button>
          </Form>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
