import React, { useState } from 'react';
import PlayerRow from "./PlayerRow";
import Table from 'react-bootstrap/Table';
import InputGroup from 'react-bootstrap/InputGroup';
import FilterButton from './FilterButton';
import SearchBar from './SearchBar';
import { nameToFilter } from '../util/util';
import * as playerData from '../test_player_data.json';
import * as cardData from '../test_card_data.json';

function CardTable({ tableName }) {
  const [cards, setCards] = useState(cardData.default);
  const [players, setPlayers] = useState(playerData.default);
  const [searchTerm, setSearchTerm] = useState('');
  const [currentPlayers, setCurrentPlayers] = useState(false);
  const [rookieCards, setRookieCards] = useState(false);
  const [autograph, setAutograph] = useState(false);
  const [patch, setPatch] = useState(false);

  return (
    <>
      <h1>{tableName}</h1>
      <InputGroup className="mb-3">
        <SearchBar state={searchTerm} setState={setSearchTerm} />
        <FilterButton name={"Current Players"} state={currentPlayers} setState={setCurrentPlayers} />
        <FilterButton name={"RC"} state={rookieCards} setState={setRookieCards} />
        <FilterButton name={"Auto"} state={autograph} setState={setAutograph} />
        <FilterButton name={"Patch"} state={patch} setState={setPatch} />
      </InputGroup>
      <Table striped bordered hover responsive variant="dark">
        <thead>
          <tr>
            <th>Picture:</th>
            <th>First Name:</th>
            <th>Last Name:</th>
            <th>Seasons Played:</th>
            <th colSpan={2}>Number of Cards:</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {players.map(values => (
            <PlayerRow firstName={values.firstName} lastName={values.lastName} seasonsPlayed={values.seasonsPlayed} seasons={values.seasons} cardData={cards.filter(nameToFilter((values.firstName + " " + values.lastName)))} />
          ))
          }
        </tbody>
      </Table>
    </>
  );
}

export default CardTable;
