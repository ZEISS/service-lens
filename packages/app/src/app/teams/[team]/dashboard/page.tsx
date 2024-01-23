import { type PropsWithChildren, Suspense } from 'react'
import { SubNav, SubNavTitle } from '@/components/sub-nav'
import { Tabs, TabsTrigger, TabsList, TabsContent } from '@/components/ui/tabs'
import { Main } from '@/components/main'
import { TotalSolutionsCard } from '@/components/dashboard/total-solutions-card'
import { LoadingCard } from '@/components/dashboard/loading-card'
import { TotalWorkloadsCard } from '@/components/dashboard/total-workloads-card'
import { TotalLensesCard } from '@/components/dashboard/total-lenses-card'
import { TotalProfilesCard } from '@/components/dashboard/total-profiles-card'

export const revalidate = 0 // no cache

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page({
  params
}: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <SubNav>
        <SubNavTitle>Home</SubNavTitle>
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
              <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                <Suspense fallback={<LoadingCard />}>
                  <TotalWorkloadsCard teamSlug={params.team} />
                </Suspense>
                <Suspense fallback={<LoadingCard />}>
                  <TotalSolutionsCard teamSlug={params.team} />
                </Suspense>
                <Suspense fallback={<LoadingCard />}>
                  <TotalLensesCard teamSlug={params.team} />
                </Suspense>
                <Suspense fallback={<LoadingCard />}>
                  <TotalProfilesCard teamSlug={params.team} />
                </Suspense>
              </div>
              {/* <div className="grid gap-4">
                <Suspense fallback={<LoadingCard />}>
                  <WorkloadsListCard />
                </Suspense>
              </div>  */}
            </TabsContent>
          </Tabs>
        </div>
      </Main>
    </>
  )
}
