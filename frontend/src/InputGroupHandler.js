import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import React, { useRef } from 'react';
import CheckBoxButton from './CheckBoxButton';

function InputGroupHandler(props) {

    const searchRef = useRef();

    function currentPlayersFilter(player){
        return player.seasonsPlayed.includes('present')
    }

    function searchHandler(event) {
        //TODO: come up with search logic
        props.setSearchString(searchRef.current.value)
        console.log(searchRef.current.value)
    }

    return (
        <>
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="Search"
                    aria-label="Search"
                    aria-describedby="basic-addon2"
                    onChange={searchHandler}
                    ref={searchRef}
                />
                <CheckBoxButton name={"Current Players"} filter={currentPlayersFilter} data={props.players} setData={props.setPlayers} checked={props.checked} setChecked={props.setChecked} originalData={props.originalData} setOriginalData={props.setOriginalData}/>
            </InputGroup>
        </>
    )
}

export default InputGroupHandler;