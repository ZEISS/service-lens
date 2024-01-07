import { Separator } from '@/components/ui/separator'
import { GeneralForm } from './components/general-form'

export default function Page() {
  return (
    <>
      <div>
        <h3 className="text-lg font-medium">General</h3>
        <p className="text-sm text-muted-foreground">
          This is how others will see you on the site.
        </p>
      </div>
      <Separator />
      <GeneralForm />
    </>
  )
}
