import { AddLensButton } from './components/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { LensesDataTable } from './components/data-table'

export default function Lenses() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Lenses
          <SubNavSubtitle>Review specifications for workloads</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddLensButton />
        </SubNavActions>
      </SubNav>
      <Main>
        <LensesDataTable />
      </Main>
    </>
  )
}
