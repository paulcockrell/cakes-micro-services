import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import { Media } from 'reactstrap'

class Table extends Component {
  render() {
    const cakes = this.props.cakes || []
    const cakesTable = cakes.map(cake => {
      return (
        <Media key={ cake.id }>
          <Media left>
            <Media object src={ cake.imageUrl } alt={ cake.name } width="64" height="64" />
          </Media>
          <Media body>
            <Media heading>
              { cake.name }
            </Media>
            Want to learn more about this lovely cake? <Link className="button button-primary" to={ `/show/${cake.id}` }>Click here</Link>
          </Media>
        </Media>
      )
    })

    return (
      <div>
        { cakesTable }
      </div>
    )
  }
}

export default Table
