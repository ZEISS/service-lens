"use client"

import { useActionState } from "react"

import Form from "next/form"

import { Button } from "@/components/ui/button"
import { Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { createWorkloadReview } from "@/db/actions"

export interface ReviewCreateCardFormProps {
  workloadId: string
  lensId: string
}

export function ReviewCreateCardForm({ workloadId, lensId }: ReviewCreateCardFormProps) {
  const [_, formAction, pending] = useActionState(createWorkloadReview, null)

  return (
    <Form action={formAction} id="create-review-form">
      <input type="hidden" name="workloadId" value={workloadId} />
      <input type="hidden" name="lensId" value={lensId} />
      <input type="hidden" name="notes" value="" />
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Review</CardTitle>
          <CardDescription>The status of the review.</CardDescription>
          <CardAction>
            <Button variant="outline" type="submit" disabled={pending}>
              Start Review
            </Button>
          </CardAction>
        </CardHeader>

        <CardContent>
          <div className="flex items-center gap-2">
            <Label>No review started.</Label>
          </div>
        </CardContent>
      </Card>
    </Form>
  )
}
