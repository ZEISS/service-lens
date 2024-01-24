import Link from 'next/link'
import { headers } from 'next/headers'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'

export type NavItem = {
  link: string
  name: string
}

const nav: NavItem[] = []

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
      {nav?.map((item, idx) => (
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
