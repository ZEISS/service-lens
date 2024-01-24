import { SubNav, SubNavTitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewSolutionForm } from '@/components/solutions/new-form'
import { PropsWithChildren } from 'react'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page({
  searchParams
}: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <SubNav>
        <SubNavTitle>New Solution</SubNavTitle>
      </SubNav>
      <Section>
        <NewSolutionForm />
      </Section>
    </>
  )
}
