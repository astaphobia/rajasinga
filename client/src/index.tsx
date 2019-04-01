import 'bulma/css/bulma.css'
import React from 'react'
import { render } from 'react-dom'
import { BrowserRouter } from 'react-router-dom'

import Routes from './routes'
import TopNav from './components/TopNav'

render(
  <div className="container">
    <BrowserRouter>
      <>
        <TopNav />
        <Routes />
      </>
    </BrowserRouter>
  </div>,
  document.getElementById('root')
)
