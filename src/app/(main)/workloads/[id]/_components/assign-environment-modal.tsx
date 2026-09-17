"use client"

import { useActionState, useState } from "react"

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
import { ApiComboBox } from "@/components/api-combobox"

import { createWorkloadAction } from "../../_components/add-workload-modal.action"

interface AddEnvironmentModalProps {
  workloadId: string
}

export function AssignEnvironmentModal({ workloadId }: AddEnvironmentModalProps) {
  const [state, formAction, pending] = useActionState(createWorkloadAction, null)
  const [environment, setEnvironment] = useState({label: "", value: ""});

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button size="sm" variant="outline">
          Assign
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Assign Environment</DialogTitle>
          <DialogDescription>Fill in the information below to create a new workload.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id={`assign-environment-form-${workloadId}`}>
          <Input type="hidden" name="workloadId" value={workloadId} />
          <Input type="hidden" name="environmentId" value={environment.value} />
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.name}>
              <FieldLabel htmlFor="title">Environment</FieldLabel>
              <ApiComboBox
                className="w-full"
                selectedItem={environment}
                url=""
                onSelect={(item) => {
                  setEnvironment(item)
                }}
              />
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
          <Button type="submit" form="add-workload-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
