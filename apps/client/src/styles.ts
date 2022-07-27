import { injectGlobal } from '@magister_zito/vue3-styled-components'

injectGlobal`
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  #app {
    height: 100vh;
  }

  body {
    background-color: #fefefe;
    font-family: -apple-system, BlinkMacSystemFont, sans-serif;
    padding: 0;
    margin: 0;
    height: 100vh;
    width: 100vw;
    overflow: hidden;
  }

  a {
    text-decoration: none;
    color: inherit;
  }
`
