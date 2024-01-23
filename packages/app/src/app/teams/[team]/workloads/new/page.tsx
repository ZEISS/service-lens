import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewWorkloadForm } from '@/components/workloads/new-form'
import { Suspense, type PropsWithChildren } from 'react'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default function Page({ params }: PropsWithChildren<NextPageProps>) {
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
          <NewWorkloadForm teamSlug={params.team} />
        </Suspense>
      </Section>
    </>
  )
}
