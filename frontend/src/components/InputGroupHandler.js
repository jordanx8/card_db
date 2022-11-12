
import InputGroup from 'react-bootstrap/InputGroup';
import React from 'react';
import FilterButton from './FilterButton';
import { currentPlayersFilter } from '../util/filters'
import SearchBar from './SearchBar';

function InputGroupHandler(props) {

    return (
        <>
            <InputGroup className="mb-3">
                <SearchBar />
                <FilterButton name={"Current Players"} filter={currentPlayersFilter} data={props.players} setData={props.setPlayers}/>
            </InputGroup>
        </>
    )
}

export default InputGroupHandler;