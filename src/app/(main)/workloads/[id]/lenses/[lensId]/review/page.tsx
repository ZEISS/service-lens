import { notFound } from "next/navigation"

import { ArrowLeftIcon, ArrowRightIcon } from "lucide-react"

import { Button } from "@/components/ui/button"
import { ButtonGroup } from "@/components/ui/button-group"
import { Card, CardContent } from "@/components/ui/card"
import { Field, FieldLabel } from "@/components/ui/field"
import { Textarea } from "@/components/ui/textarea"
import { getWorkloadById } from "@/db/queries/workloads"
import type { SearchParams } from "@/types"

import { Breadcrumbs } from "./_components/breadcrumbs"
import { JumpTo } from "./_components/jump-to"
import { QuestionForm } from "./_components/questions-form"

interface ReviewPageProps {
  params: Promise<{ id: string; lensId: string }>
  searchParams: Promise<SearchParams>
}

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

  const pillar = lens.lensPillars.find((p) => p.id === pillarId) ?? null
  const question = pillar?.questions.find((q) => q.id === questionId) ?? pillar?.questions[0]

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs workload={workload} lens={lens} />

      {/* */}
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

      {/* Question Form */}
      <QuestionForm question={question} />
    </div>
  )
}
