import React from 'react'

function NotFound({ location }: { location: any }) {
  return (
    <section>
      <div className="hero-body">
        <div className="container has-text-centered">
          <h1 className="center">
            404 Not Found <code>{location.pathname}</code>
          </h1>
        </div>
      </div>
    </section>
  )
}

export default NotFound
