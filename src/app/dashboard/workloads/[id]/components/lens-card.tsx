'use client'

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import Link from 'next/link'
import { Separator } from '@/components/ui/separator'
import DateFormat from '@/components/date-format'
import { Lens } from '@/db/models/lens'

export type LensCardProps = {
  lens?: Lens
  workloadId?: string
  className?: string
}

export function LensCard({ workloadId, lens, ...props }: LensCardProps) {
  return (
    <Card {...props}>
      <CardHeader className="space-y-1">
        <Link href={`/dashboard/workloads/${workloadId}/lenses/${lens?.id}`}>
          <CardTitle className="text-2xl">{lens?.name}</CardTitle>
        </Link>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="flex items-center justify-between">
          <div className="space-y-1">
            <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
              Last updated
            </h2>
            <DateFormat date={lens?.dataValues?.updatedAt} />
          </div>
        </div>
        <Separator />
        <p>{lens?.description || 'No description provided.'}</p>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
  )
}
