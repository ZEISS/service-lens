import { DotsHorizontalIcon } from '@radix-ui/react-icons'
import { Button } from '@/components/ui/button'
import { PlusIcon } from '@radix-ui/react-icons'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
  DropdownMenuLabel
} from '@/components/ui/dropdown-menu'
import Link from 'next/link'
import type { SolutionTemplate } from '@/db/models/solution-templates'

interface CommentActionsProps {
  templates?: SolutionTemplate[]
}

export function AddSolution({ templates }: CommentActionsProps) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant={'outline'}>
          <PlusIcon className="h-4 w-4" />
          <span className="sr-only">Create new solution</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-[160px]">
        <DropdownMenuLabel>Templates</DropdownMenuLabel>
        <Link href="/dashboard/solutions/new?template=_blank">
          <DropdownMenuItem>Blank</DropdownMenuItem>
        </Link>
        {templates?.map((tmpl, idx) => (
          <Link key={idx} href={`/dashboard/solutions/new?template=${tmpl.id}`}>
            <DropdownMenuItem>{tmpl.title}</DropdownMenuItem>
          </Link>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
