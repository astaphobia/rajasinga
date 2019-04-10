import React from 'react'
import { Link } from 'react-router-dom'

function TopNav() {
  return (
    <div className="hero-head">
      <nav className="navbar">
        <div className="container">
          <div className="navbar-brand">
            <Link className="navbar-item" to={'/'}>
              <h1 className="brand">RAJASINGA</h1>
            </Link>
            <span className="navbar-burger burger" data-target="navbarMenu">
              <span />
              <span />
              <span />
            </span>
          </div>
          <div id="navbarMenu" className="navbar-menu">
            <div className="navbar-end">
              <span className="navbar-item">
                <Link className="button is-white is-outlined" to={`/`}>
                  <span className="icon">
                    <i className="fa fa-home" />
                  </span>
                  <span>Home</span>
                </Link>
              </span>
              <span className="navbar-item">
                <Link className="button is-white is-outlined" to={`/about`}>
                  <span className="icon">
                    <i className="fa fa-superpowers" />
                  </span>
                  <span>About Me</span>
                </Link>
              </span>
            </div>
          </div>
        </div>
      </nav>
    </div>
  )
}

export default TopNav
