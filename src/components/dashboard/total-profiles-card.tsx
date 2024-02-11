'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { use } from 'react'
import { api } from '@/trpc/client'
import { MagicWandIcon } from '@radix-ui/react-icons'
import DateFormat from '@/components/date-format'
import Link from 'next/link'

export interface TotalProfilesCardProps<TeamSlug = string> {
  teamSlug: TeamSlug
}

export function TotalProfilesCard({ teamSlug }: TotalProfilesCardProps) {
  const { count } = use(api.profiles.listByTeam.query({ slug: teamSlug }))

  return (
    <Link href={`/teams/${teamSlug}/profiles`} passHref>
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">Total Profiles</CardTitle>
          <MagicWandIcon className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{count}</div>
          <DateFormat className="text-xs text-muted-foreground" />
        </CardContent>
      </Card>
    </Link>
  )
}

export default TotalProfilesCard
