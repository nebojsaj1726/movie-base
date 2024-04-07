import { useGetHomePageMovies } from "hooks/useMovies"
import { MovieList } from "components/MovieList"
import { Layout } from "components/Layout"

export const Home = () => {
  const { data, isLoading, error } = useGetHomePageMovies()

  return (
    <Layout>
      <div className="bg-deep-blue pb-8">
        <p className="text-gray-200 px-8 md:px-16 pt-10 text-xl">
          Discover and explore. Welcome to our movie hub!
        </p>
        <MovieList
          title="Latest movies"
          movies={data?.latestMovies}
          isLoading={isLoading}
          error={error}
        />
        <MovieList
          title="Featured movies"
          movies={data?.featuredMovies}
          isLoading={isLoading}
          error={error}
        />
        <MovieList
          title="Movie of the day"
          movies={data ? [data.movieOfTheDay] : []}
          isLoading={isLoading}
          error={error}
        />
        <p className="text-gray-200 px-8 md:px-16">* Ratings are IMDb.</p>
      </div>
    </Layout>
  )
}
