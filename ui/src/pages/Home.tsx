import { useRandomMoviesQuery } from "hooks/useMovies"
import { Link } from "react-router-dom"
import logo from "assets/logo.svg"

export const Home = () => {
  const { data, isLoading, error } = useRandomMoviesQuery(10, ["Drama"])

  console.log(data, isLoading, error)

  return (
    <div>
      <header className="bg-dark-midnight text-gray-200 px-12 py-6 flex justify-between">
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
        <div className="ml-6 flex-grow max-w-xl">
          <input
            type="text"
            placeholder="Search"
            className="bg-slate-light w-full text-gray-200 px-4 py-2 rounded focus:outline-none focus:bg-gray-700"
          />
        </div>
      </header>
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mt-8">Welcome to the Home Page</h1>
        {/* Add your home page content here */}
      </div>
    </div>
  )
}
