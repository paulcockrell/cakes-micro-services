import React, { Component } from 'react'
import { Container, Row, Col, Media, Button } from 'reactstrap'
import { Link } from 'react-router-dom'

class Show extends Component {
  state = {
    cake: null,
  }

  getCake = (id) => {
    // TODO: Use environment variable for api endpoint
    fetch(`http://localhost:8080/cakes/${id}`)
      .then(response => response.json())
      .then(cake => this.setState({ cake }))
      .catch(err => console.log(err))
  }

  deleteCake = (id) => {
    fetch(`http://localhost:8080/cakes/${id}`, {
      method: 'DELETE'
    })
    .then(_ => this.props.history.push("/"))
    .catch(err => console.error("Error deleting cake", err))
  }

  componentDidMount() {
    const { id } = this.props.match.params
    this.getCake(id)
  }

  render() {
    const { cake } = this.state
    if (!cake) return <i>Loading...</i>

    return (
      <Container className="App">
        <Row>
          <Col>
            <h1>Viewing cake ID {cake.id}</h1>
          </Col>
        </Row>
        <Row>
          <Col>
            <Media key={ cake.id }>
              <Media left>
                <Media object src={ cake.image_url } alt={ cake.name } width="400" height="400" />
              </Media>
              <Media body>
                <Media heading>
                  { cake.name }
                </Media>
                ID: { cake.id } <br/>
                Name: { cake.name } <br/>
                Comment: { cake.comment } <br/>
                Yum factor: { cake.yumFactor } <br/>
                <Link to={`/edit/${cake.id}`}>Edit</Link>
                <Button onClick={ this.deleteCake.bind(this, cake.id) }>Delete</Button>
                <Link to="/">Back</Link>
              </Media>
            </Media>
          </Col>
        </Row>
      </Container>
    )
  }
}

export default Show
