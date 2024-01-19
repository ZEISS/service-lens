import { Metadata } from 'next'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ActionButton } from '@/app/dashboard/components/add-button'
import { Main } from '@/components/main'

export const metadata: Metadata = {
  title: 'Dashboard',
  description: 'Dashboard'
}

export default async function Page() {
  return (
    <>
      <SubNav>
        <SubNavTitle>
          Dashboard
          <SubNavSubtitle>Manage and review workflows</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <ActionButton />
        </SubNavActions>
      </SubNav>
      <Main className="space-y-8 p-8">
        <div className="flex-1 space-y-4">
          <div className="flex items-center justify-between space-y-2"></div>
          <Tabs defaultValue="overview" className="space-y-4">
            <TabsList>
              <TabsTrigger value="overview">Overview</TabsTrigger>
              <TabsTrigger value="analytics" disabled>
                Analytics
              </TabsTrigger>
              <TabsTrigger value="reports" disabled>
                Reports
              </TabsTrigger>
              <TabsTrigger value="notifications" disabled>
                Notifications
              </TabsTrigger>
            </TabsList>
            <TabsContent value="overview" className="space-y-4">
              {/* <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                <Suspense fallback={<LoadingCard />}>
                  <TotalWorkloadsCard />
                </Suspense>
                <Suspense fallback={<LoadingCard />}>
                  <TotalSolutionsCard />
                </Suspense>
              </div>
              <div className="grid gap-4">
                <Suspense fallback={<LoadingCard />}>
                  <WorkloadsListCard />
                </Suspense>
              </div> */}
            </TabsContent>
          </Tabs>
        </div>
      </Main>
    </>
  )
}
