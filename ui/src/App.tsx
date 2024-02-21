import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import { Movies } from "./pages/Movies"
import { Movie } from "./pages/Movie"
import { Home } from "./pages/Home"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"

export const queryClient = new QueryClient()

export const App = () => (
  <QueryClientProvider client={queryClient}>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/movies" element={<Movies />} />
        <Route path="/movies/:id" element={<Movie />} />
      </Routes>
    </Router>
  </QueryClientProvider>
)
