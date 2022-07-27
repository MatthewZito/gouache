import styled from '@magister_zito/vue3-styled-components'
import { RouterLink } from 'vue-router'

export const Header = styled.div`
  display: flex;
  width: 100%;
  height: 3.5rem;
  background-color: black;
  padding: 0.25rem;
`

export const HeaderBtn = styled.button.attrs({
  type: 'button',
})`
  color: red;
  width: 50px;
  border-radius: 100%;
  cursor: pointer;
`

export const NavDrawer = styled.nav`
  position: absolute;
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 300px;
  background-color: gray;
  transition: all 0.5s ease;
`

export const NavDrawerItem = styled(RouterLink, { isActive: Boolean })`
  position: absolute;
  z-index: 9;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  height: 60px;
  padding: 1rem;
  border-bottom: 1px solid black;
  cursor: pointer;

  background-color: ${({ isActive }) => (isActive ? 'darkgray' : 'gray')};

  &:hover {
    background-color: darkgray;
  }

  & a {
    color: inherit;
    text-decoration: none;
  }
`

export const MainWrapper = styled.main`
  display: flex;
  padding: 1rem;
  align-items: center;
  position: relative;
  width: 100%;
`
