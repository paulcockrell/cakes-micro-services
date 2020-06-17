import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Media } from 'reactstrap'

class Table extends Component {
  render() {
    const cakes = this.props.cakes.map(cake => {
      return (
        <Media key={ cake.id }>
          <Media left>
            <Media object src={ cake.image_url } alt={ cake.name } width="64" height="64" />
          </Media>
          <Media body>
            <Media heading>
              { cake.name }
            </Media>
            Want to learn more about this lovely cake? <Link to={ `/show/${cake.id}` }>Click here</Link>
          </Media>
        </Media>
      )
    })

    return (
      <div>
        { cakes }
      </div>
    )
  }
}

export default Table
