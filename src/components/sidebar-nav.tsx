import Link from 'next/link'

import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import { headers } from 'next/headers'

export type SidebarItem = {
  href: string
  title: string
}

export interface SidebarNavProps extends React.HTMLAttributes<HTMLElement> {
  items: SidebarItem[]
}

export function SidebarNav({ className, items, ...props }: SidebarNavProps) {
  const headerList = headers()
  const pathname = headerList.get('x-pathname')

  return (
    <aside
      className={cn(
        'flex space-x-2 -mx-4 lg:w-1/5 lg:flex-col lg:space-x-0 lg:space-y-1',
        className
      )}
      {...props}
    >
      {items.map(item => (
        <Link
          key={item.href}
          href={item.href}
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            pathname === item.href
              ? 'bg-muted hover:bg-muted'
              : 'hover:bg-transparent hover:bg-muted hover:rounded',
            'justify-start'
          )}
        >
          {item.title}
        </Link>
      ))}
    </aside>
  )
}
