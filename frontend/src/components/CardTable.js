import React, { useState } from 'react';
import PlayerRow from "./PlayerRow";
import Table from 'react-bootstrap/Table';
import InputGroup from 'react-bootstrap/InputGroup';
import FilterButton from './FilterButton';
import SearchBar from './SearchBar';
import { useQuery } from "@apollo/client";
import { nameToFilter } from '../util/util';
import { GET_PLAYERS_QUERY, GET_CARDS_QUERY } from '../util/queries';
import * as playerData from '../test_player_data.json';
import * as cardData from '../test_card_data.json';

function CardTable({ tableName }) {
  const [cardData, setCardData] = useState([]);
  const { data: playersData, loading: playersLoading, error: playersError } = useQuery(GET_PLAYERS_QUERY);
  const { data: cardsData, loading: cardsLoading, error: cardsError } = useQuery(GET_CARDS_QUERY, {
    variables: { tableName: "Pelicans" },
    onCompleted: (data) => {
      console.log(data.cards)
      setCardData(data.cards)
    }
  });

  const [searchTerm, setSearchTerm] = useState('');
  const [currentPlayers, setCurrentPlayers] = useState(false);
  const [rookieCards, setRookieCards] = useState(false);
  const [autograph, setAutograph] = useState(false);
  const [patch, setPatch] = useState(false);

  if (playersLoading || cardsLoading) return "Loading...";
  if (playersError) return <pre>{playersError.message}</pre>
  if (cardsError) return <pre>{cardsError.message}</pre>

  return (
    <>
      <h1>{tableName}</h1>
      <InputGroup className="mb-3">
        {/* TODO: Filter Buttons to work */}
        <SearchBar state={searchTerm} setState={setSearchTerm} allCards={cardsData.cards} cardState={cardData} setCardState={setCardData}/>
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
          {playersData.players.map(values => (
            <PlayerRow firstName={values.firstName} lastName={values.lastName} seasonsPlayed={values.seasonsPlayed} seasons={values.seasons} cardData={cardData.filter(nameToFilter((values.firstName + " " + values.lastName)))} />
          ))
          }
        </tbody>
      </Table>
    </>
  );
}

export default CardTable;
