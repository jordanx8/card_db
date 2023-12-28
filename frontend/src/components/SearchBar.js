import Form from 'react-bootstrap/Form';
import React from 'react';

function SearchBar({ state, setState, allCards, cardState, setCardState }) {
    // TODO: is state/setstate necessary? could that be combined into the other state shit?
    function searchHandler(event) {
        setState(event.target.value)
        let filteredCards = allCards.filter(function (card) {
            return JSON.stringify(card).toLowerCase().includes(event.target.value.toLowerCase())
        });
        setCardState(filteredCards)
    }

    return (
        <Form.Control
            placeholder="Search"
            aria-label="Search"
            onChange={searchHandler}
            value={state}
        />
    )
}

export default SearchBar