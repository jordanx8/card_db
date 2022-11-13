import Form from 'react-bootstrap/Form';
import React from 'react';

function SearchBar({ state, setState }) {
    function searchHandler(event) {
        setState(event.target.value)
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