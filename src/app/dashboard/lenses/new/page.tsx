import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewSolutionForm } from './components/new-form'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          New Lens
          <SubNavSubtitle>Lenses help to evalute workloads.</SubNavSubtitle>
        </SubNavTitle>
      </SubNav>
      <Section>
        <NewSolutionForm />
      </Section>
    </>
  )
}
