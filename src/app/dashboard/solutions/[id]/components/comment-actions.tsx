'use client'

import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { z } from 'zod'
import { Button } from '@/components/ui/button'
import { useAction } from '@/trpc/client'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu'
import { SolutionComment } from '@/db/models/solution-comments'
import { rhfDeleteCommentAction } from './comment-actions.action'
import { rhfDeleteCommentActionSchema } from './comment-actions.schema'

interface CommentActionsProps {
  comment?: SolutionComment
}

export function CommentActions({ comment }: CommentActionsProps) {
  const mutation = useAction(rhfDeleteCommentAction)

  const handleDelete = async (
    data: z.infer<typeof rhfDeleteCommentActionSchema>
  ) => await mutation.mutateAsync(data)

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
        >
          <DotsHorizontalIcon className="h-4 w-4" />
          <span className="sr-only">Open menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        <DropdownMenuItem
          onClick={() => comment?.id && handleDelete(comment.id)}
        >
          Delete
          <DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
