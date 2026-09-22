import { notFound } from "next/navigation"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getWorkloadById } from "@/db/queries/workloads"
import { PillarsDataTable } from "./_components/pillars-data-table"

import { Breadcrumbs } from "./_components/breadcrumbs"

export default async function Page({ params }: { params: Promise<{ id: string; lensId: string }> }) {
  const { id, lensId } = await params

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

      {/* Pillars */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Pillars</CardTitle>
          <CardDescription>Pillars of the lens to assess.</CardDescription>
          <CardAction>
            <Button variant="outline" asChild>
              <Link href={`/workloads/${workload.id}/lenses/${lens.id}/review`}>Review</Link>
            </Button>
          </CardAction>
        </CardHeader>
        <CardContent>
          <PillarsDataTable data={lens.lensPillars} lensId={lens.id} workloadId={workload.id} />
        </CardContent>
      </Card>

    </div>
  )
}
