'use client'

import { Button } from '@/components/ui/button'
import { useAction } from '@/trpc/client'
import { rhfActionPushlishLens } from '@/app/dashboard/lenses/actions/lens.action'

interface PublishButtonProps {
  lensId: string
}

export function PublishButton({ lensId }: PublishButtonProps) {
  const mutation = useAction(rhfActionPushlishLens)
  const handleOnClickPublish = async () => await mutation.mutateAsync(lensId)

  return (
    <form action={rhfActionPushlishLens} onSubmit={handleOnClickPublish}>
      <Button variant="outline" type="submit">
        Publish
      </Button>
    </form>
  )
}
