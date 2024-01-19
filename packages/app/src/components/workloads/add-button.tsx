import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { PropsWithChildren } from 'react'

export interface AddWorkloadButtonProp {
  teamSlug: string
}

export function AddWorkloadButton({
  teamSlug
}: PropsWithChildren<AddWorkloadButtonProp>) {
  return (
    <Link href={`/teams/${teamSlug}/workloads/new`} passHref>
      <Button variant={'outline'}>Add Workload</Button>
    </Link>
  )
}
