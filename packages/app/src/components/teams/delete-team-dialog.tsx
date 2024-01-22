'use client'

import {
  AlertDialog,
  AlertDialogTrigger,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogAction,
  AlertDialogCancel
} from '@/components/ui/alert-dialog'
import { Button } from '@/components/ui/button'
import { useState } from 'react'
import { useAction } from '@/trpc/client'
import { useRouter } from 'next/navigation'
import { rhfActionDeleteTeam } from '@/actions/team.action'

type DeleteTeamDialogProps = {}

export const DeleteTeamDialog = (props: DeleteTeamDialogProps) => {
  const [open, setOpen] = useState(false)

  const mutation = useAction(rhfActionDeleteTeam)
  const router = useRouter()
  const handleOnClickDelete = async () => {
    await mutation.mutate()
    setOpen(false)
    router.replace('/dashboard')
  }

  return (
    <>
      <AlertDialog open={open} onOpenChange={setOpen}>
        <AlertDialogTrigger asChild>
          <Button variant="destructive">Delete Team</Button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
            <AlertDialogDescription>
              This action cannot be undone. This will permanently delete the and
              remove the data.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <form
              action={handleOnClickDelete}
              onSubmit={event => {
                mutation.mutateAsync().then(() => router.replace('/dashboard'))
                event.preventDefault()
              }}
            >
              <Button
                variant="destructive"
                type="submit"
                disabled={mutation.status === 'loading'}
              >
                Delete
              </Button>
            </form>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
