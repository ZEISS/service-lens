import { Separator } from '@/components/ui/separator'
import { PropsWithChildren, Suspense } from 'react'
import { LoadingSpinner } from '@/components/loading-spinner'
import { SettingsMembersForm } from '@/components/teams/settings-members-form'
import { AddMemberForm } from '@/components/teams/add-member-form'

export interface NextPageProps<Team = string> {
  params: { team: Team }
}

export default function Page({ params }: NextPageProps) {
  return (
    <>
      <div>
        <h3 className="text-lg font-medium">Members</h3>
        <p className="text-sm text-muted-foreground">
          Manage the members of your team.
        </p>
      </div>
      <Separator />

      <Suspense fallback={<LoadingSpinner />}>
        <SettingsMembersForm teamId={params.team} />
      </Suspense>

      <Suspense fallback={<LoadingSpinner />}>
        <AddMemberForm />
      </Suspense>
    </>
  )
}
