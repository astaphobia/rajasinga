import 'bulma/css/bulma.css'
import React from 'react'
import { render } from 'react-dom'

import Routes from './routes'

render(
    <div className="container">
        <Routes />
    </div>,
    document.getElementById('root')
)
