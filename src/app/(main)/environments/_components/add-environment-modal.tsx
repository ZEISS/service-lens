"use client"

import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Plus } from "lucide-react"
import Form from "next/form"
import { useActionState } from "react"
import { createEnvironmentAction } from "./add-environment-modal.action"

export function AddEnvironmentModal() {
  const [state, formAction, pending] = useActionState(createEnvironmentAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="default">
          <Plus />
          <span>Add Environment</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Environment</DialogTitle>
          <DialogDescription>Fill in the information below to create a new environment.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id="add-environment-form">
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.name}>
              <FieldLabel htmlFor="title">Name</FieldLabel>
              <Input
                id="name"
                name="name"
                defaultValue={state?.values?.name}
                disabled={pending}
                placeholder="Atlantis"
                autoComplete="off"
              />
              <FieldDescription>Provide a concise title for your design.</FieldDescription>
              {state?.errors?.properties?.name && (
                <FieldError>{state?.errors?.properties?.name.errors.pop()}</FieldError>
              )}
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="add-environment-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
