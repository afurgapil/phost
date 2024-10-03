import React from "react";
import { render, screen } from "@testing-library/react";
import Layout from "../../components/Layout";

jest.mock("next/head", () => ({
  __esModule: true,
  default: ({ children }) => <>{children}</>,
}));

jest.mock("../../components/Header", () => () => <div>Mocked Header</div>);
jest.mock("../../components/Footer", () => () => <div>Mocked Footer</div>);

describe("Layout component", () => {
  test("renders the default title", () => {
    render(<Layout title="Phost">Test Content</Layout>);
    const titleElement = screen.getByText("Phost");
    expect(titleElement).toBeInTheDocument();
  });

  test("renders the provided title", () => {
    const customTitle = "Custom Page Title";
    render(<Layout title={customTitle}>Test Content</Layout>);
    const titleElement = screen.getByText(customTitle);
    expect(titleElement).toBeInTheDocument();
  });

  test("renders the Header component", () => {
    render(<Layout>Test Content</Layout>);
    expect(screen.getByText("Mocked Header")).toBeInTheDocument();
  });

  test("renders the Footer component", () => {
    render(<Layout>Test Content</Layout>);
    expect(screen.getByText("Mocked Footer")).toBeInTheDocument();
  });

  test("renders children content", () => {
    render(
      <Layout>
        <div>Test Child Content</div>
      </Layout>
    );
    expect(screen.getByText("Test Child Content")).toBeInTheDocument();
  });

  test("renders meta tags correctly", () => {
    render(<Layout>Test Content</Layout>);
    const metaDescription = document.querySelector("meta[name='description']");
    const metaViewport = document.querySelector("meta[name='viewport']");

    expect(metaDescription).toHaveAttribute("content", "Free image host");
    expect(metaViewport).toHaveAttribute(
      "content",
      "width=device-width, initial-scale=1"
    );
  });

  test("renders the favicon link", () => {
    render(<Layout>Test Content</Layout>);
    const faviconLink = document.querySelector("link[rel='icon']");
    expect(faviconLink).toHaveAttribute("href", "/favicon.ico");
  });
});
