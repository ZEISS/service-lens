import { Button } from '@/components/ui/button'
import { PlusIcon } from '@radix-ui/react-icons'
import Link from 'next/link'

interface AddWorkloadButtonProps {}

export function AddWorkloadButton({}: AddWorkloadButtonProps) {
  return (
    <Link href="/dashboard/workloads/new" passHref>
      <Button variant={'outline'}>
        <PlusIcon />
      </Button>
    </Link>
  )
}
