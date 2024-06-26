/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "dark-midnight": "#0e2336",
        "slate-light": "#182c3e",
        "deep-blue": "#0f2b3c",
      },
      width: {
        "17/100": "17%",
        "21/100": "21%",
      },
    },
  },
  plugins: [],
}
