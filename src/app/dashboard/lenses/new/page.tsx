import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewSolutionForm } from '@/components/lenses/new-form'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          New Lens
          <SubNavSubtitle>
            Measure workloads against best practices.
          </SubNavSubtitle>
        </SubNavTitle>
      </SubNav>
      <Section>
        <NewSolutionForm />
      </Section>
    </>
  )
}
