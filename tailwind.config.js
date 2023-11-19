/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/*.django", 
    "./views/**/*.django"
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          "primary": "#305c96",
          "secondary": "#8199d3",
          "accent": "#016568",
          "neutral": "#1c191f",
          "base-100": "#424153",
          "info": "#afc2e4",
          "success": "#00D389",
          "warning": "#a3730a",
          "error": "#f54d7c",
        },
      },
    ],
  },
  plugins: [
    require("@tailwindcss/typography"), 
    require("daisyui")
  ],
}

