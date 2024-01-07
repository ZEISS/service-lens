import { Button } from '@/components/ui/button'
import { PlusIcon } from '@radix-ui/react-icons'
import Link from 'next/link'

interface AddProfileButtonProps {}

export function AddProfileButton({}: AddProfileButtonProps) {
  return (
    <Link href="/dashboard/profiles/new" passHref>
      <Button variant={'outline'}>
        <PlusIcon />
      </Button>
    </Link>
  )
}
