import Link from 'next/link'
import { headers, cookies } from 'next/headers'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import { PropsWithChildren } from 'react'

export interface MainNavProps extends React.HTMLAttributes<HTMLElement> {
  scope?: string
}

export function MainNav({
  className,
  scope,
  ...props
}: PropsWithChildren<MainNavProps>) {
  const headerList = headers()
  const pathname = headerList.get('x-pathname')

  const nav = [
    {
      name: 'Dashboard',
      link: `/teams/${scope}/dashboard`
    },
    {
      name: 'Workloads',
      link: `/teams/${scope}/workloads`
    },
    {
      name: 'Solutions',
      link: `/teams/${scope}/solutions`
    },
    {
      name: 'Lenses',
      link: `/teams/${scope}/lenses`
    },
    {
      name: 'Profiles',
      link: `/teams/${scope}/profiles`
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
