import React, { Component } from 'react'
import { Container, Row, Col, Media, Button } from 'reactstrap'
import { Link } from 'react-router-dom'

const endpoint = process.env.REACT_APP_API_ENDPOINT || "http://localhost:8080/cakes"

class Show extends Component {
  state = {
    cake: null,
  }

  getCake = (id) => {
    // TODO: Use environment variable for api endpoint
    fetch(`${endpoint}/${id}`)
      .then(response => response.json())
      .then(cake => this.setState({ cake }))
      .catch(err => console.log(err))
  }

  deleteCake = (id) => {
    fetch(`${endpoint}/${id}`, {
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
            <h1>Waracle Cake PWA</h1>
            <h2>Show cake</h2>
            <hr />
          </Col>
        </Row>
        <Row>
          <Col>
              <Media object src={ cake.imageUrl } alt={ cake.name } width="300" height="300" />
          </Col>
        </Row>
        <Row>
          <Col>
              ID: { cake.id } <br/>
          </Col>
        </Row>
        <Row>
          <Col>
              Name: { cake.name } <br/>
          </Col>
        </Row>
        <Row>
          <Col>
              Comment: { cake.comment } <br/>
          </Col>
        </Row>
        <Row>
          <Col>
              Yum factor: { cake.yumFactor } <br/>
          </Col>
        </Row>
        <Row>
          <Col>
              <Button tag={ Link } to={ `/edit/${cake.id}` } color="primary">Edit</Button>{' '}
              <Button tag={ Link } onClick={ this.deleteCake.bind(this, cake.id) } color="danger">Delete</Button>{' '}
              <Button tag={ Link } to="/" color="secondary">Back</Button>
          </Col>
        </Row>
      </Container>
    )
  }
}

export default Show
