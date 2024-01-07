import { Suspense, use, useDeferredValue } from 'react'
import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewProfileForm } from '@/components/dashboard/profiles/new-form'
import { LoadingSpinner } from '@/components/loading-spinner'
import { api } from '@/trpc/server-http'

async function ServerLoader() {
  const questions = await api.listProfilesQuestions.query()
  const selectedChoices = questions?.reduce(
    (prev, curr) => ({ ...prev, [curr.ref]: [] }),
    {}
  )

  return (
    <NewProfileForm selectedChoices={selectedChoices} questions={questions} />
  )
}

export default async function Page() {
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
          <ServerLoader />
        </Suspense>
      </Section>
    </>
  )
}
