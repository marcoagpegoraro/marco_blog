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
          "primary": "#008eff",
          "secondary": "#992d00",
          "accent": "#0049ff",
          "neutral": "#0e0207",
          "base-100": "#282018",
          "info": "#00d1ff",
          "success": "#00e594",
          "warning": "#daa600",
          "error": "#cf0027",
        },
      },
    ],
  },
  plugins: [
    require("@tailwindcss/typography"), 
    require("daisyui")
  ],
}

