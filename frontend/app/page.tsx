
export default function Home() {
  return (
      <div className="grid grid-rows-[auto_1fr_auto] items-center justify-items-center min-h-screen p-8 gap-16 sm:p-20 font-sans bg-white dark:bg-black text-black dark:text-white">
        <main className="flex flex-col gap-8 row-start-2 items-center text-center sm:text-left max-w-3xl">

          <h1 className="text-3xl sm:text-4xl font-bold text-center">
            Welcome to <span className="text-blue-600">IdeaHub</span>
          </h1>
          <p className="text-base sm:text-lg text-gray-700 dark:text-gray-300">
            IdeaHub is an AI-powered platform to manage, organize, and summarize your innovative ideas.
            You can create, view, update, and delete ideas, and even get AI-generated summaries and tags!
          </p>

          <div className="flex gap-4 flex-col sm:flex-row items-center">
            <a
                className="rounded-full bg-blue-600 hover:bg-blue-700 text-white font-medium h-10 sm:h-12 px-6 transition-colors"
                href="/ideas"
            >
              Get Started
            </a>
            <a
                className="rounded-full border border-gray-300 dark:border-gray-600 text-sm sm:text-base px-6 py-2 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                href="https://github.com/your-repo"
                target="_blank"
                rel="noopener noreferrer"
            >
              View Source Code
            </a>
          </div>
        </main>

        <footer className="row-start-3 text-xs text-gray-500 dark:text-gray-400 text-center">
          <p>© {new Date().getFullYear()} IdeaHub. All rights reserved.</p>
        </footer>
      </div>
  );
}
