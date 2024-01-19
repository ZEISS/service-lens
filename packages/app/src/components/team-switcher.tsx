'use client'

import * as React from 'react'
import { CaretSortIcon, CheckIcon } from '@radix-ui/react-icons'
import { cn } from '@/lib/utils'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { NewTeamFormValues } from '@/components/teams/new-form.schema'
import { defaultValues } from '@/components/teams/new-form.schema'
import { TeamsCreateSchema } from '@/server/routers/schemas/teams'
import { useAction } from '@/trpc/client'
import { rhfAction } from '@/components/teams/new-form.action'
import { Textarea } from '@/components/ui/textarea'
import {
  Form,
  FormControl,
  FormItem,
  FormLabel,
  FormDescription,
  FormMessage,
  FormField
} from '@/components/ui/form'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator
} from '@/components/ui/command'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from '@/components/ui/dialog'
import { useForm } from 'react-hook-form'
import { Input } from '@/components/ui/input'
import {
  Popover,
  PopoverContent,
  PopoverTrigger
} from '@/components/ui/popover'
import { api } from '@/trpc/client'
import { useRouter } from 'next/navigation'

type Team = (typeof groups)[number]['teams'][number]

type PopoverTriggerProps = React.ComponentPropsWithoutRef<typeof PopoverTrigger>

interface TeamSwitcherProps extends PopoverTriggerProps { }

export default function TeamSwitcher({ className }: TeamSwitcherProps) {
  const me = React.use(api.me.query())
  const user = React.use(api.users.get.query())
  const teams = React.useMemo(
    () => user?.teams?.map(team => ({ label: team.name, value: team.id })),
    [user?.teams]
  )

  const [open, setOpen] = React.useState(false)
  const [showNewTeamDialog, setShowNewTeamDialog] = React.useState(false)
  const [selectedTeam, setSelectedTeam] = React.useState<Team>(teams[0])
  const router = useRouter()

  const form = useForm<NewTeamFormValues>({
    resolver: zodResolver(TeamsCreateSchema),
    defaultValues,
    mode: 'onChange'
  })

  const mutation = useAction(rhfAction)
  const handleSubmit = async (data: NewTeamFormValues) =>
    await mutation.mutateAsync({ ...data })

  React.useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/teams/${mutation.data.slug}/settings`)
    }
  }, [router, mutation.status, mutation.data])

  return (
    <Dialog open={showNewTeamDialog} onOpenChange={setShowNewTeamDialog}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a team"
            className={cn('w-[200px] justify-between', className)}
          >
            <Avatar className="mr-2 h-5 w-5">
              <AvatarImage
                src={me?.user.image ?? ''}
                alt={me?.user.name ?? ''}
              />
              <AvatarFallback>SC</AvatarFallback>
            </Avatar>
            {selectedTeam?.label}
            <CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command>
            <CommandList>
              <CommandInput placeholder="Search team..." />
              <CommandEmpty>No team found.</CommandEmpty>
              <CommandGroup heading="Personal Account">
                <CommandItem className="text-sm">
                  <Avatar className="mr-2 h-5 w-5">
                    <AvatarImage
                      src={me?.user.image ?? ''}
                      alt={me?.user.name ?? ''}
                      className="grayscale"
                    />
                    <AvatarFallback>SC</AvatarFallback>
                  </Avatar>
                  {me?.user.name}
                </CommandItem>
              </CommandGroup>
              <CommandGroup heading="Teams">
                {teams?.map(team => (
                  <CommandItem
                    key={team.value}
                    className="text-sm"
                    onSelect={() => {
                      setSelectedTeam(team)
                      setOpen(false)
                    }}
                  >
                    {team.label}
                  </CommandItem>
                ))}
              </CommandGroup>
              {/* {groups.map(group => (
                <CommandGroup key={group.label} heading={group.label}>
                  {group.teams.map(team => (
                    <CommandItem
                      key={team.value}
                      onSelect={() => {
                        setSelectedTeam(team)
                        setOpen(false)
                      }}
                      className="text-sm"
                    >
                      <Avatar className="mr-2 h-5 w-5">
                        <AvatarImage
                          src={`https://avatar.vercel.sh/${team.value}.png`}
                          alt={team.label}
                          className="grayscale"
                        />
                        <AvatarFallback>SC</AvatarFallback>
                      </Avatar>
                      {team.label}
                      <CheckIcon
                        className={cn(
                          'ml-auto h-4 w-4',
                          selectedTeam.value === team.value
                            ? 'opacity-100'
                            : 'opacity-0'
                        )}
                      />
                    </CommandItem>
                  ))}
                </CommandGroup>
              ))} */}
            </CommandList>
            <CommandSeparator />
            <CommandList>
              <CommandGroup>
                <DialogTrigger asChild>
                  <CommandItem
                    onSelect={() => {
                      setOpen(false)
                      setShowNewTeamDialog(true)
                    }}
                  >
                    Create Team
                  </CommandItem>
                </DialogTrigger>
                <CommandItem onSelect={() => router.push('/account/teams')}>
                  Manage Teams
                </CommandItem>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create team</DialogTitle>
          <DialogDescription>
            Add a new team to manage products and customers.
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            action={rhfAction}
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">
                    <h1>Name</h1>
                  </FormLabel>
                  <FormControl>
                    <Input placeholder="Name ..." {...field} />
                  </FormControl>
                  <FormDescription>Give it a great name.</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="slug"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">Slug</FormLabel>
                  <FormControl>
                    <Input placeholder="Slug ..." {...field} />
                  </FormControl>
                  <FormDescription>
                    {`This is the short name used for URLs (e.g.
                'solution-architects', 'order-service')`}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="contactEmail"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">
                    <h1>Contact email</h1>
                  </FormLabel>
                  <FormControl>
                    <Input placeholder="team@acme.com" {...field} />
                  </FormControl>
                  <FormDescription>
                    Add a shared inbox for you team (optional).
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <div className="grid w-full">
                  <FormItem>
                    <FormLabel className="sr-only">
                      <h1>Description</h1>
                    </FormLabel>
                    <FormControl>
                      <Textarea
                        {...field}
                        className="w-full"
                        placeholder="Add a description ..."
                      />
                    </FormControl>
                    <FormDescription>A desciption of your team</FormDescription>
                    <FormMessage />
                  </FormItem>
                </div>
              )}
            />

            <DialogFooter>
              <Button
                variant="outline"
                onClick={() => setShowNewTeamDialog(false)}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={
                  form.formState.isSubmitting || !form.formState.isValid
                }
              >
                Continue
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
