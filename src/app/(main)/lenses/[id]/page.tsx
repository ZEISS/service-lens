import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getLensById } from "@/db/queries/lenses"
import { notFound } from "next/navigation"
import { Breadcrumbs } from "../_components/breadcrumbs"

interface LensPageProps {
  params: Promise<{ id: string }>
}

export default async function LensPage({ params }: LensPageProps) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const lens = await getLensById(id)

  if (!lens) {
    notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs lens={lens} />

      {/* Title */}
      <h1 className="scroll-m-20 text-balance font-extrabold text-4xl tracking-tight">{lens.name}</h1>

      {/* Metdata */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Metadata</CardTitle>
          <CardDescription />
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">ID</label>
            <p className="mt-1 text-sm">{lens.id}</p>
          </div>

          <Separator />

          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Name</label>
            <p className="mt-1 text-sm">{lens.name}</p>
          </div>

          <Separator />

          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Version</label>
            <p className="mt-1 text-sm">{lens.version}</p>
          </div>

          <Separator />

          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Description</label>
            <p className="mt-1 text-sm">{lens.description || "No description provided."}</p>
          </div>
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
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Created</label>
            <p className="mt-1 text-sm">{lens.createdAt?.toLocaleString()}</p>
          </div>

          <Separator />

          {/* Updated At */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Last Modified</label>
            <p className="mt-1 text-sm">{lens.updatedAt?.toLocaleString()}</p>
          </div>

          {/* Deleted At */}
          {lens.deletedAt && (
            <>
              <Separator />
              <div>
                <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Deleted</label>
                <p className="mt-1 text-sm">{lens.deletedAt?.toLocaleString()}</p>
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
          <pre className="overflow-x-auto whitespace-pre-wrap rounded bg-muted p-3 text-xs">
            {JSON.stringify(lens, null, 2)}
          </pre>
        </CardContent>
      </Card>
    </div>
  )
}
