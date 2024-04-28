export const yearOptions = () => {
  const currentYear = new Date().getFullYear()
  const years = Array.from(
    { length: currentYear - 1913 },
    (_, index) => currentYear - index
  )
  return years.map((year) => ({ label: String(year), value: year }))
}

export const ratingOptions = () => {
  const ratings = Array.from({ length: 9 }, (_, index) => index + 1)
  return ratings.map((rating, index) => ({
    label: `> ${index + 1}`,
    value: rating,
  }))
}

export const genreOptions = () => {
  const genres = [
    "Action",
    "Adventure",
    "Animation",
    "Comedy",
    "Crime",
    "Drama",
    "Documentary",
    "Science Fiction",
    "History",
    "Horror",
    "Fantasy",
    "Music",
    "Romance",
    "Thriller",
    "War",
    "Western",
  ]
  return genres.map((genre) => ({
    label: genre,
    value: genre,
  }))
}
