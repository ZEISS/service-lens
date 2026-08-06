import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getTagById } from "@/db/queries/tags"
import { notFound } from "next/navigation"
import { Breadcrumbs } from "../_components/breadcrumbs"

export default async function Page({ params }: { params: Promise<{ id: number }> }) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const tag = await getTagById(id)

  if (!tag) {
    return notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs tag={tag} />

      {/* Title */}
      <h1 className="scroll-m-20 text-balance font-extrabold text-4xl tracking-tight">{tag.name}</h1>

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
            <p className="mt-1 text-sm">{tag.createdAt?.toLocaleString()}</p>
          </div>

          <Separator />

          {/* Updated At */}
          <div>
            <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Last Modified</label>
            <p className="mt-1 text-sm">{tag.updatedAt?.toLocaleString()}</p>
          </div>

          {/* Deleted At */}
          {tag.deletedAt && (
            <>
              <Separator />
              <div>
                <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">Deleted</label>
                <p className="mt-1 text-sm">{tag.deletedAt?.toLocaleString()}</p>
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
            {JSON.stringify(tag, null, 2)}
          </pre>
        </CardContent>
      </Card>
    </div>
  )
}
