import { AddLensButton } from '@/components/lenses/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { LensesDataTable } from '@/components/lenses/data-table'

export default function Lenses() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Lenses
          <SubNavSubtitle>
            Measure any architecture against best practices.
          </SubNavSubtitle>
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
