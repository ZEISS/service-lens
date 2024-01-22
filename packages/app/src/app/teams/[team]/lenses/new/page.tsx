import { SubNav, SubNavTitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewLensForm } from '@/components/lenses/new-form'
import { type PropsWithChildren } from 'react'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export const revalidate = 0 // no cache

export default function Page(props: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <SubNav>
        <SubNavTitle>New Lens</SubNavTitle>
      </SubNav>
      <Section>
        <NewLensForm teamSlug={props.params.team} />
      </Section>
    </>
  )
}
