"use client"

import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Plus } from "lucide-react"
import Form from "next/form"
import { useActionState } from "react"
import { createLensAction } from "./add-lens-modal.action"

export function AddLensModal() {
  const [state, formAction, pending] = useActionState(createLensAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="default">
          <Plus />
          <span>Add Lens</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Lens</DialogTitle>
        </DialogHeader>
        <Form action={formAction} id="add-lens-form" className="w-full max-w-md">
          <FieldGroup>
            <Field>
              <FieldLabel htmlFor="spec">Select specification file</FieldLabel>
              <Input id="spec" name="spec" type="file" accept=".json, .yaml, .yml" disabled={pending} />
              <FieldDescription>You are only allowed to upload JSON or YAML files.</FieldDescription>
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="add-lens-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
