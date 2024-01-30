import { SubNav, SubNavTitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewLensForm } from '@/components/lenses/new-form'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
}

export const revalidate = 0 // no cache

export default function Page(props: NextPageProps) {
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
