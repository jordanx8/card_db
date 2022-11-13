import React, { useState } from 'react';
import PlayerRow from "./PlayerRow";
import Table from 'react-bootstrap/Table';
import InputGroupHandler from './InputGroupHandler';
import { nameToFilter } from '../util/filters';

function CardTable(props) {
  const [cards, setCards] = useState(props.cardData.default);
  const [players, setPlayers] = useState(props.playerData.default);

  return (
    <>
      <h1>{props.tableName}</h1>
      <InputGroupHandler setSearchString={props.setSearchString} cards={cards} setCards={setCards} players={players} setPlayers={setPlayers} />
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
