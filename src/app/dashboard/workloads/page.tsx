import { AddWorkloadButton } from '@/components/workloads/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { WorkloadDataTable } from '@/components/workloads/data-table'

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
      <Main className="space-y-8 p-8">
        <WorkloadDataTable />
      </Main>
    </>
  )
}
