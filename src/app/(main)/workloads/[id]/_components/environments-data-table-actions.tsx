"use client"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import Link from "next/link"
import { EllipsisVertical } from "lucide-react"
import Form from "next/form"
import { useActionState } from "react"
import { removeEnvironmentAction } from "./environments-data-table-actions.action"

interface DataTableRowActionsProps {
  workloadId: string
  environmentId: string
}

export function EnvironmentsTableRowActions({ workloadId, environmentId }: DataTableRowActionsProps) {
  const [_state, formAction, pending] = useActionState(removeEnvironmentAction, null)

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="flex size-8 text-muted-foreground data-[state=open]:bg-muted" size="icon">
          <EllipsisVertical />
          <span className="sr-only">Open menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-32">
        <DropdownMenuItem>
          <Link href={`/environments/${environmentId}`}>Edit</Link>
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem variant="destructive">
          <Form action={formAction}>
            <input type="hidden" id="workloadId" name="workloadId" value={workloadId} />
            <input type="hidden" id="environmentId" name="environmentId" value={environmentId} />
            <button type="submit" disabled={pending}>
              Remove
            </button>
          </Form>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
