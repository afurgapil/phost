import React from "react";
import { within, render, screen, fireEvent } from "@testing-library/react";
import Home from "../../pages/index";
import { ThemeProvider } from "../../context/ThemeContext";
const renderWithTheme = (initialDarkMode) => {
  return render(
    <ThemeProvider initialDarkMode={initialDarkMode}>
      <Home />
    </ThemeProvider>
  );
};

describe("Home component", () => {
  test("renders correctly in light mode", () => {
    renderWithTheme(false);
    const mainDiv = screen.getByTestId("home-page");
    expect(mainDiv).toHaveClass("bg-slate-200");
  });

  test("renders correctly in dark mode", () => {
    renderWithTheme(true);
    const mainDiv = screen.getByTestId("home-page");
    expect(mainDiv).toHaveClass("bg-gray-800");
  });

  test("uploads a file and displays the image", async () => {
    renderWithTheme(false);
    const file = new File(["dummy content"], "test.png", { type: "image/png" });
    const input = screen.getByLabelText(/upload your image/i);
    fireEvent.change(input, { target: { files: [file] } });
  });

  test("copies the image URL to the clipboard", async () => {
    renderWithTheme(false);
    const file = new File(["dummy content"], "test.png", { type: "image/png" });
    const input = screen.getByLabelText(/upload your image/i);
    fireEvent.change(input, { target: { files: [file] } });
  });
  test("renders uploaded image and remove button when a file is uploaded", async () => {
    renderWithTheme(false);
    const file = new File(["dummy content"], "test.png", { type: "image/png" });
    const input = screen.getByLabelText(/upload your image/i);
    fireEvent.change(input, { target: { files: [file] } });

    const uploadedImg = await screen.findByAltText("Uploaded Preview");
    expect(uploadedImg).toBeInTheDocument();

    const removeButton = screen.getByTestId("remove-button");
    expect(removeButton).toBeInTheDocument();

    const icon = within(removeButton).getByTestId("delete-icon");
    expect(icon).toBeInTheDocument();

    fireEvent.click(removeButton);
    expect(uploadedImg).not.toBeInTheDocument();
  });
});
