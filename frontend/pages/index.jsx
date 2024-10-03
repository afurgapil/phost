import Layout from "../components/Layout";
import React, { useState, useEffect } from "react";
import { useTheme } from "../context/ThemeContext";
import { RiDeleteBin7Fill } from "react-icons/ri";

const Home = () => {
  const { darkMode } = useTheme();
  const [file, setFile] = useState(null);
  const [uploadedImage, setUploadedImage] = useState(null);

  const handleFileUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
      setFile(file);
      const reader = new FileReader();
      reader.onload = () => {
        setUploadedImage(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <Layout title="Phost">
      <main>
        <div
          data-testid="home-page"
          className={`min-h-screen w-full transition duration-300 ease-in-out flex flex-col justify-start items-center ${
            darkMode ? "bg-gray-800" : "bg-slate-200"
          }`}
        >
          <div className="flex flex-col items-center mt-48">
            <div
              className={`${
                darkMode ? "border-white" : "border-black"
              } mb-6 border px-16 py-8  rounded-md shadow-md w-full max-w-md `}
            >
              <label
                className={`${
                  darkMode ? "text-slate-200" : "text-gray-700"
                } block font-semibold text-2xl  m-2 p-4 rounded-md`}
                htmlFor="imageUpload"
              >
                Upload Your Image
              </label>
              <input
                id="imageUpload"
                accept="image/*"
                className="block w-full text-sm text-gray-500 bg-slate-300 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold hover:file:bg-blue-100 m-2 rounded-[9999px]"
                type="file"
                onChange={handleFileUpload}
              />
            </div>

            {uploadedImage && (
              <div className="mt-4">
                <img
                  src={uploadedImage}
                  alt="Uploaded Preview"
                  className="max-w-full h-auto border rounded-md"
                />
                <button
                  className="flex items-center text-red-500 mt-2"
                  onClick={() => setUploadedImage(null)}
                >
                  <RiDeleteBin7Fill className="mr-2" />
                  Remove Image
                </button>
              </div>
            )}
          </div>
        </div>
      </main>
    </Layout>
  );
};

export default Home;
