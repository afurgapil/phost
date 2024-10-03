import React from "react";
import "../styles/globals.css";
import "../styles/fonts.css";
import { ThemeProvider } from "../context/ThemeContext";
function MyApp({ Component, pageProps }) {
  return (
    <ThemeProvider initialDarkMode={true}>
      <Component {...pageProps} />
    </ThemeProvider>
  );
}

export default MyApp;
