import {
  SubNav,
  SubNavTitle,
  SubNavSubtitle,
  SubNavActions
} from '@/components/sub-nav'
import { Section } from '@/components/section'
import { api } from '@/trpc/server-http'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import DateFormat from '@/components/date-format'
import { ActionsDropdown } from './components/actions-dropdown'
import { EditProfileForm } from '@/components/dashboard/profiles/edit-form'

export type PageProps = {
  params: { id: string }
}

export default async function Page({ params }: PageProps) {
  const profile = await api.getProfile.query(params?.id)
  const questions = await api.listProfilesQuestions.query()

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {profile?.name}
          <SubNavSubtitle>{profile?.description}</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <ActionsDropdown profile={profile} />
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
            <div className="grid gap-8">
              <Card>
                <CardHeader className="space-y-1">
                  <CardTitle className="text-2xl">Overview</CardTitle>
                </CardHeader>
                <CardContent className="grid gap-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                        Last updated
                      </h2>
                      <DateFormat date={profile?.dataValues?.updatedAt} />
                    </div>
                  </div>
                  <Separator />
                  <p>{profile?.description || 'No description provided.'}</p>
                </CardContent>
              </Card>
              {profile && (
                <EditProfileForm profile={profile} questions={questions} />
              )}
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
