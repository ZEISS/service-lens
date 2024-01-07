import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewSolutionForm } from './components/new-form'
import { api } from '@/trpc/server-invoker'
import { SolutionTemplate } from '@/db/models/solution-templates'

export type PageProps = {
  searchParams: { template: string }
}

export default async function Page({ searchParams }: PageProps) {
  const template =
    searchParams.template === '_blank'
      ? new SolutionTemplate()
      : await api.getSolutionTemplate.query(searchParams.template)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Solutions
          <SubNavSubtitle>Design, discuss, review, and build.</SubNavSubtitle>
        </SubNavTitle>
      </SubNav>
      <Section>
        <NewSolutionForm template={template ?? new SolutionTemplate()} />
      </Section>
    </>
  )
}
