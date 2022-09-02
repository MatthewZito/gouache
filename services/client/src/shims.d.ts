declare interface Window {}

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<
    Record<string, never>,
    Record<string, never>,
    any
  >
  export default component
}

interface ImportMeta {
  env: {
    VITE_AUTH_API?: string
    VITE_REPORTING_API?: string
    VITE_RESOURCE_API?: string
    VITE_CY_TEST?: string
  }
}
