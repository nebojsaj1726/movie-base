import { useEffect, useRef, useState } from "react"
import { Link } from "react-router-dom"
import { Movie } from "types"
import posterImage from "assets/poster.jpeg"

interface MovieCardProps {
  movie: Movie
}

export const MovieCard = ({ movie }: MovieCardProps) => {
  const cardRef = useRef<HTMLDivElement>(null)
  const [isOnRight, setIsOnRight] = useState(false)
  const [showDetails, setShowDetails] = useState(false)
  let hoverTimer: NodeJS.Timeout

  const handleMouseEnter = () => {
    hoverTimer = setTimeout(() => {
      setShowDetails(true)
    }, 500)
  }

  const handleMouseLeave = () => {
    clearTimeout(hoverTimer)
    setShowDetails(false)
  }

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
    <div
      className="group relative w-full md:w-17/100"
      ref={cardRef}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      <div className="bg-slate-light rounded-lg shadow-md text-gray-200">
        <Link to={`/movies/${movie.id}`}>
          <img
            src={movie.imageURL || posterImage}
            alt={movie.title}
            className="rounded-t-lg object-contain w-full h-full"
          />
          <div className="px-2 py-1">
            <div className="flex gap-3 text-gray-300">
              <p>{movie.year}</p>
              <p>
                <span className="mr-1 text-yellow-300">&#9733;</span>
                {movie.rate}
              </p>
            </div>
            <h3 className="font-semibold overflow-hidden whitespace-nowrap overflow-ellipsis">
              {movie.title}
            </h3>
          </div>
        </Link>
      </div>
      <div
        className={`absolute top-0 flex flex-col bg-slate-light border-4 border-slate-500 text-gray-200 
        rounded-md mx-2 p-4 z-10 w-72 
        invisible h-0 opacity-0 ${
          showDetails &&
          "lg:group-hover:visible lg:group-hover:h-full lg:group-hover:opacity-100"
        }
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
