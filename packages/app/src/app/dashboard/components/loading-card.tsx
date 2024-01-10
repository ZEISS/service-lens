'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { MagicWandIcon } from '@radix-ui/react-icons'
import { Skeleton } from '@/components/ui/skeleton'

export default function LoadingCard() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          <Skeleton className="h-4 w-[250px]" />
        </CardTitle>
        <MagicWandIcon className="h-4 w-4 text-muted-foreground" />
      </CardHeader>
      <CardContent>
        <Skeleton className="h-4 w-[250px]" />
      </CardContent>
    </Card>
  )
}
