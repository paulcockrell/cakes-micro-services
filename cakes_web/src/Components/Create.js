import React, { Component } from 'react'
import { Container, Row, Col, Form, Button, FormGroup, Label, Input } from 'reactstrap'
import { Link } from 'react-router-dom'

class Create extends Component {
  state = {
    name: '',
    comment: '',
    imageUrl: '',
    yumFactor: 0,
  }

  handleChange = e => {
    this.setState({ [e.target.name]: e.target.value })
  }

  createCake = () => {
    const { name, comment, imageUrl, yumFactor } = this.state
    fetch(`http://localhost:8080/cakes`, {
      method: 'POST',
      body: JSON.stringify({ name, comment, imageUrl, yumFactor: Number(yumFactor) }),
    })
    .then(rsp => rsp.json())
    .then(cake => this.props.history.push(`/show/${cake.id}`))
    .catch(err => console.error("Error creating cake", err))
  }

  handleSubmit = (evt) => {
    evt.preventDefault()
    this.createCake()
  }

  render() {
    const { name, comment, imageUrl, yumFactor } = this.state

    return(
      <Container>
      <Row>
        <Col>
          <h1>Waracle Cake PWA</h1>
          <h2>Create cake</h2>
          <hr />
        </Col>
      </Row>
        <Row>
          <Col>
            <Form onSubmit={ this.handleSubmit.bind(this) }>
              <FormGroup>
                <Label for="name">Name</Label>
                <Input type="text" name="name" id="name" value={ name } onChange={ this.handleChange } />
              </FormGroup>
              <FormGroup>
                <Label for="yumFactor">Yum Factor</Label>
                <Input type="select" name="yumFactor" id="yumFactor" value={yumFactor} onChange={ this.handleChange }>
                  {
                    [...Array(10).keys()].map(idx => {
                      return <option key={ idx } value={ idx }>{ idx }</option> 
                    })
                  }
                </Input>
              </FormGroup>
              <FormGroup>
                <Label for="comment">Comment</Label>
                <Input type="textarea" name="comment" id="comment" value={ comment } onChange={ this.handleChange } />
              </FormGroup>
              <FormGroup>
                <Label for="imageUrl">Image</Label>
                <Input type="text" name="imageUrl" id="imageUrl" value={ imageUrl } onChange={ this.handleChange } />
              </FormGroup>
              <Button type="submit" value="submit" color="primary">Submit</Button>{' '}
              <Button tag={Link} to="/" color="secondary">Cancel</Button>
            </Form>
          </Col>
        </Row>
      </Container>
    )
  }
}

export default Create
