import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { AddProfileButton } from '@/components/dashboard/profiles/add-button'
import { ProfileDataTable } from '@/components/dashboard/profiles/data-table'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          <p>Profiles</p>
          <SubNavSubtitle>
            Provide business context for a workload
          </SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddProfileButton />
        </SubNavActions>
      </SubNav>
      <Main>
        <ProfileDataTable />
      </Main>
    </>
  )
}
