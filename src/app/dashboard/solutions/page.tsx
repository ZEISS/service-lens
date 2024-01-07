import { AddSolution } from './components/add-solution'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import SolutionsDataTable from './components/data-table'

export default function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Solutions
          <SubNavSubtitle>
            Design, review, and execute solutions.
          </SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddSolution />
        </SubNavActions>
      </SubNav>
      <Main>
        <SolutionsDataTable />
      </Main>
    </>
  )
}
