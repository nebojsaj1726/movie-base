import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      hooks: path.join(__dirname, "./src/hooks"),
      assets: path.join(__dirname, "./src/assets"),
      components: path.join(__dirname, "./src/components"),
      types: path.join(__dirname, "./src/types"),
      utils: path.join(__dirname, "./src/utils"),
    },
  },
})
