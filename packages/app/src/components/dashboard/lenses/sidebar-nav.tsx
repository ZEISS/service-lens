'use client'

import Link from 'next/link'
import { cn } from '@/lib/utils'
import { Lens } from '@/db/models/lens'
import { usePathname } from 'next/navigation'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from '@/components/ui/accordion'
import { PropsWithChildren } from 'react'
import { buttonVariants } from '@/components/ui/button'

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
}: PropsWithChildren<SidebarNavProps>) {
  const pathname = usePathname()

  return (
    <nav
      className={cn(
        'flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1',
        className
      )}
      {...props}
    >
      <Accordion type="single" collapsible className="w-full">
        {lens?.pillars?.map((pillar, idx) => (
          <AccordionItem key={idx} value="item-1">
            <AccordionTrigger>{pillar.name}</AccordionTrigger>
            <AccordionContent>
              {pillar.questions?.map(question => (
                <Link
                  key={question.ref}
                  href={`/dashboard/workloads/${params.id}/lenses/${params.lensId}/question/${question.id}`}
                  className={cn(
                    buttonVariants({ variant: 'outline' }),
                    pathname === question.ref
                      ? 'bg-muted hover:bg-muted'
                      : 'hover:bg-transparent hover:bg-muted hover:rounded',
                    'whitespace-normal p-4 justify-start h-full'
                  )}
                >
                  {question.name}
                </Link>
              ))}
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
    </nav>
  )
}
