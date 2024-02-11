import { Form } from '@/components/ui/form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'

export default async function Setup() {
  return (
    <>
      <form>
        <div className="container relative h-[800px] flex-col items-center justify-center md:grid lg:max-w-none lg:px-0">
          <div className="lg:p-8">
            <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
              <div className="flex flex-col space-y-2 text-center">
                <h1 className="text-2xl font-semibold tracking-tight">
                  Setup your account
                </h1>
                <p className="text-sm text-muted-foreground">
                  Creating your team.
                </p>
              </div>
            </div>
          </div>
        </div>
      </form>
    </>
  )
}
