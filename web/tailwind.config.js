/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/templates/*.go.html", "./**/static/js/*.js"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms")],
};
