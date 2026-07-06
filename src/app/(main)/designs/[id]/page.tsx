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
      <div className="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-6 mb-8">
        <div className="flex-1">
          <h1 className="text-4xl font-bold tracking-tight mb-2">{design.title}</h1>
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
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-primary/10 rounded-lg flex items-center justify-center">
                <EditIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="text-sm font-medium">Status</p>
                <p className="text-xs text-muted-foreground">{design.deletedAt ? "Deleted" : "Active"}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-primary/10 rounded-lg flex items-center justify-center">
                <ClockIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="text-sm font-medium">Created</p>
                <p className="text-xs text-muted-foreground">{formatRelativeTime(design.createdAt)}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-primary/10 rounded-lg flex items-center justify-center">
                <EditIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="text-sm font-medium">Modified</p>
                <p className="text-xs text-muted-foreground">{formatRelativeTime(design.updatedAt)}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-4">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-primary/10 rounded-lg flex items-center justify-center">
                <UserIcon className="h-4 w-4 text-primary" />
              </div>
              <div>
                <p className="text-sm font-medium">Content</p>
                <p className="text-xs text-muted-foreground">
                  {design.body ? `${getContentStatistics(design.body)?.wordCount || 0} words` : "No content"}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-4 gap-6">
        {/* Main Content */}
        <div className="xl:col-span-3 space-y-6">
          {/* Design Content */}
          <Card>
            <CardHeader>
              <CardTitle>Content</CardTitle>
              <CardDescription>The main content of this design written in Markdown format.</CardDescription>
            </CardHeader>
            <CardContent>
              {design.body ? (
                <div className="prose prose-gray max-w-none dark:prose-invert">
                  <ReactMarkdown
                    components={{
                      h1: ({ children }) => (
                        <h1 className="text-3xl font-bold mb-6 text-foreground border-b border-border pb-2">
                          {children}
                        </h1>
                      ),
                      h2: ({ children }) => (
                        <h2 className="text-2xl font-semibold mb-4 mt-8 text-foreground">{children}</h2>
                      ),
                      h3: ({ children }) => (
                        <h3 className="text-xl font-medium mb-3 mt-6 text-foreground">{children}</h3>
                      ),
                      h4: ({ children }) => (
                        <h4 className="text-lg font-medium mb-2 mt-4 text-foreground">{children}</h4>
                      ),
                      p: ({ children }) => <p className="mb-4 leading-7 text-foreground">{children}</p>,
                      ul: ({ children }) => <ul className="list-disc pl-6 mb-4 space-y-1">{children}</ul>,
                      ol: ({ children }) => <ol className="list-decimal pl-6 mb-4 space-y-1">{children}</ol>,
                      li: ({ children }) => <li className="text-foreground leading-relaxed">{children}</li>,
                      blockquote: ({ children }) => (
                        <blockquote className="border-l-4 border-primary/50 pl-6 py-2 my-4 italic text-muted-foreground bg-muted/30 rounded-r-lg">
                          {children}
                        </blockquote>
                      ),
                      code: ({ children, className }) => {
                        const isInlineCode = !className
                        if (isInlineCode) {
                          return (
                            <code className="bg-muted px-2 py-1 rounded-md text-sm font-mono text-foreground border">
                              {children}
                            </code>
                          )
                        }
                        // Code block
                        const language = className?.replace("language-", "") || "text"
                        return <code className={`language-${language}`}>{children}</code>
                      },
                      pre: ({ children }) => (
                        <pre className="bg-muted p-4 rounded-lg overflow-x-auto mb-4 font-mono text-sm border relative">
                          {children}
                        </pre>
                      ),
                      a: ({ href, children }) => (
                        <a
                          href={href}
                          className="text-primary hover:text-primary/80 underline underline-offset-4 decoration-primary/50 hover:decoration-primary transition-colors"
                        >
                          {children}
                        </a>
                      ),
                      hr: () => <hr className="border-border my-8" />,
                      table: ({ children }) => (
                        <div className="overflow-x-auto my-6 rounded-lg border border-border">
                          <table className="w-full border-collapse">{children}</table>
                        </div>
                      ),
                      thead: ({ children }) => <thead className="bg-muted/50">{children}</thead>,
                      tbody: ({ children }) => <tbody className="divide-y divide-border">{children}</tbody>,
                      th: ({ children }) => (
                        <th className="px-6 py-3 text-left text-xs font-medium text-muted-foreground uppercase tracking-wider">
                          {children}
                        </th>
                      ),
                      td: ({ children }) => (
                        <td className="px-6 py-4 text-sm text-foreground whitespace-nowrap">{children}</td>
                      ),
                    }}
                  >
                    {design.body}
                  </ReactMarkdown>
                </div>
              ) : (
                <div className="flex flex-col items-center justify-center py-12 text-center">
                  <div className="rounded-full bg-muted p-3 mb-4">
                    <EditIcon className="h-6 w-6 text-muted-foreground" />
                  </div>
                  <h3 className="font-semibold text-lg mb-2">No content yet</h3>
                  <p className="text-muted-foreground mb-4 max-w-sm">
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
        <div className="xl:col-span-1 space-y-6">
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
                <label className="text-sm font-medium text-muted-foreground">ID</label>
                <p className="text-sm font-mono bg-muted px-2 py-1 rounded mt-1 break-all">{design.id}</p>
              </div>

              <Separator />

              {/* Title */}
              <div>
                <label className="text-sm font-medium text-muted-foreground">Title</label>
                <p className="text-sm mt-1">{design.title}</p>
              </div>

              {/* Description */}
              {design.description && (
                <>
                  <Separator />
                  <div>
                    <label className="text-sm font-medium text-muted-foreground">Description</label>
                    <p className="text-sm mt-1 leading-relaxed">{design.description}</p>
                  </div>
                </>
              )}

              {/* Status */}
              <Separator />
              <div>
                <label className="text-sm font-medium text-muted-foreground">Status</label>
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
                <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">
                  <CalendarIcon className="h-4 w-4" />
                  Created
                </label>
                <p className="text-sm mt-1">{formatDate(design.createdAt)}</p>
                <p className="text-xs text-muted-foreground">{formatRelativeTime(design.createdAt)}</p>
              </div>

              <Separator />

              {/* Updated At */}
              <div>
                <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">
                  <EditIcon className="h-4 w-4" />
                  Last Modified
                </label>
                <p className="text-sm mt-1">{formatDate(design.updatedAt)}</p>
                <p className="text-xs text-muted-foreground">{formatRelativeTime(design.updatedAt)}</p>
              </div>

              {/* Deleted At */}
              {design.deletedAt && (
                <>
                  <Separator />
                  <div>
                    <label className="text-sm font-medium text-muted-foreground flex items-center gap-2">
                      <TrashIcon className="h-4 w-4" />
                      Deleted
                    </label>
                    <p className="text-sm mt-1">{formatDate(design.deletedAt)}</p>
                    <p className="text-xs text-muted-foreground">{formatRelativeTime(design.deletedAt)}</p>
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
                          <label className="text-sm font-medium text-muted-foreground">Words</label>
                          <p className="text-2xl font-bold">{stats.wordCount.toLocaleString()}</p>
                        </div>
                        <div>
                          <label className="text-sm font-medium text-muted-foreground">Characters</label>
                          <p className="text-2xl font-bold">{stats.characterCount.toLocaleString()}</p>
                        </div>
                        <div>
                          <label className="text-sm font-medium text-muted-foreground">Paragraphs</label>
                          <p className="text-2xl font-bold">{stats.paragraphCount}</p>
                        </div>
                        <div>
                          <label className="text-sm font-medium text-muted-foreground">Reading Time</label>
                          <p className="text-2xl font-bold">{stats.readingTimeMinutes} min</p>
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
              <pre className="text-xs bg-muted p-3 rounded overflow-x-auto whitespace-pre-wrap">
                {JSON.stringify(design, null, 2)}
              </pre>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
