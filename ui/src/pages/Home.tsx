import { useGetHomePageMovies } from "hooks/useMovies"
import { MovieList } from "components/MovieList"
import { Layout } from "components/Layout"

export const Home = () => {
  const { data, isLoading, error } = useGetHomePageMovies()

  return (
    <Layout>
      <div className="bg-deep-blue pb-8">
        <h1 className="text-gray-200 px-8 md:px-16 pt-12 text-3xl text-center">
          <span className="text-blue-500">Discover</span> and{" "}
          <span className="text-yellow-500">explore</span>. Welcome to our movie
          hub!
        </h1>
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
        <p className="text-gray-200 mt-6 px-8 md:px-16">* Ratings are IMDb.</p>
      </div>
    </Layout>
  )
}
