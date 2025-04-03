/// <reference types="vite/client" />

interface Window {
  /**
   * window._env_ is defined in `./public/config/config.js`. It contains
   * variables that are required for the application to run in production.
   */
  _env_: Partial<ImportMetaEnv>
}

interface ImportMetaEnv {
  /**
   * Address of the help page.
   */
  readonly HELP_PAGE: string
}
