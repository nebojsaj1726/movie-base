import { Layout } from "components/Layout"
import { Spinner } from "components/Spinner"
import { useShowByIdQuery } from "hooks/useMovies"
import { useEffect } from "react"
import { useParams } from "react-router-dom"

export const Show = () => {
  const { id } = useParams<{ id: string }>()
  const { data, isLoading, error } = useShowByIdQuery(id ?? "")

  useEffect(() => {
    window.scrollTo(0, 0)
  }, [])

  return (
    <Layout>
      <div className="px-8 md:px-20 py-16 bg-deep-blue min-h-screen">
        {isLoading && <Spinner />}
        {error && <div className="text-red-600">Error: {error.message}</div>}
        {data && (
          <div className="flex flex-col gap-10 md:flex-row">
            <div>
              <img
                src={data.imageURL}
                alt={data.title}
                className="max-w-full md:max-w-lg rounded-lg"
              />
            </div>
            <div className="text-gray-300 text-lg pt-6 max-w-2xl">
              <h1 className="text-3xl md:text-4xl font-semibold mb-6 text-gray-200">
                {data.title}
              </h1>
              <div className="flex items-center gap-2 mb-1">
                <p className="">{data.year}</p>
                <p>
                  <span className="text-yellow-300">&#9733;</span>
                  {data.rate}
                </p>
              </div>
              <p className="mb-6">{data.description}</p>
              <p className="mb-2">Genres: {data.genres}</p>
              <p>Cast: {data.actors}</p>
            </div>
          </div>
        )}
      </div>
    </Layout>
  )
}
