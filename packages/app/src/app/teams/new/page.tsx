import { SubNav, SubNavTitle, SubNavSubtitle } from '@/components/sub-nav'
import { Section } from '@/components/section'
import { NewTeamForm } from '@/components/teams/new-form'
import { Suspense } from 'react'

export default function Page() {
    return (
        <>
            <SubNav>
                <SubNavTitle>
                    New Team
                    <SubNavSubtitle>Create a new team.</SubNavSubtitle>
                </SubNavTitle>
            </SubNav>
            <Section>
                <Suspense>
                    <NewTeamForm />
                </Suspense>
            </Section>
        </>
    )
}
