"use client"

import { useActionState, useEffect, useState } from "react"

import Form from "next/form"

import type { GetProfileResponse } from "@/app/api/profiles/route"
import type { ApiComboBoxFetchFunc } from "@/components/api-combobox"
import { ApiComboBox } from "@/components/api-combobox"
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

import { profilesAssignAction } from "./profiles-assign-modal.action"

interface AddProfileModalProps {
  workloadId: string
}

export function ProfilesAssignModal({ workloadId }: AddProfileModalProps) {
  const [state, formAction, pending] = useActionState(profilesAssignAction, null)
  const [lens, setLens] = useState({ label: "", value: "" })
  const [open, setOpen] = useState(false)

  useEffect(() => {
    if (state?.success) {
      setOpen(false)
    }
  }, [state])

  const fetchItems: ApiComboBoxFetchFunc<any> = (e, setItems) => {
    fetch(`/api/profiles?search=${e}`)
      .then((res) => res.json() as Promise<GetProfileResponse>)
      .then((v) => {
        setItems(v.items.map((item) => ({ value: item.id, label: item.name })))
      })
      .catch(() => {
        setItems([])
      })
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button size="sm" variant="outline">
          Assign
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>Assign Profile</DialogTitle>
          <DialogDescription>Fill in the information below to assign a profile to the workload.</DialogDescription>
        </DialogHeader>
        <Form action={formAction} id={`assign-environment-form`}>
          <Input type="hidden" name="workloadId" value={workloadId} />
          <Input type="hidden" name="profileId" value={lens.value} />
          <FieldGroup>
            <Field data-invalid={!!state?.errors?.properties?.profileId}>
              <FieldLabel htmlFor="title">Lens</FieldLabel>
              <ApiComboBox
                className="w-full"
                selectedItem={lens}
                onSelect={(item) => {
                  setLens(item)
                }}
                onFetch={fetchItems}
              />
              {state?.errors?.properties?.profileId && (
                <FieldError>{state?.errors?.properties?.profileId.errors.pop()}</FieldError>
              )}
            </Field>
          </FieldGroup>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button type="submit" form="assign-environment-form" disabled={pending}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
