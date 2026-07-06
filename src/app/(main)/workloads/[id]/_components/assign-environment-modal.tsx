"use client"

import { useActionState } from "react"

import Form from "next/form"


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

import { createWorkloadAction } from "../../_components/add-workload-modal.action"

interface AddEnvironmentModalProps {
  workloadId: string
}

export function AssignEnvironmentModal({ workloadId }: AddEnvironmentModalProps) {
  const [state, formAction, pending] = useActionState(createWorkloadAction, null)

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm">Assign</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Assign Environment</DialogTitle>
          <DialogDescription>Fill in the information below to create a new workload.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id={`assign-environment-form-${workloadId}`}>
          <Input type="hidden" name="workloadId" value={workloadId} />
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.name}>
              <FieldLabel htmlFor="title">Name</FieldLabel>
              <Input
                id="name"
                name="name"
                defaultValue={state?.values?.name}
                disabled={pending}
                placeholder="Fate of Atlantis"
                autoComplete="off"
              />
              {state?.errors?.properties?.name && (
                <FieldError>{state?.errors?.properties?.name.errors.pop()}</FieldError>
              )}
            </Field>
            <Field data-invalid={!!state?.errors?.properties?.description}>
              <FieldLabel htmlFor="description">Description</FieldLabel>
              <Input
                id="description"
                name="description"
                defaultValue={state?.values?.description}
                disabled={pending}
                placeholder="Good old Mother Nature."
                autoComplete="off"
              />
              {state?.errors?.properties?.description && (
                <FieldError>{state?.errors?.properties?.description.errors.pop()}</FieldError>
              )}
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="add-workload-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
