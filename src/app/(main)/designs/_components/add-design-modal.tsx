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
import { createDesignAction } from "./add-design-modal.action"

export function AddDesignModal() {
  const [state, formAction, pending] = useActionState(createDesignAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="default">
          <Plus />
          <span>Add Design</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Design</DialogTitle>
          <DialogDescription>Fill in the information below to create a new design.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id="add-design-form">
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.title}>
              <FieldLabel htmlFor="title">Title</FieldLabel>
              <Input
                id="title"
                name="title"
                defaultValue={state?.values?.title}
                disabled={pending}
                placeholder="Fate of Atlantis"
                autoComplete="off"
              />
              <FieldDescription>Provide a concise title for your design.</FieldDescription>
              {state?.errors?.properties?.title && (
                <FieldError>{state?.errors?.properties?.title.errors.pop()}</FieldError>
              )}
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="add-design-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
