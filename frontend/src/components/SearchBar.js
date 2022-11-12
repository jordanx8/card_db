import Form from 'react-bootstrap/Form';
import React, { useRef } from 'react';

function SearchBar(props) {
    const searchRef = useRef();

    function searchHandler(event) {
        //TODO: come up with search logic
        console.log(searchRef.current.value)
    }

    return (
        <Form.Control
            placeholder="Search"
            aria-label="Search"
            aria-describedby="basic-addon2"
            onChange={searchHandler}
            ref={searchRef}
        />
    )
}

export default SearchBar