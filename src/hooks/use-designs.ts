import { graphql } from "@/gql/gql"
import { useQuery } from "@tanstack/react-query"
import { request } from "graphql-request"

const API_URL = "http://localhost:3000/api/graphql"

const designsQueryDocument = graphql(`
    query Query {
        designs {
            id
            title
            body
        }
    }
`)

export const fetchDesigns = async () => {
  return request(API_URL, designsQueryDocument)
}

export function useDesigns() {
  return useQuery({
    queryKey: ["designs"],
    queryFn: fetchDesigns,
  })
}

const Designs = () => {
  const { data } = useDesigns()
  return data
}
