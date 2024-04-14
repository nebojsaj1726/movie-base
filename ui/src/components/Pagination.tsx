import { useEffect } from "react"

interface PaginationProps {
  currentPage: number
  totalPages: number
  onPageChange: (page: number) => void
}

export const Pagination = ({
  currentPage,
  totalPages,
  onPageChange,
}: PaginationProps) => {
  const pagesToShow = window.innerWidth < 640 ? 4 : 8

  let startPage = Math.max(1, currentPage - Math.floor(pagesToShow / 2))
  let endPage = Math.min(totalPages, startPage + pagesToShow - 1)

  if (endPage === totalPages && endPage - startPage + 1 < pagesToShow) {
    startPage = Math.max(1, endPage - pagesToShow + 1)
  }

  if (startPage === 1 && endPage - startPage + 1 < pagesToShow) {
    endPage = Math.min(totalPages, startPage + pagesToShow - 1)
  }

  const pages = [...Array(endPage - startPage + 1)].map((_, i) => startPage + i)

  useEffect(() => {
    window.scrollTo({ top: 0, behavior: "smooth" })
  }, [currentPage])

  return (
    <div className="flex justify-center py-8">
      <nav>
        <ul className="flex items-center space-x-2">
          <li>
            <button
              onClick={() => onPageChange(currentPage - 1)}
              disabled={currentPage === 1}
              className="px-3 py-1 bg-gray-700 text-gray-300 rounded hover:bg-gray-600"
            >
              &laquo; Prev
            </button>
          </li>
          {pages.map((page) => (
            <li key={page}>
              <button
                onClick={() => onPageChange(page)}
                className={`px-3 py-1 bg-gray-700 ${
                  page === currentPage
                    ? "border-4 border-blue-900 text-blue-700"
                    : " hover:bg-gray-600 text-gray-300"
                } rounded`}
              >
                {page}
              </button>
            </li>
          ))}
          <li>
            <button
              onClick={() => onPageChange(currentPage + 1)}
              disabled={currentPage === totalPages}
              className="px-3 py-1 bg-gray-700 text-gray-300 rounded hover:bg-gray-600"
            >
              Next &raquo;
            </button>
          </li>
        </ul>
      </nav>
    </div>
  )
}
