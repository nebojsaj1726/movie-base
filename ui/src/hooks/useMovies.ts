import {
  useQuery,
  UseQueryResult,
  UseQueryOptions,
  keepPreviousData,
} from "@tanstack/react-query"
import { gql, GraphQLClient } from "graphql-request"
import { HomePageData, Movie } from "types"

const endpoint = import.meta.env.VITE_GRAPHQL_ENDPOINT
const client = new GraphQLClient(endpoint)

export const useRandomMoviesQuery = ({
  count,
  genre,
  year,
  rating,
  enabled,
}: {
  count?: number
  genre?: string[]
  year?: number
  rating?: number
  enabled?: boolean
}): UseQueryResult<Movie[], Error> => {
  const options: UseQueryOptions<Movie[], Error> = {
    queryKey: ["randomMovies", count, genre, year, rating],
    queryFn: async () => {
      const query = gql`
        query GetRandomMovies(
          $count: Int
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
            actors
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
    enabled,
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
            actors
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
    placeholderData: keepPreviousData,
  }

  return useQuery(options)
}

export const useGetMoviesQuery = ({
  limit,
  offset,
  genre,
  year,
  rating,
}: {
  limit?: number
  offset?: number
  genre?: string[]
  year?: number
  rating?: number
}): UseQueryResult<{ movies: Movie[]; totalCount: number }, Error> => {
  const options: UseQueryOptions<
    { movies: Movie[]; totalCount: number },
    Error
  > = {
    queryKey: ["getMovies", limit, offset, genre, year, rating],
    queryFn: async () => {
      const query = gql`
        query GetMovies(
          $limit: Int
          $offset: Int
          $genre: [String!]
          $year: Int
          $rating: Float
        ) {
          getMovies(
            limit: $limit
            offset: $offset
            genre: $genre
            year: $year
            rating: $rating
          ) {
            movies {
              id
              title
              rate
              year
              description
              genres
              duration
              imageURL
              actors
            }
            totalCount
          }
        }
      `

      const data = await client.request<{
        getMovies: { movies: Movie[]; totalCount: number }
      }>(query, {
        limit,
        offset,
        genre,
        year,
        rating,
      })
      return data.getMovies
    },
  }

  return useQuery(options)
}

export const useGetHomePageMovies = (): UseQueryResult<HomePageData, Error> => {
  const options: UseQueryOptions<HomePageData, Error> = {
    queryKey: ["getHomePageMovies"],
    queryFn: async () => {
      const query = gql`
        query GetHomePageMovies {
          getHomePageData {
            latestMovies {
              id
              title
              rate
              year
              description
              genres
              duration
              imageURL
              actors
            }
            featuredMovies {
              id
              title
              rate
              year
              description
              genres
              duration
              imageURL
              actors
            }
            movieOfTheDay {
              id
              title
              rate
              year
              description
              genres
              duration
              imageURL
              actors
            }
          }
        }
      `

      const data = await client.request<{ getHomePageData: HomePageData }>(
        query
      )
      return data.getHomePageData
    },
  }

  return useQuery(options)
}
