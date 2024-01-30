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
import { columns } from '@/components/workloads/lenses-data-columns'
import { MoreButton } from './components/more-button'
import { DataTable } from '@/components/data-table'
import type { DataTableOptions } from '@/components/data-table'
import { ListWorkloadLensSchema } from '@/server/routers/schemas/workload'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export interface NextPageProps<TeamSlug = string, WorkloadId = string> {
  params: { team: TeamSlug; id: WorkloadId }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export const revalidate = 0 // no cache

export default async function Page(props: NextPageProps) {
  const searchParams = ListWorkloadLensSchema.parse({
    ...props.searchParams,
    ...props.params
  })
  const { rows, count } = await api.workloads.listLens.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)
  const workload = await api.getWorkload.query(props.params.id)

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
            </div>
            <div className="grid gap-4 py-4">
              <DataTable
                columns={columns}
                data={rows}
                pageCount={pageCount}
                pageSize={searchParams.limit}
                pageIndex={searchParams.offset}
                options={options}
              />
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
