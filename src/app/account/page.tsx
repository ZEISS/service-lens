import { ProfileForm } from './components/profile-form'
import { api } from '@/trpc/server-invoker'

export default async function Page() {
  const me = await api.me.query()

  return <>{me && <ProfileForm session={me} />}</>
}
