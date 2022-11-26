import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';

function AddPlayerForm() {
    const [formState, setFormState] = useState({
        formFirstName: "",
        formLastName: "",
        formSeasonsPlayed: ""
    });

    function handleInputChange(event) {
        const value = event.target.value
        const name = event.target.id
        setFormState({
            ...formState,
            [name]: value
        });
    }

    return (
        <>
            <h1>Add New Player:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formFirstName">
                    <Form.Label>First Name</Form.Label>
                    <Form.Control placeholder="Enter first name" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formLastName">
                    <Form.Label>Last Name</Form.Label>
                    <Form.Control placeholder="Enter last name" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formSeasonsPlayed">
                    <Form.Label>Season's Played</Form.Label>
                    <Form.Control placeholder="Enter season's played" />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddPlayerForm