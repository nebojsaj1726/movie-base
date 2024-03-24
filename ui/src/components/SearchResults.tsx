import { useSearchMoviesByKeywordQuery } from "hooks/useMovies"
import { Link } from "react-router-dom"

interface SearchResultsProps {
  searchInput: string
}

export const SearchResults = ({ searchInput }: SearchResultsProps) => {
  const { data, error } = useSearchMoviesByKeywordQuery(searchInput)

  return (
    <div className="absolute z-10 mt-2 w-full">
      <ul className="bg-gray-700 rounded-md">
        {error && (
          <div className="p-4 text-red-600">Error: {error.message}</div>
        )}
        {data &&
          data.map((movie) => (
            <li key={movie.id} className="px-4 py-4 border-b border-gray-200">
              <Link to={`/movies/${movie.id}`}>
                <p className="font-semibold">
                  {`${movie.title} (${movie.year})`}
                </p>
                <div className="text-sm">{movie.genres}</div>
              </Link>
            </li>
          ))}
      </ul>
    </div>
  )
}
