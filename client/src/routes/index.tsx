import React from 'react'
import { Switch, Route } from 'react-router-dom'

import Home from '../pages/Home'
import LoginPage from '../pages/Login'
import AboutPage from '../pages/About'
import NotFound from '../pages/NotFound'

function Routes() {
  return (
    <Switch>
      <React.Fragment>
        <Route path="/" exact component={Home} />
        <Route path="/login" component={LoginPage} />
        <Route path="/about" component={AboutPage} />
      </React.Fragment>
    </Switch>
  )
}

export default Routes
