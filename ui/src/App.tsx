import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import { Movies } from "./pages/Movies"
import { Movie } from "./pages/Movie"
import { Home } from "./pages/Home"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { Shows } from "./pages/Shows"
import { Show } from "./pages/Show"

export const queryClient = new QueryClient()

export const App = () => (
  <QueryClientProvider client={queryClient}>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/movies" element={<Movies />} />
        <Route path="/movies/:id" element={<Movie />} />
        <Route path="/shows" element={<Shows />} />
        <Route path="/shows/:id" element={<Show />} />
      </Routes>
    </Router>
  </QueryClientProvider>
)
