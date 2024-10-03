import React from "react";
import { render, screen } from "@testing-library/react";
import { ThemeProvider } from "../../context/ThemeContext";
import NotFoundPage from "../../pages/nf";

const renderWithTheme = (darkMode) => {
  return render(
    <ThemeProvider initialDarkMode={darkMode}>
      <NotFoundPage />
    </ThemeProvider>
  );
};

describe("NotFoundPage", () => {
  test("renders correctly in light mode", () => {
    renderWithTheme(false);
    const mainDiv = screen.getByTestId("not-found-page");
    expect(mainDiv).toHaveClass("bg-white");
    expect(mainDiv).toHaveClass("text-gray-800");
  });

  test("renders correctly in dark mode", () => {
    renderWithTheme(true);
    const mainDiv = screen.getByTestId("not-found-page");
    expect(mainDiv).toHaveClass("bg-gray-800");
    expect(mainDiv).toHaveClass("text-white");
  });

  test("displays correct content", () => {
    renderWithTheme(false);
    expect(screen.getByText("404")).toBeInTheDocument();
    expect(
      screen.getByText("Sorry, the image you are looking for was not found")
    ).toBeInTheDocument();
  });

  test("contains a link to the home page", () => {
    renderWithTheme(false);
    const homeLink = screen.getByText("Click here to return to the home page");
    expect(homeLink).toBeInTheDocument();
    expect(homeLink.closest("a")).toHaveAttribute("href", "/");
  });
});
