import Link from "next/link";
import { useTheme } from "../context/ThemeContext";

export default function Header() {
  const { darkMode, toggleDarkMode } = useTheme();

  return (
    <header
      className={`${darkMode ? "bg-gray-900 text-white" : "bg-slate-300"}`}
    >
      <nav className="flex flex-row justify-around items-center mx-4">
        <Link href="/" className="font-caveatBold text-5xl">
          Phost
        </Link>
        <div className="flex flex-row justify-end w-full p-4">
          <button
            onClick={toggleDarkMode}
            className={`text-5xl ${
              darkMode ? "text-yellow-300" : "text-gray-700"
            } transition duration-300 ease-in-out`}
          >
            {darkMode ? "🌞" : "🌙"}
          </button>
        </div>
      </nav>
    </header>
  );
}
