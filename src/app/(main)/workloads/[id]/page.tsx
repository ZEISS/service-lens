import { notFound } from "next/navigation"

import { Muted } from "@/components/typography/muted"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getWorkloadById } from "@/db/queries/workloads"

import { Breadcrumbs } from "../_components/breadcrumbs"
import { EnvironmentsDataTable } from "./_components/environments-data-table"

export default async function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const workload = await getWorkloadById(id)
  const environments = workload?.environments.map((env) => ({ ...env.environment })) ?? []

  if (!workload) {
    return notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs workload={workload} />

      {/* Title */}
      <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight text-balance">{workload.name}</h1>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Meta</CardTitle>
          <CardDescription>A brief description of the workload.</CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-sm">{workload.description || "No description provided."}</p>
        </CardContent>
      </Card>

      {/* Environments */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Environments</CardTitle>
          <CardDescription>Associated environments for this workload.</CardDescription>
        </CardHeader>
        <CardContent>
          <EnvironmentsDataTable data={environments} />
        </CardContent>
      </Card>

      {/* Timestamps */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Timeline</CardTitle>
          <CardDescription>Date and time of creation and updates</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <Muted>Created</Muted>
          <p>{workload.createdAt?.toLocaleString()}</p>

          <Separator />

          {/* Updated At */}
          <Muted>Updated</Muted>
          <p className="text-sm mt-1">{workload.updatedAt?.toLocaleString()}</p>

          {/* Deleted At */}
          {workload.deletedAt && (
            <>
              <Separator />
              <div>
                <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Deleted</label>
                <p className="text-sm mt-1">{workload.deletedAt?.toLocaleString()}</p>
              </div>
            </>
          )}
        </CardContent>
      </Card>

      {/* Raw Data (for debugging/development) */}
      <Card>
        <CardHeader>
          <CardTitle>Raw Data</CardTitle>
          <CardDescription>Technical details and raw database properties.</CardDescription>
        </CardHeader>
        <CardContent>
          <pre className="text-xs bg-muted p-3 rounded overflow-x-auto whitespace-pre-wrap">
            {JSON.stringify(workload, null, 2)}
          </pre>
        </CardContent>
      </Card>
    </div>
  )
}
