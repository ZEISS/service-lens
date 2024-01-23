import Link from 'next/link'
import { cn } from '@/lib/utils'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from '@/components/ui/accordion'
import { PropsWithChildren } from 'react'
import { api } from '@/trpc/server-invoker'
import { buttonVariants } from '@/components/ui/button'

export type SidebarNavProps = {
  params: { team: string; lensId: string; id: string }
}

export async function SidebarNav({
  params,
  ...props
}: PropsWithChildren<SidebarNavProps>) {
  const lens = await api.getLens.query(params.lensId)

  return (
    <nav
      className={'flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1'}
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
                  href={`/teams/${params.team}/workloads/${params.id}/lenses/${params.lensId}/question/${question.id}`}
                  className={cn(
                    buttonVariants({ variant: 'outline' }),
                    'hover:bg-transparent hover:bg-muted hover:rounded',
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
