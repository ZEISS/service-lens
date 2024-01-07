import Link from 'next/link'
import { headers } from 'next/headers'

import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'

const nav = [
  {
    name: 'Dashboard',
    link: '/dashboard'
  },
  {
    name: 'Workloads',
    link: '/dashboard/workloads'
  },
  {
    name: 'Solutions',
    link: '/dashboard/solutions'
  },
  {
    name: 'Lenses',
    link: '/dashboard/lenses'
  },
  {
    name: 'Profiles',
    link: '/dashboard/profiles'
  }
]

export function MainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  const headerList = headers()
  const pathname = headerList.get('x-pathname')

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
