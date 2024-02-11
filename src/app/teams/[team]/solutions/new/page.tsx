import { SubNav, SubNavTitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewSolutionForm } from '@/components/solutions/new-form'

export default async function Page() {
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
