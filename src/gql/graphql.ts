/* eslint-disable */
import type { GraphQLResolveInfo } from "graphql"
import type { GraphQLCodegenTypescriptTypeDefsDocumentNode } from "graphql-codegen-typescript-typedefs"
import type { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core"
export type Maybe<T> = T | null
export type InputMaybe<T> = T | null | undefined
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] }
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> }
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never }
export type Incremental<T> = T | { [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string }
  String: { input: string; output: string }
  Boolean: { input: boolean; output: boolean }
  Int: { input: number; output: number }
  Float: { input: number; output: number }
}

export type Design = {
  __typename?: "Design"
  body: Scalars["String"]["output"]
  id: Scalars["Int"]["output"]
  title: Scalars["String"]["output"]
}

export type Query = {
  __typename?: "Query"
  designs?: Maybe<Array<Maybe<Design>>>
}

export type QueryQueryVariables = Exact<{ [key: string]: never }>

export type QueryQuery = {
  __typename?: "Query"
  designs?: Array<{ __typename?: "Design"; id: number; title: string; body: string } | null> | null
}

export const QueryDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "Query" },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "designs" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "id" } },
                { kind: "Field", name: { kind: "Name", value: "title" } },
                { kind: "Field", name: { kind: "Name", value: "body" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<QueryQuery, QueryQueryVariables>

export type ResolverTypeWrapper<T> = Promise<T> | T

export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>
}
export type Resolver<
  TResult,
  TParent = Record<PropertyKey, never>,
  TContext = Record<PropertyKey, never>,
  TArgs = Record<PropertyKey, never>,
> = ResolverFn<TResult, TParent, TContext, TArgs> | ResolverWithResolve<TResult, TParent, TContext, TArgs>

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => Promise<TResult> | TResult

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => TResult | Promise<TResult>

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>

export type SubscriptionResolver<
  TResult,
  TKey extends string,
  TParent = Record<PropertyKey, never>,
  TContext = Record<PropertyKey, never>,
  TArgs = Record<PropertyKey, never>,
> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>

export type TypeResolveFn<TTypes, TParent = Record<PropertyKey, never>, TContext = Record<PropertyKey, never>> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo,
) => Maybe<TTypes> | Promise<Maybe<TTypes>>

export type IsTypeOfResolverFn<T = Record<PropertyKey, never>, TContext = Record<PropertyKey, never>> = (
  obj: T,
  context: TContext,
  info: GraphQLResolveInfo,
) => boolean | Promise<boolean>

export type NextResolverFn<T> = () => Promise<T>

export type DirectiveResolverFn<
  TResult = Record<PropertyKey, never>,
  TParent = Record<PropertyKey, never>,
  TContext = Record<PropertyKey, never>,
  TArgs = Record<PropertyKey, never>,
> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo,
) => TResult | Promise<TResult>

/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  Design: ResolverTypeWrapper<Design>
  String: ResolverTypeWrapper<Scalars["String"]["output"]>
  Int: ResolverTypeWrapper<Scalars["Int"]["output"]>
  Query: ResolverTypeWrapper<Record<PropertyKey, never>>
  Boolean: ResolverTypeWrapper<Scalars["Boolean"]["output"]>
}

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  Design: Design
  String: Scalars["String"]["output"]
  Int: Scalars["Int"]["output"]
  Query: Record<PropertyKey, never>
  Boolean: Scalars["Boolean"]["output"]
}

export type DesignResolvers<
  ContextType = any,
  ParentType extends ResolversParentTypes["Design"] = ResolversParentTypes["Design"],
> = {
  body?: Resolver<ResolversTypes["String"], ParentType, ContextType>
  id?: Resolver<ResolversTypes["Int"], ParentType, ContextType>
  title?: Resolver<ResolversTypes["String"], ParentType, ContextType>
}

export type QueryResolvers<
  ContextType = any,
  ParentType extends ResolversParentTypes["Query"] = ResolversParentTypes["Query"],
> = {
  designs?: Resolver<Maybe<Array<Maybe<ResolversTypes["Design"]>>>, ParentType, ContextType>
}

export type Resolvers<ContextType = any> = {
  Design?: DesignResolvers<ContextType>
  Query?: QueryResolvers<ContextType>
}

export const Schema: string =
  "schema{query:Query}type Design{body:String!id:Int!title:String!}type Query{designs:[Design]}"
export const TypeDefs = {
  kind: "Document",
  definitions: [
    {
      kind: "SchemaDefinition",
      directives: [],
      operationTypes: [
        {
          kind: "OperationTypeDefinition",
          operation: "query",
          type: { kind: "NamedType", name: { kind: "Name", value: "Query" } },
        },
      ],
    },
    {
      kind: "ObjectTypeDefinition",
      name: { kind: "Name", value: "Design" },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "body" },
          arguments: [],
          type: { kind: "NonNullType", type: { kind: "NamedType", name: { kind: "Name", value: "String" } } },
          directives: [],
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "id" },
          arguments: [],
          type: { kind: "NonNullType", type: { kind: "NamedType", name: { kind: "Name", value: "Int" } } },
          directives: [],
        },
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "title" },
          arguments: [],
          type: { kind: "NonNullType", type: { kind: "NamedType", name: { kind: "Name", value: "String" } } },
          directives: [],
        },
      ],
    },
    {
      kind: "ObjectTypeDefinition",
      name: { kind: "Name", value: "Query" },
      interfaces: [],
      directives: [],
      fields: [
        {
          kind: "FieldDefinition",
          name: { kind: "Name", value: "designs" },
          arguments: [],
          type: { kind: "ListType", type: { kind: "NamedType", name: { kind: "Name", value: "Design" } } },
          directives: [],
        },
      ],
    },
  ],
  loc: { start: 0, end: 92 },
} as unknown as GraphQLCodegenTypescriptTypeDefsDocumentNode
