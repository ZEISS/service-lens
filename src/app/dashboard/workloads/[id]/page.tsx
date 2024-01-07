import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { OverviewCard } from './components/overview-card'
import { ProfileCard } from './components/profile-card'
import { Section } from '@/components/section'
import { api } from '@/trpc/server-http'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { PropertiesCard } from './components/properties-card'
import { LensCard } from './components/lens-card'
import { MoreButton } from './components/more-button'
import { LensesCard } from './components/lenses-card'

export type PageProps = {
  params: { id: string }
}

export default async function Page({ params }: PageProps) {
  const workload = await api.getWorkload.query(params?.id)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {workload?.name}
          <SubNavSubtitle>Manage and review workflows</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <MoreButton />
        </SubNavActions>
      </SubNav>
      <Section>
        <Tabs defaultValue="overview" className="h-full space-y-6">
          <TabsList>
            <TabsTrigger value="overview" className="relative">
              Overview
            </TabsTrigger>
            <TabsTrigger value="properties">Properties</TabsTrigger>
            <TabsTrigger value="permissions" disabled>
              Permissions
            </TabsTrigger>
          </TabsList>
          <TabsContent
            value="overview"
            className="border-none p-0 outline-none"
          >
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-7">
              {workload && (
                <OverviewCard workload={workload} className="col-span-4" />
              )}
              {workload?.profile && (
                <ProfileCard
                  profile={workload.profile}
                  className="col-span-3"
                />
              )}
              {workload?.lenses && (
                <LensesCard
                  workloadId={workload.id}
                  lenses={workload?.lenses}
                  className="col-span-7"
                />
              )}
            </div>
          </TabsContent>
          <TabsContent
            value="properties"
            className="h-full flex-col border-none p-0 data-[state=active]:flex"
          >
            {workload && <PropertiesCard workload={workload} />}
          </TabsContent>
          <TabsContent
            value="permissions"
            className="h-full flex-col border-none p-0 data-[state=active]:flex"
          >
            Permissions
          </TabsContent>
        </Tabs>
      </Section>
    </>
  )
}
