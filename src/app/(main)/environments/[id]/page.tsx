import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { findEnvironmentById } from "@/db/queries/environments"
import { notFound } from "next/navigation"
import { Breadcrumbs } from "../_components/breadcrumbs"

export default async function Page({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const environment = await findEnvironmentById(id)

  if (!environment) {
    return notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs environment={environment} />

      {/* Title */}
      <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight text-balance">{environment.name}</h1>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Meta</CardTitle>
          <CardDescription>A brief description of the environment.</CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-sm">{environment.description || "No description provided."}</p>
        </CardContent>
      </Card>

      {/* Timestamps */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Timeline</CardTitle>
          <CardDescription>Date and time of creation and updates</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          {/* Created At */}
          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Created</label>
            <p className="text-sm mt-1">{environment.createdAt?.toLocaleString()}</p>
          </div>

          <Separator />

          {/* Updated At */}
          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Last Modified</label>
            <p className="text-sm mt-1">{environment.updatedAt?.toLocaleString()}</p>
          </div>

          {/* Deleted At */}
          {environment.deletedAt && (
            <>
              <Separator />
              <div>
                <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Deleted</label>
                <p className="text-sm mt-1">{environment.deletedAt?.toLocaleString()}</p>
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
            {JSON.stringify(environment, null, 2)}
          </pre>
        </CardContent>
      </Card>
    </div>
  )
}
