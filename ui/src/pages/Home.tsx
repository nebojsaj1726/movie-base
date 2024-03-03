import { Link } from "react-router-dom"
import logo from "assets/logo.svg"
import { SearchBar } from "components/SearchBar"

export const Home = () => {
  return (
    <div>
      <header className="bg-dark-midnight text-gray-200 px-6 md:px-12 py-6 flex flex-col md:flex-row gap-6 justify-between md:text-lg">
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
        <SearchBar />
      </header>
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mt-8">Welcome to the Home Page</h1>
        {/* Add your home page content here */}
      </div>
    </div>
  )
}
