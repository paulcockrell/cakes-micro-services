import React, { Component } from 'react'
import { Container, Row, Col } from 'reactstrap'
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

  addCakeToState = (cake) => {
    this.setState(prevState => ({
      cakes: [...prevState.cakes, cake]
    }))
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
          </Col>
        </Row>
        <Row>
          <Col>
            <Table cakes={ cakes } />
          </Col>
        </Row>
      </Container>
    )
  }
}

export default App;
