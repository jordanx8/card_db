import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';
import { ADD_PLAYER_MUTATION } from '../../util/mutations';
import { gql, useMutation } from '@apollo/client';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function AddPlayerForm() {
    const [addPlayer, {data, loading, error}] = useMutation(ADD_PLAYER_MUTATION);


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

    async function handleSubmit() {
        const id = toast.loading("Adding player...")
        const {data} = await addPlayer({ variables: { firstName: formState.formFirstName, lastName: formState.formLastName, seasonsPlayed: formState.formSeasonsPlayed } })
        toast.update(id, {render: data.AddPlayer, type:"success", isLoading: false, autoClose: 5000, closeOnClick: true, draggable: true})
    }

    return (
        <>
            <ToastContainer
                position="top-right"
                draggable
                autoClose={5000}
                closeOnClick
            />
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
                <Button variant="primary" type="button" onClick={handleSubmit}>
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddPlayerForm