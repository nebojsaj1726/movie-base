import { useEffect, useRef, useState } from "react"
import { Link } from "react-router-dom"
import { Movie } from "types"

interface MovieCardProps {
  movie: Movie
}

export const MovieCard = ({ movie }: MovieCardProps) => {
  const cardRef = useRef<HTMLDivElement>(null)
  const [isOnRight, setIsOnRight] = useState(false)

  useEffect(() => {
    const handleResize = () => {
      if (!cardRef.current) return
      const cardRect = cardRef.current.getBoundingClientRect()
      const viewportWidth = document.documentElement.clientWidth
      const distanceToRight = viewportWidth - cardRect.right

      setIsOnRight(distanceToRight < 300)
    }

    handleResize()
    window.addEventListener("resize", handleResize)

    return () => {
      window.removeEventListener("resize", handleResize)
    }
  }, [])

  return (
    <div className="group relative" ref={cardRef}>
      <div className="bg-slate-light rounded-lg shadow-md w-fit sm:w-52 2xl:w-60 text-gray-200">
        <Link to={`/movies/${movie.id}`}>
          <img
            src={movie.imageURL}
            alt={movie.title}
            className="rounded-t-lg mb-2 w-11/12 mx-auto sm:w-full"
          />
          <div className="px-3">
            <div className="flex gap-3 text-gray-300">
              <p>{movie.year}</p>
              <p>
                <span className="mr-1 text-yellow-300">&#9733;</span>
                {movie.rate}
              </p>
            </div>
            <h3 className="font-semibold mb-2 overflow-hidden whitespace-nowrap overflow-ellipsis">
              {movie.title}
            </h3>
          </div>
        </Link>
      </div>
      <div
        className={`absolute top-0 flex flex-col bg-slate-light border-4 border-slate-500 text-gray-200 rounded-md mx-2 p-4 z-10 w-72 invisible
        lg:group-hover:visible h-0 lg:group-hover:h-full opacity-0 lg:group-hover:opacity-100 
        transition-opacity duration-300 ${
          isOnRight ? "right-full" : "left-full"
        }`}
      >
        <h3 className="text-lg font-semibold mb-1">{movie.title}</h3>
        <div className="flex gap-3 text-gray-300 text-sm mb-3">
          <p>{movie.year}</p>
          <p>
            <span className="mr-1 text-yellow-300">&#9733;</span>
            {movie.rate}
          </p>
        </div>
        <p className="text-sm mb-2 flex-1 overflow-hidden overflow-ellipsis">
          {movie.description}
        </p>
        <p className="text-sm">{movie.genres}</p>
        <p className="text-sm">{movie.duration}</p>
        <p className="text-sm overflow-hidden max-h-10">{`Cast: ${movie.actors}`}</p>
      </div>
    </div>
  )
}
