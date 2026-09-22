import { notFound } from "next/navigation"
import type { SearchParams } from "@/types"
import { JumpTo } from "./_components/jump-to"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Field, FieldLabel } from "@/components/ui/field"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"

import { getWorkloadById } from "@/db/queries/workloads"

import { Breadcrumbs } from "./_components/breadcrumbs"
import { ArrowLeftIcon, ArrowRightIcon } from "lucide-react"

import { Button } from "@/components/ui/button"
import { ButtonGroup } from "@/components/ui/button-group"
import { lv } from "date-fns/locale"

interface ReviewPageProps {
  params: Promise<{ id: string; lensId: string }>
  searchParams: Promise<SearchParams>
}

const items = [
  { label: "Select a fruit", value: null },
  { label: "Apple", value: "apple" },
  { label: "Banana", value: "banana" },
  { label: "Blueberry", value: "blueberry" },
  { label: "Grapes", value: "grapes" },
  { label: "Pineapple", value: "pineapple" },
]

export default async function Page({ params, searchParams }: ReviewPageProps) {
  const { id, lensId } = await params
  const { pillarId, questionId } = await searchParams

  if (!id) {
    notFound()
  }

  const workload = await getWorkloadById(id)
  const lens = workload?.lenses.find((l) => l.id === lensId) ?? null

  if (!(workload && lens)) {
    return notFound()
  }

  return (
      <div className="@container/main flex flex-col gap-4 md:gap-6">
        {/* Navigation */}
      <Breadcrumbs workload={workload} lens={lens} />

      <ButtonGroup>
        <ButtonGroup className="hidden sm:flex">
          <Button variant="outline" size="icon" aria-label="Previous">
            <ArrowLeftIcon />
          </Button>
          <Button variant="outline" size="icon" aria-label="Next">
            <ArrowRightIcon />
          </Button>
        </ButtonGroup>
        <ButtonGroup>
          <Button variant="outline">Revert</Button>
          <JumpTo pillars={lens.lensPillars} />
        </ButtonGroup>
        <ButtonGroup>
          <Button variant="outline">Save</Button>
          <Button variant="outline">Save & Next</Button>
        </ButtonGroup>
      </ButtonGroup>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2"></CardTitle>
            <CardDescription>{workload.description || "No description provided."}</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex items-center space-x-2">
              <Switch id="airplane-mode" /><Label htmlFor="airplane-mode">This question does not apply to the workload.</Label>
            </div>
          </CardContent>
      </Card>

      {/* Choices */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Answers</CardTitle>
          <CardDescription>Select the answers that best fit the workload.</CardDescription>
        </CardHeader>
        <CardContent>
        </CardContent>
      </Card>

        {/* Notes */}
        <Card>
          <CardContent>
            <Field data-disabled>
              <FieldLabel htmlFor="textarea-disabled">Notes (optional)</FieldLabel>
              <Textarea id="textarea-disabled" placeholder="Add your notes here." />
            </Field>
          </CardContent>
        </Card>
    </div>
  )
}
