import Link from 'next/link'
import { headers } from 'next/headers'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'

export interface MainNavProps extends React.HTMLAttributes<HTMLElement> {
  teamId: string
}

export function MainNav({ className, teamId, ...props }: MainNavProps) {
  const headerList = headers()
  const pathname = headerList.get('x-pathname')

  const nav = [
    {
      name: 'Dashboard',
      link: `/team/${teamId}/dashboard`
    },
    {
      name: 'Workloads',
      link: `/team/${teamId}/workloads`
    },
    {
      name: 'Solutions',
      link: `/team/${teamId}/solutions`
    },
    {
      name: 'Lenses',
      link: `/team/${teamId}/lenses`
    },
    {
      name: 'Profiles',
      link: `/team/${teamId}/profiles`
    }
  ]

  return (
    <nav
      className={cn('flex items-center space-x-4 lg:space-x-6', className)}
      {...props}
    >
      {nav.map((item, idx) => (
        <Link
          key={idx}
          href={item.link}
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            'hover:bg-transparent hover:bg-muted hover:rounded',
            'justify-start'
          )}
        >
          {item.name}
        </Link>
      ))}
    </nav>
  )
}
