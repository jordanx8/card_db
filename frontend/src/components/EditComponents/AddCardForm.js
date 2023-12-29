import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState } from 'react';
import ListDropdown from './ListDropdown';
import { useQuery, useMutation } from "@apollo/client";
import { GET_PLAYERS_QUERY } from '../../util/queries';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { filterPlayerListDownTo } from '../../util/util'
import { ADD_CARD_MUTATION } from '../../util/mutations';

function AddCardForm() {
    const [formState, setFormState] = useState({
        playerName: "",
        season: "",
        manufacturer: "",
        set: "",
        parallel: "",
        cardNumber: "",
        imageLink: "",
        tableName: "Pelicans",
        notes: []
    });

    const { data: playersData, loading: playersLoading, error: playersError } = useQuery(GET_PLAYERS_QUERY, {
        onCompleted: (data) => {
            setSelectedPlayer(data.players[0].firstName + " " + data.players[0].lastName)
            setFormState({
                ...formState,
                "playerName": data.players[0].firstName + " " + data.players[0].lastName,
                "season": data.players[0].seasons[0],
            });
        }
    });
    const [addCard, { data, loading, error }] = useMutation(ADD_CARD_MUTATION);

    function handleInputChange(event) {
        const value = event.target.value
        const name = event.target.id
        if (name === "formNotes") {
            if (event.target.checked) {
                let newArray = [...formState.notes, event.target.getAttribute('controlid')]
                setFormState({
                    ...formState,
                    "notes": newArray,
                });
            } else {
                let newArray = formState.notes.filter(x => x !== event.target.getAttribute('controlid'))
                setFormState({
                    ...formState,
                    "notes": newArray,
                });
            }
        } else {
            if (name === "playerName") {
                setFormState({
                    ...formState,
                    "season": playersData.players.filter(filterPlayerListDownTo(value))[0].seasons[0]
                });
            }
            setFormState({
                ...formState,
                [name]: value
            });
        }
    }
    const [selectedPlayer, setSelectedPlayer] = useState("");
    const [wChecked, setWChecked] = useState(false);
    const [numChecked, setNumChecked] = useState(false);


    function PlayerDropdownHandler(event) {
        setSelectedPlayer(event.target.value);
    }

    function numberHandler(event) {
        setNumChecked(!numChecked)
    }

    function wHandler(event) {
        setWChecked(!wChecked)
    }

    async function handleSubmit() {
        const id = toast.loading("Adding card...")
        const { data } = await addCard({
            variables: {
                playerName: formState.playerName,
                season: formState.season,
                manufacturer: formState.manufacturer,
                set: formState.set,
                parallel: formState.parallel,
                cardNumber: formState.cardNumber,
                imageLink: formState.imageLink,
                tableName: formState.tableName,
                notes: formState.notes,
                team: formState.team
            }
        })
        toast.update(id, { render: data.AddCard, type: "success", isLoading: false, autoClose: 5000, closeOnClick: true, draggable: true })
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
            <h1>Add New Card:</h1>
            <Form>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="playerName">
                    <Form.Label>Select a Player</Form.Label>
                    <Form.Select value={selectedPlayer} onChange={PlayerDropdownHandler}>
                        {playersData.players.filter(x => x.seasons.length !== 0).map(values => (
                            <ListDropdown text={values.firstName + " " + values.lastName} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="team">
                    <Form.Label>Team</Form.Label>
                    <Form.Control placeholder="pels?" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="season">
                    <Form.Label>Select a Season</Form.Label>
                    <Form.Select>
                        {playersData.players.filter(filterPlayerListDownTo(selectedPlayer))[0].seasons.map(season => (
                            <ListDropdown text={season} />
                        ))
                        }
                    </Form.Select>
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="manufacturer">
                    <Form.Label>Manufacturer</Form.Label>
                    <Form.Control placeholder="probably Panini" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="set">
                    <Form.Label>Set</Form.Label>
                    <Form.Control placeholder="Enter card's set" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="parallel">
                    <Form.Label>Parallel</Form.Label>
                    <Form.Control placeholder="Enter parallel" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="cardNumber">
                    <Form.Label>Card Number</Form.Label>
                    <Form.Control placeholder="Enter card number" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="imageLink">
                    <Form.Label>Image Link</Form.Label>
                    <Form.Control placeholder="Enter image link" />
                </Form.Group>
                <Form.Group onChange={handleInputChange} className="mb-3" controlId="formNotes">
                    <Form.Check inline type="checkbox" label="RC" controlId="RC" />
                    <Form.Check inline type="checkbox" label="Sticker" controlId="Sticker" />
                    <Form.Check inline type="checkbox" label="College" controlId="College" />
                    <Form.Check inline type="checkbox" label="Pre-RC" controlId="Pre-RC" />
                    <Form.Check inline type="checkbox" label="Auto" controlId="Auto" />
                    <Form.Check inline type="checkbox" label="Patch" controlId="Patch" />
                    {/* TODO: Get the below forms thangs to work */}
                    <Form.Check value={wChecked} onChange={wHandler} type="checkbox" label="w/ Others" controlId="w/ Others" />
                    {wChecked &&
                        <Form.Control placeholder="who? (comma-separated)" />}
                    <Form.Check value={numChecked} onChange={numberHandler} type="checkbox" label="#'d" />
                    {numChecked &&
                        <Form.Control placeholder="Enter serial number" />}
                </Form.Group>
                <Button variant="primary" type="button" onClick={handleSubmit}>
                    Submit
                </Button>
            </Form>
        </>
    )
}

export default AddCardForm