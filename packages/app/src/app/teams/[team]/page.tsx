import { api } from '@/trpc/server-http'
import { redirect } from 'next/navigation'

export interface NextPageProps<Team = string> {
  params: { team: Team }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page({ params }: NextPageProps) {
  const team = await api.teams.getByName.query({ slug: params.team })

  if (!team) {
    redirect('/home')
  }

  return redirect(`/teams/${team.slug}/dashboard`)
}
