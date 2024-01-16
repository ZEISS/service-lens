import { ProfileForm } from './components/profile-form'
import { api } from '@/trpc/server-invoker'
import { Separator } from '@/components/ui/separator'

export default async function Page() {
  const me = await api.me.query()

  return (
    <>
      <div>
        <h3 className="text-lg font-medium">General</h3>
        <p className="text-sm text-muted-foreground">
          This is how others will see you on the site.
        </p>
      </div>
      <Separator />
      {me && <ProfileForm session={me} />}
    </>
  )
}
