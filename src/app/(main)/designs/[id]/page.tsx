import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { getDesignById } from "@/db/queries/designs"
import { CalendarIcon, ClockIcon, EditIcon, TrashIcon, UserIcon } from "lucide-react"
import Link from "next/link"
import { notFound } from "next/navigation"
import ReactMarkdown from "react-markdown"
import { Breadcrumbs } from "../_components/breadcrumbs"
import { DeleteButton } from "./delete-button"

interface DesignPageProps {
  params: Promise<{ id: string }>
}

function formatDate(date: Date | null): string {
  if (!date) return "Never"
  return new Intl.DateTimeFormat("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  }).format(new Date(date))
}

function formatRelativeTime(date: Date | null): string {
  if (!date) return "Never"
  const now = new Date()
  const diffMs = now.getTime() - new Date(date).getTime()
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  const diffHours = Math.floor(diffMs / (1000 * 60 * 60))
  const diffMinutes = Math.floor(diffMs / (1000 * 60))

  if (diffDays > 0) return `${diffDays} day${diffDays > 1 ? "s" : ""} ago`
  if (diffHours > 0) return `${diffHours} hour${diffHours > 1 ? "s" : ""} ago`
  if (diffMinutes > 0) return `${diffMinutes} minute${diffMinutes > 1 ? "s" : ""} ago`
  return "Just now"
}

function getContentStatistics(content: string | null) {
  if (!content) return null

  const wordCount = content
    .trim()
    .split(/\s+/)
    .filter((word) => word.length > 0).length
  const characterCount = content.length
  const characterCountNoSpaces = content.replace(/\s/g, "").length
  const paragraphCount = content.split(/\n\s*\n/).filter((p) => p.trim().length > 0).length
  const lineCount = content.split("\n").length
  const readingTimeMinutes = Math.max(1, Math.ceil(wordCount / 200)) // Average reading speed: 200 words per minute

  return {
    wordCount,
    characterCount,
    characterCountNoSpaces,
    paragraphCount,
    lineCount,
    readingTimeMinutes,
  }
}

export default async function DesignPage({ params }: DesignPageProps) {
  const { id } = await params

  if (!id) {
    notFound()
  }

  const design = await getDesignById(id)

  if (!design) {
    notFound()
  }

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      {/* Navigation */}
      <Breadcrumbs design={design} />

      {/* Header Section */}
      <div className="mb-8 flex flex-col gap-6 sm:flex-row sm:items-start sm:justify-between">
        <div className="flex-1">
          <h1 className="mb-2 font-bold text-4xl tracking-tight">{design.title}</h1>
          {design.description && <p className="text-lg text-muted-foreground">{design.description}</p>}
        </div>

        {/* Action Buttons */}
        <div className="flex gap-2">
          <Button asChild>
            <Link href={`/designs/${design.id}/edit`}>
              <EditIcon className="h-4 w-4" />
              Edit Design
            </Link>
          </Button>
          <DeleteButton design={design} />
        </div>
      </div>

      {/* Quick Info Cards */}
      <div className="mb-8 grid grid-cols-1 gap-4 md:grid-cols-4">
        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary/10">
                <EditIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="font-medium text-sm">Status</p>
                <p className="text-muted-foreground text-xs">{design.deletedAt ? "Deleted" : "Active"}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary/10">
                <ClockIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="font-medium text-sm">Created</p>
                <p className="text-muted-foreground text-xs">{formatRelativeTime(design.createdAt)}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary/10">
                <EditIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="font-medium text-sm">Modified</p>
                <p className="text-muted-foreground text-xs">{formatRelativeTime(design.updatedAt)}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary/10">
                <UserIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="font-medium text-sm">Content</p>
                <p className="text-muted-foreground text-xs">
                  {design.body ? `${getContentStatistics(design.body)?.wordCount || 0} words` : "No content"}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 gap-6 xl:grid-cols-4">
        {/* Main Content */}
        <div className="space-y-6 xl:col-span-3">
          {/* Design Content */}
          <Card>
            <CardHeader>
              <CardTitle>Content</CardTitle>
              <CardDescription>The main content of this design written in Markdown format.</CardDescription>
            </CardHeader>
            <CardContent>
              {design.body ? (
                <div className="prose prose-gray dark:prose-invert max-w-none">
                  <ReactMarkdown
                    components={{
                      h1: ({ children }) => (
                        <h1 className="mb-6 border-border border-b pb-2 font-bold text-3xl text-foreground">
                          {children}
                        </h1>
                      ),
                      h2: ({ children }) => (
                        <h2 className="mt-8 mb-4 font-semibold text-2xl text-foreground">{children}</h2>
                      ),
                      h3: ({ children }) => (
                        <h3 className="mt-6 mb-3 font-medium text-foreground text-xl">{children}</h3>
                      ),
                      h4: ({ children }) => (
                        <h4 className="mt-4 mb-2 font-medium text-foreground text-lg">{children}</h4>
                      ),
                      p: ({ children }) => <p className="mb-4 text-foreground leading-7">{children}</p>,
                      ul: ({ children }) => <ul className="mb-4 list-disc space-y-1 pl-6">{children}</ul>,
                      ol: ({ children }) => <ol className="mb-4 list-decimal space-y-1 pl-6">{children}</ol>,
                      li: ({ children }) => <li className="text-foreground leading-relaxed">{children}</li>,
                      blockquote: ({ children }) => (
                        <blockquote className="my-4 rounded-r-lg border-primary/50 border-l-4 bg-muted/30 py-2 pl-6 text-muted-foreground italic">
                          {children}
                        </blockquote>
                      ),
                      code: ({ children, className }) => {
                        const isInlineCode = !className
                        if (isInlineCode) {
                          return (
                            <code className="rounded-md border bg-muted px-2 py-1 font-mono text-foreground text-sm">
                              {children}
                            </code>
                          )
                        }
                        // Code block
                        const language = className?.replace("language-", "") || "text"
                        return <code className={`language-${language}`}>{children}</code>
                      },
                      pre: ({ children }) => (
                        <pre className="relative mb-4 overflow-x-auto rounded-lg border bg-muted p-4 font-mono text-sm">
                          {children}
                        </pre>
                      ),
                      a: ({ href, children }) => (
                        <a
                          href={href}
                          className="text-primary underline decoration-primary/50 underline-offset-4 transition-colors hover:text-primary/80 hover:decoration-primary"
                        >
                          {children}
                        </a>
                      ),
                      hr: () => <hr className="my-8 border-border" />,
                      table: ({ children }) => (
                        <div className="my-6 overflow-x-auto rounded-lg border border-border">
                          <table className="w-full border-collapse">{children}</table>
                        </div>
                      ),
                      thead: ({ children }) => <thead className="bg-muted/50">{children}</thead>,
                      tbody: ({ children }) => <tbody className="divide-y divide-border">{children}</tbody>,
                      th: ({ children }) => (
                        <th className="px-6 py-3 text-left font-medium text-muted-foreground text-xs uppercase tracking-wider">
                          {children}
                        </th>
                      ),
                      td: ({ children }) => (
                        <td className="whitespace-nowrap px-6 py-4 text-foreground text-sm">{children}</td>
                      ),
                    }}
                  >
                    {design.body}
                  </ReactMarkdown>
                </div>
              ) : (
                <div className="flex flex-col items-center justify-center py-12 text-center">
                  <div className="mb-4 rounded-full bg-muted p-3">
                    <EditIcon className="h-6 w-6 text-muted-foreground" />
                  </div>
                  <h3 className="mb-2 font-semibold text-lg">No content yet</h3>
                  <p className="mb-4 max-w-sm text-muted-foreground">
                    This design doesn&apos;t have any content yet. Click the edit button to add some content.
                  </p>
                  <Button asChild variant="outline">
                    <Link href={`/designs/${design.id}/edit`}>
                      <EditIcon className="h-4 w-4" />
                      Add Content
                    </Link>
                  </Button>
                </div>
              )}
            </CardContent>
          </Card>
        </div>

        {/* Sidebar - Design Properties */}
        <div className="space-y-6 xl:col-span-1">
          {/* Design Information */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <UserIcon className="h-5 w-5" />
                Design Information
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {/* ID */}
              <div>
                <label className="font-medium text-muted-foreground text-sm">ID</label>
                <p className="mt-1 break-all rounded bg-muted px-2 py-1 font-mono text-sm">{design.id}</p>
              </div>

              <Separator />

              {/* Title */}
              <div>
                <label className="font-medium text-muted-foreground text-sm">Title</label>
                <p className="mt-1 text-sm">{design.title}</p>
              </div>

              {/* Description */}
              {design.description && (
                <>
                  <Separator />
                  <div>
                    <label className="font-medium text-muted-foreground text-sm">Description</label>
                    <p className="mt-1 text-sm leading-relaxed">{design.description}</p>
                  </div>
                </>
              )}

              {/* Status */}
              <Separator />
              <div>
                <label className="font-medium text-muted-foreground text-sm">Status</label>
                <div className="mt-1">
                  {design.deletedAt ? (
                    <Badge variant="destructive">Deleted</Badge>
                  ) : (
                    <Badge variant="default">Active</Badge>
                  )}
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Timestamps */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <ClockIcon className="h-5 w-5" />
                Timeline
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {/* Created At */}
              <div>
                <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">
                  <CalendarIcon className="h-4 w-4" />
                  Created
                </label>
                <p className="mt-1 text-sm">{formatDate(design.createdAt)}</p>
                <p className="text-muted-foreground text-xs">{formatRelativeTime(design.createdAt)}</p>
              </div>

              <Separator />

              {/* Updated At */}
              <div>
                <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">
                  <EditIcon className="h-4 w-4" />
                  Last Modified
                </label>
                <p className="mt-1 text-sm">{formatDate(design.updatedAt)}</p>
                <p className="text-muted-foreground text-xs">{formatRelativeTime(design.updatedAt)}</p>
              </div>

              {/* Deleted At */}
              {design.deletedAt && (
                <>
                  <Separator />
                  <div>
                    <label className="flex items-center gap-2 font-medium text-muted-foreground text-sm">
                      <TrashIcon className="h-4 w-4" />
                      Deleted
                    </label>
                    <p className="mt-1 text-sm">{formatDate(design.deletedAt)}</p>
                    <p className="text-muted-foreground text-xs">{formatRelativeTime(design.deletedAt)}</p>
                  </div>
                </>
              )}
            </CardContent>
          </Card>

          {/* Content Statistics */}
          {design.body && (
            <Card>
              <CardHeader>
                <CardTitle>Content Statistics</CardTitle>
                <CardDescription>Analytics and insights about the design content.</CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                {(() => {
                  const stats = getContentStatistics(design.body)
                  if (!stats) return null

                  return (
                    <>
                      <div className="grid grid-cols-2 gap-4">
                        <div>
                          <label className="font-medium text-muted-foreground text-sm">Words</label>
                          <p className="font-bold text-2xl">{stats.wordCount.toLocaleString()}</p>
                        </div>
                        <div>
                          <label className="font-medium text-muted-foreground text-sm">Characters</label>
                          <p className="font-bold text-2xl">{stats.characterCount.toLocaleString()}</p>
                        </div>
                        <div>
                          <label className="font-medium text-muted-foreground text-sm">Paragraphs</label>
                          <p className="font-bold text-2xl">{stats.paragraphCount}</p>
                        </div>
                        <div>
                          <label className="font-medium text-muted-foreground text-sm">Reading Time</label>
                          <p className="font-bold text-2xl">{stats.readingTimeMinutes} min</p>
                        </div>
                      </div>

                      <Separator />

                      <div className="space-y-2">
                        <div className="flex justify-between text-sm">
                          <span className="text-muted-foreground">Lines</span>
                          <span className="font-medium">{stats.lineCount}</span>
                        </div>
                        <div className="flex justify-between text-sm">
                          <span className="text-muted-foreground">Characters (no spaces)</span>
                          <span className="font-medium">{stats.characterCountNoSpaces.toLocaleString()}</span>
                        </div>
                        <div className="flex justify-between text-sm">
                          <span className="text-muted-foreground">Average words per paragraph</span>
                          <span className="font-medium">{Math.round(stats.wordCount / stats.paragraphCount)}</span>
                        </div>
                      </div>
                    </>
                  )
                })()}
              </CardContent>
            </Card>
          )}

          {/* Raw Data (for debugging/development) */}
          <Card>
            <CardHeader>
              <CardTitle>Raw Data</CardTitle>
              <CardDescription>Technical details and raw database properties.</CardDescription>
            </CardHeader>
            <CardContent>
              <pre className="overflow-x-auto whitespace-pre-wrap rounded bg-muted p-3 text-xs">
                {JSON.stringify(design, null, 2)}
              </pre>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
