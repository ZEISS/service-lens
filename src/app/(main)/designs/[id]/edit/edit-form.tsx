"use client"

import { useState } from "react"

import dynamic from "next/dynamic"
import Link from "next/link"
import { useRouter } from "next/navigation"

import { zodResolver } from "@hookform/resolvers/zod"
import { FormProvider, useForm } from "react-hook-form"
import { toast } from "sonner"
import { z } from "zod"

import { Button } from "@/components/ui/button"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import type { TDesign } from "@/db/schemas/design"

import { type UpdateDesignFormData, updateDesignAction } from "./actions"

// Dynamically import the markdown editor to avoid SSR issues
const MDEditor = dynamic(() => import("@uiw/react-md-editor"), { ssr: false })

// Form schema
const editDesignSchema = z.object({
  title: z.string().min(5, "Title must be at least 5 characters").max(255, "Title must be at most 255 characters"),
  description: z.string().max(1024, "Description must be at most 1024 characters").optional().nullable(),
  body: z.string().optional().nullable(),
})

interface EditDesignFormProps {
  design: TDesign
}

export function EditDesignForm({ design }: EditDesignFormProps) {
  const router = useRouter()
  const [saving, setSaving] = useState(false)
  const [markdownValue, setMarkdownValue] = useState<string>(design.body || "")

  const form = useForm<UpdateDesignFormData>({
    resolver: zodResolver(editDesignSchema),
    defaultValues: {
      title: design.title,
      description: design.description || "",
      body: design.body || "",
    },
  })

  const {
    handleSubmit,
    setValue,
    formState: { errors },
  } = form

  // Handle markdown editor changes
  const handleMarkdownChange = (value?: string) => {
    const newValue = value || ""
    setMarkdownValue(newValue)
    setValue("body", newValue)
  }

  // Handle form submission
  const onSubmit = async (data: UpdateDesignFormData) => {
    setSaving(true)
    try {
      const result = await updateDesignAction(design.id, data)

      if (result.success) {
        toast.success("Design updated successfully")
        router.push(`/designs/${design.id}`)
      } else {
        toast.error(result.error || "Failed to update design")
      }
    } catch (error) {
      console.error("Error updating design:", error)
      toast.error("Failed to update design")
    } finally {
      setSaving(false)
    }
  }

  return (
    <FormProvider {...form}>
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="title">Title</FieldLabel>
            <Input id="title" {...form.register("title")} placeholder="Enter design title" disabled={saving} />
            <FieldDescription>Provide a clear and concise title for your design.</FieldDescription>
            {errors.title && <p className="text-sm text-destructive mt-1">{errors.title.message}</p>}
          </Field>

          <Field>
            <FieldLabel htmlFor="description">Description</FieldLabel>
            <Textarea
              id="description"
              {...form.register("description")}
              placeholder="Enter a brief description (optional)"
              rows={3}
              disabled={saving}
            />
            <FieldDescription>Optional description to provide context about your design.</FieldDescription>
            {errors.description && <p className="text-sm text-destructive mt-1">{errors.description.message}</p>}
          </Field>

          <Field>
            <FieldLabel htmlFor="body">Content</FieldLabel>
            <div className="mt-2">
              <MDEditor
                value={markdownValue}
                onChange={handleMarkdownChange}
                preview="edit"
                height={400}
                data-color-mode="light"
              />
            </div>
            <FieldDescription>
              Write your design content using Markdown syntax. You can include headers, lists, code blocks, links, and
              more.
            </FieldDescription>
            {errors.body && <p className="text-sm text-destructive mt-1">{errors.body.message}</p>}
          </Field>
        </FieldGroup>

        <div className="flex items-center gap-4 pt-6">
          <Button type="submit" disabled={saving} className="min-w-30">
            {saving ? "Saving..." : "Save Changes"}
          </Button>
          <Button variant="outline" type="button" asChild disabled={saving}>
            <Link href={`/designs/${design.id}`}>Cancel</Link>
          </Button>
        </div>
      </form>
    </FormProvider>
  )
}

// Ensure the module is properly exported
export default EditDesignForm
