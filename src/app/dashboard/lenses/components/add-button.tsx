import { Button } from '@/components/ui/button'
import { PlusIcon } from '@radix-ui/react-icons'
import Link from 'next/link'

interface AddLensButtonProps {}

export function AddLensButton({}: AddLensButtonProps) {
  return (
    <Link href="/dashboard/lenses/new" passHref>
      <Button variant={'outline'}>
        <PlusIcon />
      </Button>
    </Link>
  )
}
