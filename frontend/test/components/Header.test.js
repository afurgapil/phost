import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import Header from "../../components/Header";
import { ThemeProvider } from "../../context/ThemeContext";

jest.mock("next/link", () => {
  return ({ children, href }) => {
    return <a href={href}>{children}</a>;
  };
});

const renderWithTheme = (component, { initialDarkMode = false } = {}) => {
  return render(
    <ThemeProvider initialDarkMode={initialDarkMode}>{component}</ThemeProvider>
  );
};

describe("Header", () => {
  test("renders correctly in light mode", () => {
    renderWithTheme(<Header />);
    const header = screen.getByRole("banner");
    expect(header).toHaveClass("bg-slate-300");
    expect(header).not.toHaveClass("bg-gray-900");
    expect(header).not.toHaveClass("text-white");
  });

  test("renders correctly in dark mode", () => {
    renderWithTheme(<Header />, { initialDarkMode: true });
    const header = screen.getByRole("banner");
    expect(header).toHaveClass("bg-gray-900");
    expect(header).toHaveClass("text-white");
    expect(header).not.toHaveClass("bg-slate-300");
  });

  test("renders logo with correct text and link", () => {
    renderWithTheme(<Header />);
    const logo = screen.getByText("Phost");
    expect(logo).toBeInTheDocument();
    expect(logo.tagName).toBe("A");
    expect(logo).toHaveAttribute("href", "/");
  });

  test("renders theme toggle button", () => {
    renderWithTheme(<Header />);
    const toggleButton = screen.getByRole("button");
    expect(toggleButton).toBeInTheDocument();
  });

  test("theme toggle button changes from light to dark", () => {
    renderWithTheme(<Header />);
    const toggleButton = screen.getByRole("button");
    expect(toggleButton).toHaveTextContent("🌙");
    fireEvent.click(toggleButton);
    expect(toggleButton).toHaveTextContent("🌞");
  });

  test("theme toggle button changes from dark to light", () => {
    renderWithTheme(<Header />, { initialDarkMode: true });
    const toggleButton = screen.getByRole("button");
    expect(toggleButton).toHaveTextContent("🌞");
    fireEvent.click(toggleButton);
    expect(toggleButton).toHaveTextContent("🌙");
  });
});
