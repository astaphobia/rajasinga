import React from 'react'
import { Switch, Route } from 'react-router-dom'

import Home from '../pages/Home'
import LoginPage from '../pages/Login'

function Routes() {
  return (
    <Switch>
      <>
        <Route path="/" exact component={Home} />
        <Route path="/login" component={LoginPage} />
      </>
    </Switch>
  )
}

export default Routes
