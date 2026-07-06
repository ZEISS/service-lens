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
import { Field, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Plus } from "lucide-react"
import Form from "next/form"
import { useActionState } from "react"
import { createTagAction } from "./add-tag-modal.action"

export function AddTagModal() {
  const [state, formAction, pending] = useActionState(createTagAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="default">
          <Plus />
          <span>Add Tag</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Tag</DialogTitle>
          <DialogDescription>Fill in the information below to create a new tag.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id="add-tag-form">
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.name}>
              <FieldLabel htmlFor="name">Name</FieldLabel>
              <Input
                id="name"
                name="name"
                defaultValue={state?.values?.name}
                disabled={pending}
                placeholder="Dig"
                autoComplete="off"
              />
              {state?.errors?.properties?.name && (
                <FieldError>{state?.errors?.properties?.name.errors.pop()}</FieldError>
              )}
            </Field>
            <Field data-invalid={!!state?.errors?.properties?.value}>
              <FieldLabel htmlFor="value">Value</FieldLabel>
              <Input
                id="value"
                name="value"
                defaultValue={state?.values?.value}
                disabled={pending}
                placeholder="Atlantis"
                autoComplete="off"
              />
              {state?.errors?.properties?.value && (
                <FieldError>{state?.errors?.properties?.value.errors.pop()}</FieldError>
              )}
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="add-tag-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
