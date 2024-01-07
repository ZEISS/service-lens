import {
  SubNav,
  SubNavTitle,
  SubNavSubtitle,
  SubNavActions
} from '@/components/sub-nav'
import { Section } from '@/components/section'
import { api } from '@/trpc/server-http'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/table'
import { Separator } from '@/components/ui/separator'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import DateFormat from '@/components/date-format'
import { ActionsDropdown } from './components/actions-dropdown'
import { PublishButton } from './components/publish-button'

export type PageProps = {
  params: { id: string }
}

export default async function Page({ params }: PageProps) {
  const lens = await api.getLens.query(params?.id)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {lens?.name}
          <SubNavSubtitle>{lens?.description}</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          {lens?.isDraft && <PublishButton lensId={lens.id} />}
          <ActionsDropdown lens={lens} />
        </SubNavActions>
      </SubNav>
      <Section>
        <Tabs defaultValue="overview" className="h-full space-y-6">
          <div className="space-between flex items-center">
            <TabsList>
              <TabsTrigger value="overview" className="relative">
                Overview
              </TabsTrigger>
              <TabsTrigger value="properties">Properties</TabsTrigger>
              <TabsTrigger value="permissions" disabled>
                Permissions
              </TabsTrigger>
            </TabsList>
          </div>
          <TabsContent
            value="overview"
            className="border-none p-0 outline-none"
          >
            <div className="grid gap-4">
              <Card>
                <CardHeader className="space-y-1">
                  <CardTitle className="text-2xl">Overview</CardTitle>
                </CardHeader>
                <CardContent className="grid gap-4">
                  <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                    <div className="space-y-1">
                      <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                        Version
                      </h2>
                      <p>{lens?.version}</p>
                    </div>
                    <div className="space-y-1">
                      <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                        Number of Pillars
                      </h2>
                      <p>{lens?.pillars?.length}</p>
                    </div>
                    <div className="space-y-1">
                      <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                        Status
                      </h2>
                      <p>{lens?.isDraft ? 'DRAFT' : 'PUBLISHED'}</p>
                    </div>
                  </div>
                </CardContent>
                <CardContent className="grid gap-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                        Last updated
                      </h2>
                      <DateFormat date={lens?.dataValues?.updatedAt} />
                    </div>
                  </div>
                  <Separator />
                  <p>{lens?.description || 'No description provided.'}</p>
                </CardContent>
              </Card>
              <Card>
                <CardHeader className="space-y-1">
                  <CardTitle className="text-2xl">Pillars</CardTitle>
                </CardHeader>
                <CardContent>
                  <Table>
                    <TableHeader>
                      <TableRow>
                        <TableHead>Name</TableHead>
                        <TableHead>Questions</TableHead>
                        <TableHead>High Risk</TableHead>
                        <TableHead>Medium Risk</TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {lens?.pillars?.map(pillar => (
                        <TableRow key={pillar.id}>
                          <TableCell className="font-medium">
                            {pillar.name}
                          </TableCell>
                          <TableCell>{pillar.questions?.length}</TableCell>
                          <TableCell></TableCell>
                          <TableCell className="text-right"></TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </CardContent>
              </Card>
            </div>
          </TabsContent>
          <TabsContent
            value="properties"
            className="h-full flex-col border-none p-0 data-[state=active]:flex"
          >
            Nothing yet.
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
