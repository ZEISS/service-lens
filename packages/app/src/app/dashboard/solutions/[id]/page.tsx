import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
  CardDescription
} from '@/components/ui/card'
import DateFormat from '@/components/date-format'
import { Section } from '@/components/section'
import { api } from '@/trpc/server-http'
import { CommentForm } from './components/comment-form'
import { CommentActions } from './components/comment-actions'
import { remark } from 'remark'
import html from 'remark-html'
import Markdown from 'react-markdown'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { ActionsMenu } from './components/actions-menu'

export type PageProps = {
  params: { id: string }
}

export default async function Page({ params }: PageProps) {
  const solution = await api.getSolution.query(params?.id)

  const processedContent = await remark()
    .use(html)
    .process(solution?.body ?? '')

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {solution?.title}
          <SubNavSubtitle>Manage and review workflows</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          {solution && <ActionsMenu solution={solution} />}
        </SubNavActions>
      </SubNav>
      <Section>
        <Card>
          <CardHeader>
            <CardTitle className="flex flex-row items-center justify-between">
              <Button variant="ghost" className="relative h-8 w-8 rounded-full">
                <Avatar className="h-8 w-8">
                  <AvatarImage
                    src={solution?.user?.image}
                    alt={solution?.user?.name}
                  />
                  <AvatarFallback>{}</AvatarFallback>
                </Avatar>
              </Button>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <Markdown
              components={{
                h1(props) {
                  const { node, ...rest } = props
                  return (
                    <h1
                      className="scroll-m-20 text-4xl font-extrabold tracking-tight mt-6 lg:text-5x"
                      {...rest}
                    />
                  )
                },
                h2(props) {
                  const { node, ...rest } = props
                  return (
                    <h1
                      className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight mt-6 first:mt-0"
                      {...rest}
                    />
                  )
                },
                p(props) {
                  const { node, ...rest } = props
                  return (
                    <p
                      className="leading-7 [&:not(:first-child)]:mt-6"
                      {...rest}
                    />
                  )
                }
              }}
            >
              {solution?.body}
            </Markdown>
          </CardContent>
          <CardFooter>
            <CardDescription>
              Updated on <DateFormat date={solution?.dataValues?.updatedAt} />
            </CardDescription>
          </CardFooter>
        </Card>

        {solution?.comments?.map(comment => (
          <Card key={comment.id} className="my-6">
            <CardHeader>
              <CardTitle className="flex flex-row items-center justify-between">
                <Button
                  variant="ghost"
                  className="relative h-8 w-8 rounded-full"
                >
                  <Avatar className="h-8 w-8">
                    <AvatarImage
                      src={comment.user?.image}
                      alt={comment.user?.name}
                    />
                    <AvatarFallback>{}</AvatarFallback>
                  </Avatar>
                </Button>
                <CommentActions comment={comment} />
              </CardTitle>
            </CardHeader>
            <CardContent>{comment.body}</CardContent>
            <CardFooter>
              <CardDescription>
                Commented on
                <DateFormat date={comment?.dataValues?.updatedAt} />
              </CardDescription>
            </CardFooter>
          </Card>
        ))}

        <CommentForm solutionId={params.id} />
      </Section>
    </>
  )
}
