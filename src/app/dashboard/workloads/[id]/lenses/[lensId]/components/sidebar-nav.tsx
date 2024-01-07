'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

import { cn } from '@/lib/utils'
import { Lens } from '@/db/models/lens'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

export type SidebarNavProps = {
  params: { lensId: string; id: string }
  lens?: Lens
  className?: string
}

export function SidebarNav({
  className,
  lens,
  params,
  ...props
}: SidebarNavProps) {
  const pathname = usePathname()

  return (
    <nav
      className={cn(
        'flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1',
        className
      )}
      {...props}
    >
      {lens?.pillars?.map((pillar, idx) => (
        <Card key={idx}>
          <CardHeader>
            <CardTitle>{pillar.name}</CardTitle>
          </CardHeader>
          <CardContent>
            {pillar.questions?.map(question => (
              <Link
                key={question.ref}
                href={`/dashboard/workloads/${params.id}/lenses/${params.lensId}/question/${question.id}`}
              >
                <Button
                  variant="outline"
                  className="w-full break-normal justify-start"
                >
                  {question.name}
                </Button>
              </Link>
            ))}
          </CardContent>
        </Card>
      ))}

      {/* {items.map(item => (
        <Link
          key={item.href}
          href={item.href}
          className={cn(
            buttonVariants({ variant: 'ghost' }),
            pathname === item.href
              ? 'bg-muted hover:bg-muted'
              : 'hover:bg-transparent hover:underline',
            'justify-start'
          )}
        >
          {item.title}
        </Link>
      ))} */}
    </nav>
  )
}
