import { SettingsGeneralForm } from '@/components/teams/settings-general-form'
import { PropsWithChildren, Suspense } from 'react'
import { LoadingSpinner } from '@/components/loading-spinner'
import { Separator } from '@/components/ui/separator'
import { DeleteTeamDialog } from '@/components/teams/delete-team-dialog'

export interface NextPageProps<Team = string> {
  params: { team: Team }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default function Page({ params }: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <div>
        <h3 className="text-lg font-medium">General</h3>
        <p className="text-sm text-muted-foreground">Team Settings</p>
      </div>
      <Separator />
      <Suspense fallback={<LoadingSpinner />}>
        <SettingsGeneralForm teamId={params.team} />
      </Suspense>
      <Separator />
      <DeleteTeamDialog />
    </>
  )
}
