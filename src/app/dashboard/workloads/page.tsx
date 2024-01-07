import { AddWorkloadButton } from './components/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { WorkloadDataTable } from './components/data-table'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Workloads
          <SubNavSubtitle>Manage and review workflows</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddWorkloadButton />
        </SubNavActions>
      </SubNav>
      <Main>
        <WorkloadDataTable />
      </Main>
    </>
  )
}
