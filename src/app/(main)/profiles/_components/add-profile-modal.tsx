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
import { createProfileAction } from "./add-profile-modal.action"

export function AddProfileModal() {
  const [state, formAction, pending] = useActionState(createProfileAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="default">
          <Plus />
          <span>Add Profile</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Create Profile</DialogTitle>
          <DialogDescription>Fill in the information below to create a new profile.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id="add-profile-form">
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.name}>
              <FieldLabel htmlFor="name">Name</FieldLabel>
              <Input
                id="name"
                name="name"
                defaultValue={state?.values?.name}
                disabled={pending}
                placeholder="Atlantis"
                autoComplete="off"
              />
              <FieldDescription>Provide a concise name for the profile.</FieldDescription>
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
          <Button type="submit" form="add-profile-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
