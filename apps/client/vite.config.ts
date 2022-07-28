import 'dotenv/config'
import { resolve } from 'path'

/* Plugins */
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'

import Vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import { defineConfig } from 'vite'

const r = (dir: string) => resolve(__dirname, dir)

export default defineConfig(({ mode }) => {
  /**
   * Are we running the app in a test harness?
   */
  const isTestEnv = !!process.env.VITE_CY_TEST || mode === 'test'

  return {
    base: '/',

    build: {
      // < limit to base64 string
      assetsInlineLimit: 10000,
    },

    // pre-bundle the following inclusions
    optimizeDeps: {
      include: ['vue', 'vue-router'],
    },

    plugins: [
      /* Vue */
      Vue({
        template: { transformAssetUrls },
      }),

      /* Auto-import the following modules as compiler macros */
      AutoImport({
        dts: 'src/types/auto-imports.d.ts',
        imports: ['vue', 'vue-router'],
      }),

      quasar({}),
    ],

    /* Alias Resolution */
    resolve: {
      alias: {
        '@': r('./src'),
      },
    },

    server: {
      host: true,
      port: 3000,
    },

    test: {
      environment: 'jsdom',
      globals: true,
      api: false,
    },
  }
})
