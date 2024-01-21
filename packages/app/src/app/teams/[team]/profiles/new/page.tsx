import { PropsWithChildren, Suspense } from 'react'
import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewProfileForm } from '@/components/dashboard/profiles/new-form'
import { LoadingSpinner } from '@/components/loading-spinner'
import { api } from '@/trpc/server-http'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page({
  params
}: PropsWithChildren<NextPageProps>) {
  const questions = await api.listProfilesQuestions.query()
  const selectedChoices = questions?.reduce(
    (prev, curr) => ({ ...prev, [curr.ref]: [] }),
    {}
  )

  return (
    <>
      <SubNav>
        <SubNavTitle>
          New Profile
          <SubNavSubtitle>
            Profiles help to provide a business context.
          </SubNavSubtitle>
        </SubNavTitle>
      </SubNav>
      <Section>
        <Suspense fallback={<LoadingSpinner />}>
          <NewProfileForm
            teamId={params.team}
            selectedChoices={selectedChoices}
            questions={questions}
          />
        </Suspense>
      </Section>
    </>
  )
}
