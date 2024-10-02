/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        caveat: ["Caveat", "sans-serif"],
        caveatSemiBold: ["Caveat-SemiBold", "sans-serif"],
        caveatBold: ["Caveat-Bold", "sans-serif"],
        caveatMedium: ["Caveat-Medium", "sans-serif"],
      },
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
      },
      maxWidth: {
        "11/12": "91.666667%",
      },
    },
  },
  plugins: [],
};
