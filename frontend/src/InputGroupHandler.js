import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';
import React, { useState } from 'react';
import { v4 as uuidv4 } from 'uuid';
import CheckBoxButton from './CheckBoxButton';

function InputGroupHandler(props) {

    function currentPlayersFilter(player){
        return player.seasonsPlayed.includes('present')
    }

    function searchHandler(event) {
        //TODO: come up with search logic
        console.log(event.target.value)
    }

    return (
        <React.Fragment key={uuidv4()}>
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="Search"
                    aria-label="Search"
                    aria-describedby="basic-addon2"
                    onChange={searchHandler}
                />
                <CheckBoxButton name={"Current Players"} filter={currentPlayersFilter} data={props.players} setData={props.setPlayers} checked={props.checked} setChecked={props.setChecked} originalData={props.originalData} setOriginalData={props.setOriginalData}/>
            </InputGroup>
        </React.Fragment>
    )
}

export default InputGroupHandler;