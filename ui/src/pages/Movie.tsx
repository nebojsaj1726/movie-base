import { useParams } from "react-router-dom"

export const Movie = () => {
  const { id } = useParams<{ id: string }>()

  return (
    <div>
      <h1>Movie Details for ID: {id}</h1>
      {/* Add your movie details content here */}
    </div>
  )
}
