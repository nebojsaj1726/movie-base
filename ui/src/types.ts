export interface Movie {
  id: string
  title: string
  rate: string
  year: string
  description: string
  genres: string
  duration: string
  imageURL: string
  actors: string
}

export interface HomePageData {
  latestMovies: Movie[]
  featuredMovies: Movie[]
  movieOfTheDay: Movie
}

interface FiltersType<T extends string | string[]> {
  year?: number
  rating?: number
  genre?: T
}

export type Filters = FiltersType<string[]>

export type FormFilters = FiltersType<string>
