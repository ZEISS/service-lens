import { notFound } from "next/navigation"
import { getDesignById } from "@/db/queries/designs"
import EditDesignForm from "./edit-form"

interface EditPageProps {
  params: Promise<{ id: string }>
}

export default async function EditPage({ params }: EditPageProps) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const design = await getDesignById(id)

  if (!design) {
    notFound()
  }

  return (
    <div className="w-full max-w-4xl mx-auto p-6">
      <div className="mb-8">
        <h1 className="text-3xl font-bold tracking-tight">Edit Design</h1>
        <p className="text-muted-foreground mt-2">
          Update your design&apos;s title, description, and content using the markdown editor below.
        </p>
      </div>

      <EditDesignForm design={design} />
    </div>
  )
}
