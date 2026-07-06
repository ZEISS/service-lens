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
      <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight text-balance">{lens.name}</h1>

      {/* Metdata */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Metadata</CardTitle>
          <CardDescription />
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">ID</label>
            <p className="text-sm mt-1">{lens.id}</p>
          </div>

          <Separator />

          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Name</label>
            <p className="text-sm mt-1">{lens.name}</p>
          </div>

          <Separator />

          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Version</label>
            <p className="text-sm mt-1">{lens.version}</p>
          </div>

          <Separator />

          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Description</label>
            <p className="text-sm mt-1">{lens.description || "No description provided."}</p>
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
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Created</label>
            <p className="text-sm mt-1">{lens.createdAt?.toLocaleString()}</p>
          </div>

          <Separator />

          {/* Updated At */}
          <div>
            <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Last Modified</label>
            <p className="text-sm mt-1">{lens.updatedAt?.toLocaleString()}</p>
          </div>

          {/* Deleted At */}
          {lens.deletedAt && (
            <>
              <Separator />
              <div>
                <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">Deleted</label>
                <p className="text-sm mt-1">{lens.deletedAt?.toLocaleString()}</p>
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
            {JSON.stringify(lens, null, 2)}
          </pre>
        </CardContent>
      </Card>
    </div>
  )
}
