import { useRandomMoviesQuery } from "hooks/useMovies"

export const Home = () => {
  const { data, isLoading, error } = useRandomMoviesQuery(10, ["Drama"])

  console.log(data, isLoading, error)

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      {/* Add your home page content here */}
    </div>
  )
}
