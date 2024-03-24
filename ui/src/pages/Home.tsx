import { Link } from "react-router-dom"
import logo from "assets/logo.svg"
import { SearchBar } from "components/SearchBar"
import { useGetHomePageMovies } from "hooks/useMovies"
import { MovieList } from "components/MovieList"

export const Home = () => {
  const { data, isLoading, error } = useGetHomePageMovies()

  return (
    <div>
      <header className="bg-dark-midnight text-gray-200 px-6 md:px-12 py-6 flex flex-col md:flex-row gap-6 justify-between md:text-lg">
        <div className="flex items-center space-x-3">
          <div>
            <Link to="/">
              <img src={logo} alt="Logo" className="h-8 mr-2 rounded-md" />
            </Link>
          </div>
          <nav>
            <ul className="flex space-x-3 font-medium">
              <li>
                <Link to="/movies" className="hover:text-white">
                  Movies
                </Link>
              </li>
              <li>
                <Link to="/shows" className="hover:text-white">
                  Shows
                </Link>
              </li>
            </ul>
          </nav>
        </div>
        <SearchBar />
      </header>
      <div className="bg-deep-blue pb-8">
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
      </div>
    </div>
  )
}
