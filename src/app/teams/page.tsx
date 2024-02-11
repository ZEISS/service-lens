import { redirect } from 'next/navigation'

export interface NextPageProps<Team = string> {
  params: { team: Team }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page({ params }: NextPageProps) {
  return redirect('/home')
}
