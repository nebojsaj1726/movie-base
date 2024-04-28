import { Layout } from "components/Layout"
import { MovieList } from "components/MovieList"
import { Pagination } from "components/Pagination"
import { useGetMoviesQuery, useRandomMoviesQuery } from "hooks/useMovies"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { Filters, FormFilters } from "types"
import { genreOptions, ratingOptions, yearOptions } from "utils/options"
import Select from "react-select"

export const Movies = () => {
  const [currentPage, setCurrentPage] = useState(1)
  const [filters, setFilters] = useState<Filters>({})
  const [showRandom, setShowRandom] = useState(false)
  const { handleSubmit, register, setValue, getValues } = useForm()

  const { data, isLoading, error } = useGetMoviesQuery({
    offset: (currentPage - 1) * 20,
    ...filters,
  })

  const {
    data: randomMoviesData,
    isLoading: randomMoviesLoading,
    error: randomMoviesError,
  } = useRandomMoviesQuery({
    enabled: showRandom,
    ...filters,
  })

  const totalPages = Math.ceil((data?.totalCount || 0) / 20)

  const handlePageChange = (page: number) => {
    setCurrentPage(page)
  }

  const updateFilters = (formData: FormFilters) => {
    const year = formData.year || undefined
    const rating = formData.rating || undefined
    const genre = formData.genre ? [formData.genre] : undefined

    setFilters({ year, rating, genre })
  }

  const onSubmit = (formData: FormFilters) => {
    updateFilters(formData)
    setShowRandom(false)
    setCurrentPage(1)
  }

  const getFormData = () => {
    const formData = {
      year: getValues("year"),
      rating: getValues("rating"),
      genre: getValues("genre"),
    }
    return formData
  }

  const handleRandomMovieClick = () => {
    setShowRandom(true)
    const formData = getFormData()
    updateFilters(formData)
  }

  return (
    <Layout>
      <div className="bg-deep-blue min-h-screen">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="px-8 md:px-24 pt-8 pb-2"
        >
          <div className="flex flex-wrap justify-between gap-y-6 max-w-6xl m-auto">
            <div className="flex flex-wrap items-center gap-4 w-full sm:w-auto">
              <Select
                {...register("year")}
                options={yearOptions()}
                onChange={(selectedOption) => {
                  setValue("year", selectedOption?.value)
                }}
                className="w-full sm:w-48"
                placeholder="Years"
              />
              <Select
                {...register("rating")}
                options={ratingOptions()}
                onChange={(selectedOption) => {
                  setValue("rating", selectedOption?.value)
                }}
                className="w-full sm:w-48"
                placeholder="Ratings"
              />
              <Select
                {...register("genre")}
                options={genreOptions()}
                onChange={(selectedOption) => {
                  setValue("genre", selectedOption?.value)
                }}
                className="w-full sm:w-48"
                placeholder="Genres"
              />
            </div>
            <div className="flex gap-4">
              <button
                type="submit"
                className="bg-blue-500 hover:bg-blue-700 text-gray-100 font-bold py-2 px-4 rounded"
              >
                Filter movies
              </button>
              <button
                type="button"
                onClick={handleRandomMovieClick}
                className="bg-yellow-500 hover:bg-yellow-700 text-gray-100 font-bold py-2 px-4 rounded md:col-span-1"
              >
                Random Movie
              </button>
            </div>
          </div>
        </form>
        <MovieList
          movies={showRandom ? randomMoviesData : data?.movies}
          isLoading={showRandom ? randomMoviesLoading : isLoading}
          error={showRandom ? randomMoviesError : error}
          theme="center"
        />
        {!showRandom && (
          <Pagination
            currentPage={currentPage}
            totalPages={totalPages}
            onPageChange={handlePageChange}
          />
        )}
      </div>
    </Layout>
  )
}
