import {
  useQuery,
  UseQueryResult,
  UseQueryOptions,
} from "@tanstack/react-query"
import { gql, GraphQLClient } from "graphql-request"
import { Movie } from "types"

const endpoint = import.meta.env.VITE_GRAPHQL_ENDPOINT
const client = new GraphQLClient(endpoint)

export const useRandomMoviesQuery = (
  count: number = 1,
  genre: string[] = [],
  year?: number,
  rating?: number
): UseQueryResult<Movie[], Error> => {
  const options: UseQueryOptions<Movie[], Error> = {
    queryKey: ["randomMovies", count, genre, year, rating],
    queryFn: async () => {
      const query = gql`
        query GetRandomMovies(
          $count: Int = 1
          $genre: [String!]
          $year: Int
          $rating: Float
        ) {
          getRandomMovies(
            count: $count
            genre: $genre
            year: $year
            rating: $rating
          ) {
            id
            title
            rate
            year
            description
            genres
            duration
            imageURL
          }
        }
      `

      const data = await client.request<{ getRandomMovies: Movie[] }>(query, {
        count,
        genre,
        year,
        rating,
      })
      return data.getRandomMovies
    },
  }

  return useQuery(options)
}

export const useMovieByIdQuery = (id: string): UseQueryResult<Movie, Error> => {
  const options: UseQueryOptions<Movie, Error> = {
    queryKey: ["movieById", id],
    queryFn: async () => {
      const query = gql`
        query GetMovieById($id: String!) {
          getMovieById(id: $id) {
            id
            title
            rate
            year
            description
            genres
            duration
            imageURL
          }
        }
      `
      const data = await client.request<{ getMovieById: Movie }>(query, { id })
      return data.getMovieById
    },
  }
  return useQuery(options)
}

export const useSearchMoviesByKeywordQuery = (
  keyword: string
): UseQueryResult<Movie[], Error> => {
  const options: UseQueryOptions<Movie[], Error> = {
    queryKey: ["searchMoviesByKeyword", keyword],
    queryFn: async () => {
      const query = gql`
        query SearchMoviesByKeyword($keyword: String!) {
          searchMoviesByKeyword(keyword: $keyword) {
            id
            title
            year
            genres
          }
        }
      `

      const data = await client.request<{ searchMoviesByKeyword: Movie[] }>(
        query,
        { keyword }
      )
      return data.searchMoviesByKeyword
    },
  }

  return useQuery(options)
}
