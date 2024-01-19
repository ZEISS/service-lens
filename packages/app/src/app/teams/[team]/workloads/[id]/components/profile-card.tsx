'use client'

import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import DateFormat from '@/components/date-format'
import { Profile } from '@/db/models/profile'

export type ProfileCardProps = {
  profile?: Profile
  className?: string
}

export function ProfileCard({ profile, ...props }: ProfileCardProps) {
  return (
    <Card {...props}>
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl">Profile</CardTitle>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="flex items-center justify-between">
          <div className="space-y-1">
            <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
              Name
            </h2>
            <p>{profile?.name}</p>
          </div>
        </div>
        <div className="flex items-center justify-between">
          <div className="space-y-1">
            <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
              Last updated
            </h2>
            <DateFormat date={profile?.dataValues?.updatedAt} />
          </div>
        </div>
        <Separator />
        <p>{profile?.description || 'No description provided.'}</p>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
  )
}
