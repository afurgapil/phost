import { useTheme } from "../context/ThemeContext";

export default function Footer() {
  const { darkMode } = useTheme();

  return (
    <footer
      className={`${
        darkMode ? "bg-gray-900 text-white" : "bg-slate-300"
      } flex justify-center items-center p-1`}
    >
      <p>&copy; 2024 Phost. All Rights Reserved.</p>
    </footer>
  );
}
