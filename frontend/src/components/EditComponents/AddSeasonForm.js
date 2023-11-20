import Button from 'react-bootstrap/Button';
import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import ListDropdown from './ListDropdown';
import { useQuery, useMutation } from "@apollo/client";
import { GET_PLAYERS_QUERY } from '../../util/queries';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { ADD_SEASON_MUTATION } from '../../util/mutations';


function AddSeasonForm() {
    const [formState, setFormState] = useState({
        formFirstName: "",
        formLastName: "",
        formSeason: "",
    });

    const { data: playersData, loading: playersLoading, error: playersError } = useQuery(GET_PLAYERS_QUERY, {
        onCompleted: (data) => {
            setFormState({
                ...formState,
                "formFirstName": data.players[0].firstName,
                "formLastName": data.players[0].lastName
            });
        }
      });
    const [addSeason, { data, loading, error }] = useMutation(ADD_SEASON_MUTATION);

    function handleInputChange(event) {
        const value = event.target.value
        const name = event.target.id
        if (name === "formPlayer") {
            const names = value.split(" ");
            setFormState({
                ...formState,
                "formFirstName": names[0],
                "formLastName": names[1]
            });
        } else {
            setFormState({
                ...formState,
                [name]: value
            });
        }
    }


    async function handleSubmit() {
        const id = toast.loading("Adding season...")
        const { data } = await addSeason({ variables: { firstName: formState.formFirstName, lastName: formState.formLastName, season: formState.formSeason } })
        toast.update(id, { render: data.AddSeason, type: "success", isLoading: false, autoClose: 5000, closeOnClick: true, draggable: true })
    }




    if (playersLoading) return "Loading...";
    if (playersError) return <pre>{playersError.message}</pre>

    return (
        <>
            <ToastContainer
                position="top-right"
                draggable
                autoClose={5000}
                closeOnClick
            />
            <h1>Add Season to Existing Player:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formPlayer">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select>
                        {playersData.players.map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formSeason">
                    <Form.Label>Season</Form.Label>
                    <Form.Control placeholder="Enter seasons to add to player (comma separated)" />
                </Form.Group>
                <Button variant="primary" type="button" onClick={handleSubmit}>
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddSeasonForm