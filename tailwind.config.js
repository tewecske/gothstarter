/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ"], // maybe add "./**/*.go"
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark"]
  }
}

