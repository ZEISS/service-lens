import { Separator } from '@/components/ui/separator'
import { SettingsGeneralForm } from '@/components/teams/settings-general-form'
import { PropsWithChildren, Suspense } from 'react'
import { LoadingSpinner } from '@/components/loading-spinner'

export interface NextPageProps<Team = string> {
  params: { team: Team }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default function Page({ params }: PropsWithChildren<NextPageProps>) {
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
        <SettingsGeneralForm teamId={params.team} />
      </Suspense>
    </>
  )
}
