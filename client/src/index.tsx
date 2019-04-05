import 'bulma/css/bulma.css'
import './assets/styles.scss'
import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter, HashRouter } from 'react-router-dom'

import Routes from './routes'
import TopNav from './components/TopNav'

render(
  <div className="hero is-warning is-fullheight">
    <BrowserRouter>
      <>
        <TopNav />
        <Routes />
      </>
    </BrowserRouter>
  </div>,
  document.getElementById('root')
)
