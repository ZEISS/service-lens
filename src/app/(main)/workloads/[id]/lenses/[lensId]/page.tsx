import { notFound } from "next/navigation"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getWorkloadById, getWorkloadReviewByLensId } from "@/db/queries/workloads"

import { Breadcrumbs } from "./_components/breadcrumbs"
import { PillarsDataTable } from "./_components/pillars-data-table"
import { ReviewCreateCardForm } from "./_components/review-create-card-form"

export default async function Page({ params }: { params: Promise<{ id: string; lensId: string }> }) {
  const { id: workloadId, lensId } = await params

  if (!workloadId) {
    notFound()
  }

  const workload = await getWorkloadById(workloadId)
  const lens = workload?.lenses.find((l) => l.id === lensId) ?? null
  const review = await getWorkloadReviewByLensId({ workloadId, lensId })

  if (!(workload && lens)) {
    return notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs workload={workload} lens={lens} />

      {/* Title */}
      <h1 className="scroll-m-20 text-balance font-extrabold text-4xl tracking-tight">{lens.name}</h1>

      {/* Overview */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Meta</CardTitle>
        </CardHeader>
        <CardContent>
          {/* Version */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Version</label>
            <p className="mt-1 text-sm">{lens.version}</p>
          </div>

          {/* Description */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Description</label>
            <p className="mt-1 text-sm">{lens.description}</p>
          </div>
        </CardContent>
      </Card>

      {/* Meta */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Meta</CardTitle>
        </CardHeader>
        <CardContent>
          {/* Created At */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Created</label>
            <p className="mt-1 text-sm">{lens.createdAt?.toLocaleString()}</p>
          </div>

          <Separator />

          {/* Updated At */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Last Modified</label>
            <p className="mt-1 text-sm">{lens.updatedAt?.toLocaleString()}</p>
          </div>
        </CardContent>
      </Card>

      {/* Review */}
      {review ? (null) : <ReviewCreateCardForm workloadId={workloadId} lensId={lensId} />}

      {/* Pillars */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Pillars</CardTitle>
          <CardDescription>Pillars of the lens to assess.</CardDescription>
        </CardHeader>
        <CardContent>
          <PillarsDataTable data={lens.lensPillars} lensId={lens.id} workloadId={workload.id} />
        </CardContent>
      </Card>
    </div>
  )
}
