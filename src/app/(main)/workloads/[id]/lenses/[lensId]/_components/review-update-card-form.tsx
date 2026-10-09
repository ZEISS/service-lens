"use client"

import { useActionState } from "react"

import Form from "next/form"

import { Button } from "@/components/ui/button"
import { Card, CardAction, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { InputGroup, InputGroupTextarea } from "@/components/ui/input-group"
import { updateWorkloadReviewAction } from "@/db/actions"
import type { TWorkloadReview } from "@/db/schema"

export interface ReviewUpdateCardFormProps {
  review: TWorkloadReview
}

export function ReviewUpdateCardForm({ review }: ReviewUpdateCardFormProps) {
  const [state, formAction, pending] = useActionState(updateWorkloadReviewAction, null)

  return (
    <Form action={formAction} id="update-review-form">
      <Input type="hidden" name="id" value={review.id} />
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Review</CardTitle>
          <CardDescription>The status of the review.</CardDescription>
        </CardHeader>

        <CardContent>
          <FieldGroup>
            <Field>
              <FieldLabel htmlFor="notes">Notes</FieldLabel>
              <InputGroup>
                <InputGroupTextarea
                  id="notes"
                  name="notes"
                  placeholder="Add your notes here ..."
                  rows={6}
                  className="min-h-24 resize-none"
                  defaultValue={review.notes ?? ""}
                  disabled={pending}
                />
              </InputGroup>
              <FieldDescription>Optional, add any additional notes or context here.</FieldDescription>
              {state?.errors?.properties?.notes && (
                <p className="mt-1 text-destructive text-sm">{state?.errors?.properties?.notes.errors}</p>
              )}
            </Field>
          </FieldGroup>
        </CardContent>

        <CardFooter>
          <Button variant="outline" type="submit" aria-disabled={pending}>
            Update
          </Button>
        </CardFooter>
      </Card>
    </Form>
  )
}
