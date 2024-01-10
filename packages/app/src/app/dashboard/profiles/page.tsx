import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { AddProfileButton } from '@/components/dashboard/profiles/add-button'
import { ProfilesDataTable } from '@/components/dashboard/profiles/data-table'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          <p>Profiles</p>
          <SubNavSubtitle>Business context of a workload.</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddProfileButton />
        </SubNavActions>
      </SubNav>
      <Main>
        <ProfilesDataTable />
      </Main>
    </>
  )
}
