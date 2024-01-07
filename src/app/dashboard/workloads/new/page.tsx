import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewProfilesForm } from './components/new-form'
import { Suspense } from 'react'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          New Workload
          <SubNavSubtitle>
            Workload describes an application or service that serve a business
            process.
          </SubNavSubtitle>
        </SubNavTitle>
      </SubNav>
      <Section>
        <Suspense>
          <NewProfilesForm />
        </Suspense>
      </Section>
    </>
  )
}
