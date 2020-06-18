import React, { Component } from 'react'
import { Container, Row, Col, Button } from 'reactstrap'
import { Link } from 'react-router-dom'
import Table from './Components/Table'
import './App.css'

class App extends Component {
  state = {
    cakes: [],
  }

  getCakes = () => {
    // TODO: Use environment variable for api endpoint
    fetch("http://localhost:8080/cakes")
      .then(response => response.json())
      .then(cakes => this.setState({ cakes }))
      .catch(err => console.log(err))
  }

  componentDidMount() {
    this.getCakes()
  }

  render() {
    const { cakes } = this.state

    return (
      <Container className="App">
        <Row>
          <Col>
            <h1>Waracle Cake PWA</h1>
            <h2>All cakes</h2>
            <hr />
          </Col>
        </Row>
        <Row>
          <Col>
            <Table cakes={ cakes } />
          </Col>
        </Row>
        <Row>
          <Col>
            <hr />
            <Button tag={Link} to="/create" color="primary">Create a new cake</Button>
          </Col>
        </Row>
      </Container>
    )
  }
}

export default App;
