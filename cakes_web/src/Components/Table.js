import React, { Component } from 'react'
import { Media } from 'reactstrap'

class Table extends Component {
  render() {
    const cakes = this.props.cakes.map(cake => {
      return (
        <Media>
          <Media left>
            <Media object src={ cake.image_url } alt={ cake.name } width="64" height="64" />
          </Media>
          <Media body>
            <Media heading>
              { cake.name }
            </Media>
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
