import { MovieCard } from "components/MovieCard"
import { Movie, Show } from "types"

interface MovieListProps {
  title?: string
  movies?: Movie[] | Show[]
  isLoading: boolean
  error: Error | null
  theme?: "center"
}

export const MovieList = ({
  title,
  movies,
  isLoading,
  error,
  theme,
}: MovieListProps) => (
  <div className="px-8 md:px-16 py-6">
    {title && (
      <h2 className="text-2xl font-medium mb-10 text-gray-200">{title}</h2>
    )}
    <div
      className={`flex flex-wrap gap-8 justify-center ${
        theme === "center" ? "" : "md:justify-start"
      }`}
    >
      {isLoading && <div className="text-gray-200">Loading...</div>}
      {error && <div className="text-red-600">Error: {error.message}</div>}
      {movies &&
        movies.map((movie) => (
          <MovieCard
            movie={movie}
            key={movie.id}
            randomMovie={movies.length === 1}
          />
        ))}
    </div>
  </div>
)
