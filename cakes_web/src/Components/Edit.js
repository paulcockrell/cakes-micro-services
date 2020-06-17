import React, { Component } from 'react'
import { Container, Row, Col, Form, Button, FormGroup, Label, Input, FormText } from 'reactstrap'
import { Link } from 'react-router-dom'

class Edit extends Component {
  state = {
      id: 0,
      name: '',
      comment: '',
      imageUrl: '',
      yumFactor: 0,
  }

  handleChange = e => {
      this.setState({ [e.target.name]: e.target.value })
  }

  getCake = (id) => {
    // TODO: Use environment variable for api endpoint
    fetch(`http://localhost:8080/cakes/${id}`)
      .then(response => response.json())
      .then(cake => {
          const { id, name, comment, imageUrl, yumFactor } = cake
          this.setState({ id, name, comment, imageUrl, yumFactor: yumFactor })
      })
      .catch(err => console.log(err))
  }

  updateCake = () => {
      const { id, name, comment, imageUrl, yumFactor } = this.state
      fetch(`http://localhost:8080/cakes/${id}`, {
          method: 'PUT',
          body: JSON.stringify({ name, comment, imageUrl, yumFactor: Number(yumFactor) }),
      })
      .then(_ => this.props.history.push(`/show/${id}`) )
      .catch(err => console.log(err))
  }

  handleSubmit = (evt) => {
      evt.preventDefault()
      this.updateCake()
  }

  componentDidMount() {
    const { id } = this.props.match.params
    this.getCake(id)
  }

  render() {
    const { name, comment, imageUrl, yumFactor } = this.state

    return(
        <Container>
            <Row>
                <Col>
                    <h1>Edit Cake</h1>
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
                        <Button type="submit" value="submit">Submit</Button>
                    </Form>
                </Col>
            </Row>
        </Container>
    )
  }
}

export default Edit