import { Layout } from "components/Layout"
import { MovieList } from "components/MovieList"
import { Pagination } from "components/Pagination"
import { useGetMoviesQuery } from "hooks/useMovies"
import { useState } from "react"

export const Movies = () => {
  const [currentPage, setCurrentPage] = useState(1)
  const { data, isLoading, error } = useGetMoviesQuery({
    offset: (currentPage - 1) * 20,
  })
  const totalPages = Math.ceil((data?.totalCount || 0) / 20)

  const handlePageChange = (page: number) => {
    setCurrentPage(page)
  }

  return (
    <Layout>
      <div className="bg-deep-blue min-h-screen">
        <MovieList
          movies={data?.movies}
          isLoading={isLoading}
          error={error}
          theme="center"
        />
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={handlePageChange}
        />
      </div>
    </Layout>
  )
}
