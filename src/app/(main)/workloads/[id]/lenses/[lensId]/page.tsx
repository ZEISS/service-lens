import { notFound } from "next/navigation"

import { Button } from "@/components/ui/button"
import { Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { getWorkloadWithReviewsForLens } from "@/db/queries/workloads"

import { Breadcrumbs } from "./_components/breadcrumbs"
import { PillarsDataTable } from "./_components/pillars-data-table"
import { ReviewCreateCardForm } from "./_components/review-create-card-form"
import { ReviewUpdateCardForm } from "./_components/review-update-card-form"

export default async function Page({ params }: { params: Promise<{ id: string; lensId: string }> }) {
  const { id, lensId } = await params

  if (!id) {
    notFound()
  }

  const workload = await getWorkloadWithReviewsForLens({ id, lensId })
  const lens = workload?.lenses.pop()
  const review = workload?.reviews.pop()

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs workload={workload} lens={lens} />

      {/* Title */}
      <h1 className="scroll-m-20 text-balance font-extrabold text-4xl tracking-tight">{lens?.name}</h1>

      {/* Meta */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Meta</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
            {/* Version */}
            <div>
              <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Version</label>
              <p className="mt-1 text-sm">{lens?.version}</p>
            </div>

            {/* Description */}
            <div>
              <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Description</label>
              <p className="mt-1 text-sm">{lens?.description}</p>
            </div>

            {/* Created At */}
            <div>
              <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Created</label>
              <p className="mt-1 text-sm">{lens?.createdAt?.toLocaleString()}</p>
            </div>

            {/* Updated At */}
            <div>
              <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Last Modified</label>
              <p className="mt-1 text-sm">{lens?.updatedAt?.toLocaleString()}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Review */}
      {review ? (
        <ReviewUpdateCardForm review={review} />
      ) : (
        <ReviewCreateCardForm workloadId={id} lensId={lensId} />
      )}

      {/* Pillars */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Pillars</CardTitle>
          <CardDescription>Pillars of the lens to assess.</CardDescription>
          <CardAction>
            <Button variant="outline" type="submit">
              Continue
            </Button>
          </CardAction>
        </CardHeader>
        <CardContent>
          {(lens?.lensPillars && workload) && <PillarsDataTable data={lens.lensPillars} lensId={lens.id} workloadId={workload.id} />}
        </CardContent>
      </Card>
    </div>
  )
}
