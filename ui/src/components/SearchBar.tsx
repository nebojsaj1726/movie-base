import { ChangeEvent, useState } from "react"
import { SearchResults } from "./SearchResults"
import { useDebounce } from "hooks/useDebounce"

export const SearchBar = () => {
  const [searchInput, setSearchInput] = useState("")
  const debouncedSearchInput = useDebounce(searchInput, 300)

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSearchInput(event.target.value)
  }

  const clearSearchInput = () => {
    setSearchInput("")
  }

  return (
    <div className="flex-grow max-w-xl relative">
      <input
        type="text"
        placeholder="Search"
        value={searchInput}
        onChange={handleInputChange}
        className="bg-slate-light w-full text-gray-200 px-4 py-2 rounded-md focus:outline-none 
        focus:bg-gray-700"
      />
      {debouncedSearchInput && (
        <SearchResults
          searchInput={debouncedSearchInput}
          onLinkClick={clearSearchInput}
        />
      )}
    </div>
  )
}
