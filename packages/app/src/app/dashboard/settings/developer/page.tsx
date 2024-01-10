import { Separator } from '@/components/ui/separator'
import { DeveloperForm } from './developer-form'

export default function Page() {
  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Developer</h3>
        <p className="text-sm text-muted-foreground">
          Danger zone. Only for developers.
        </p>
      </div>

      <Separator />
      <DeveloperForm />
    </div>
  )
}
